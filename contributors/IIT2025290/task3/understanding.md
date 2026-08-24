# Walk through what the original program does with an input, step by step, naming the functions/blocks it passes through.
Original program was made up mainly by two functions, parseTx() and settle(). ParseTx() function calculates the total expenses by each person i.e calculates of each person's balance.

Sample Input - 
Alice,50.00
Bob,10.00
Charlie,30.00
Alice,20.00
Bob,10.00

Dry run - 
Firstly, b := make(map[string]float64)
var tot float64

Hence, b = {}
tot = 0

After that, usr := p[0]
amt, _ := strconv.ParseFloat(p[1], 64)
b[usr] += amt
tot += amt
where usr = ALice and amt = 50. So, b["Alice"] += 50 and tot = 50.

Similar goes for Bob and Charlie.

After that, again Alice spends, since Alice already exists, it runs through b[usr] += amt and hence Alice's spends increase to 70 and tot = 120. Similar happens for Bob when it spends 10 more.

Then average is calculated i.e theoretically how much one should have spent by:
var avg float64
if len(b) > 0 {
    avg = tot / float64(len(b))
}
Now, the balance is calculated by 
for k := range b {
    b[k] -= avg
}
balance = actual spending - theoretical share (i.e average)
For example, Alice spent 70 but average is 40. So, his balance is positive. Hence, he is a creditor. SImilarily, Bob is a debtor.

Now settle() function starts -
Two lists are formed debtors and creditors - var debtors, creditors []kv. Debtors are with v < 0 and creditors are one with v > 0.
Then, sorting happens in both lists i.e largest creditors and debtors first.

After this balance is updated and the one whose balance is positive are the ones who lend and one with negative balance are the ones who owe.

Hence, we get the output -
Bob owes Alice $20.00
Charlie owes Alice $10.00

# Where did you hook your new feature in, and why there?
I hooked the new feature in parseTx() because that is where each csv row is already being read and amount is bring parsed. Here, amount is already available as amt. Mainly, settle() does not need to know anything about categories. Hence, it is left untouched.

# Name one thing about the original code you think is weak or would refactor, and why.
One thing that would refractor is the unused temp variable in parseTx(). It is unnecessary and does not add anything to code.
temp := tot / 2.0
_ = temp  
Second line is there just to satisfy go's unused variable rule. Hence, I would remove both these lines. This would make parseTx() cleaner and easier to understand.