Originally , the code read the CSV, mapped each person to the total amount they paid, and tracked the overall group expenditure.

It calculated the equal share (total / number of people) and subtracts this from what each person actually paid to find their net balance.

Settling: It separates the group into debtors (negative balance) and creditors (positive balance), sorts them by amount, and matches them up to figure out exactly who owes whom.

I updated the parser to check for an optional third column.

If a category exists, the expense is stored into a second map dedicated to tracking category totals. (If missing, it falls back/skips).

The parsing function now returns both the balance map and the category map.

The main function catches this new map and prints the category subtotals just before running the original settlement logic.