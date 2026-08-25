 \#Walk through what the original program does with an input, step by step, naming the functions/blocks it passes through.
1.main() → opens the CSV file, reads all transaction rows using csv.Reader, and sends them to parseTx().
2.parseTx() → goes through each transaction, extracts the user, amount, and category, updates the user's total balance and category total, then 
finds the average spending, and updates each user's balance according to how their spending compares with the average.
3.main() → receives the results from parseTx(), prints the category-wise totals, and then passes the balances to settle().
4.settle() → separates debtors and creditors, sorts them, and matches them to calculate who should pay whom and how much..
5.main() →prints the final settlement returned by settle().

2.\# Where did you hook your new feature in, and why there?
1.I added the category-wise expense tracking inside parseTx().
2.While processing each transaction, parseTx() reads the category from p[2] and adds its amount to the cats map.
3.I chose parseTx() because this is already where the transaction data is being read and processed, so the category information can be collected there without changing the settle() logic.
4.main() then receives cats from parseTx() and prints the category totals.

\#Name one thing about the original code you think is weak or would refactor, and why.
In main(), the error from r.ReadAll() is ignored using lines, _ := r.ReadAll().
If reading the CSV fails, the program continues without handling the error.
I would handle this error explicitly so the program can fail safely and show what went wrong.