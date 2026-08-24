package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func parseTx(lines [][]string)(map[string]float64, map[string]float64){
	b := make(map[string]float64)
	w := make(map[string]float64)
	var tot float64

	for _, p := range lines {
		if len(p) < 2 {
			continue
		}
		usr := p[0]
		amt, _ := strconv.ParseFloat(p[1], 64)
		b[usr] += amt
		tot += amt

		c := "Not Specified"
        if len(p) >= 3 {
            c = p[2]
        }
		w[c] += amt
	}

	
	var avg float64
	if len(b) > 0 {
		avg = tot / float64(len(b))
	}

	for k := range b {
		b[k] -= avg
	}
	return b,w
}

type kv struct {
	k string
	v float64
}

func settle(b map[string]float64) []string {
	var debtors, creditors []kv
	for k, v := range b {
		if v < 0 {
			debtors = append(debtors, kv{k, v})
		} else if v > 0 {
			creditors = append(creditors, kv{k, v})
		}
	}

	sort.Slice(debtors, func(i, j int) bool { return debtors[i].v < debtors[j].v })
	sort.Slice(creditors, func(i, j int) bool { return creditors[i].v > creditors[j].v })

	var txs []string
	i, j := 0, 0

	for i < len(debtors) && j < len(creditors) {
		d := &debtors[i]
		c := &creditors[j]

		amt := -d.v
		if c.v < amt {
			amt = c.v
		}

		txs = append(txs, fmt.Sprintf("%s owes %s $%.2f", d.k, c.k, amt))

		d.v += amt
		c.v -= amt

		if d.v >= -1e-9 {
			i++
		}
		if c.v <= 1e-9 {
			j++
		}
	}
	return txs
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run split.go <csvfile>")
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		return
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	lines, _ := r.ReadAll()

	bals, cats := parseTx(lines)

	fmt.Println("Spent per category:")
    for k, total := range cats {
        fmt.Printf("%s $%.2f\n", k, total)
    }

	res := settle(bals)
	for _, r := range res {
		fmt.Println(r)
	}
}