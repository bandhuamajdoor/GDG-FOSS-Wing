Understanding parseTx

The function named parseTx works like an accountant. It reads every line of the CSV file. Calculates the net balance for each person.

1.parseTx starts by looking at each row. It works out two figures: the amount paid by each person and the overall total spent by the whole group.

2.parseTx divides the total by the number of people. That gives the "average" the amount everyone should pay so that things are fair.

3.Finally parseTx loops through each person’s total. Subtracts the average. The result is a map of balances. If a balance is positive that person. Should be reimbursed. If a balance is negative that person underpaid. Owes money to the group.

Understanding settle

The function named settle decides who should pay whom using the number of transactions.

1.settle takes the balances map. It splits the group into two lists: debtors ( balances) and creditors (positive balances).

2.settle sorts these lists so that the biggest debtors and the biggest creditors come first.

3.settle uses a while loop and two pointers i and j to match the debtor with the biggest creditor.

4.settle calculates the amount that can move between them without either side overpaying. It records that amount as a string. Updates the balances.

5.Because floating‑point math can be slightly inaccurate, in Go settle uses a buffer, 1e-9 to check if a balance is effectively zero. When a balance reaches zero the pointer moves to the person until all balances are settled.

How I Added the Feature

I added category subtotals without breaking the existing logic by changing parseTx. I made parseTx create a map that tracks categories. Because the third column is optional I added a check, len(p) >= 3, to confirm that the row has a category before reading it. That stops the program from crashing. Now parseTx returns. The balances map and the categories map. In main() I simply loop through the categories map. Print each subtotal right before I call settle().