\# Walk through what the original program does with an input, step by step, naming the functions/blocks it passes through.





1.main() first checks os.Args.If no CSV file is provided, it prints the usage message and stops.



2.It opens the CSV file using os.Open().If the file cannot be opened, it returns.



3.It creates a CSV reader using csv.NewReader(f) and reads all the rows using ReadAll().The result is stored in lines.



4.main() sends these rows to parseTx(lines).



5.In parseTx(),map b is created to store each person's total expense.The program then goes through every row using the loop



6.For each row,it first checks if len(p)<2,invalid rows with fewer than two columns are skipped.



7\. p\[0] is taken as the user's name and p\[1] is converted from a string to float using strconv.ParseFloat().



8.The amount is added to that user's total in the map using b\[usr]+= amt. At the same time, he total expense tot is updated.



9.After all rows are processed, parseTx() calculates the average expense per person using tot/float64(len(b)).



10\. It then subtracts this average from every person's total in b. This tells us whether the person owes money or should receive money..



11.parseTx() returns the balance map to main().



12.main() passes this balance map to settle(bals).



13 settle() separates the users into debtors(negative balance) and creditors(positive balance)



14.It sorts both groups and starts matching debtors with creditors.For every match, it calculates the amount to transfer and creates a transaction such as Bob owes Alice $20.00.



15.After each transaction, the balances of the two people are updated. This continues until either all debtors or all creditors are finished.



16.settle() returns the list of transactions, and main() prints them.







\# Where did you hook your new feature in, and why there?



The feature added was optional third column in the CSV for expense categories.



Existing code already reads the first two columns as the user's name and expense amount. I kept that logic unchanged and added a separate categories map to store the total amount spent in each category. 



While processing each row, the program checks whether a third column exists. If it does,p\[2] is taken as the category and the expense amount is added to that category's total.



I added the feature in parseTx() because this is where each CSV row is already being processed, so the category and amount are available there without changing the existing settle() logic.





\# Name one thing about the original code you think is weak or would refactor, and why.



A weakness may be that the code ignores errors while reading and converting the input.I would check these errors instead of ignoring them, so the program can show an error when the input is invalid or the CSV cannot be read.

