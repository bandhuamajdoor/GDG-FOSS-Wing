# Walk through what the original program does with an input, step by step, naming the functions/blocks it passes through.
main() reads the CSV and sends all the rows to parseTx(). parseTx() goes through each row, takes the name and amount, and keeps track of how much each person paid. It then calculates the average and adjusts everyone's balance. After that, main() sends those balances to settle(), which figures out who owes whom and prints the final transactions.

# Where did you hook your new feature in, and why there?
I placed the category tracking in parseTx() as that's where each CSV row is being read already. If a row has a third value , I add that amount to that category . I kept it separate from the current balance calculation so the settle() function doesn't have to change.

# Name one thing about the original code you think is weak or would refactor, and why.
I would take out the error part in the strconv. ParseFloat() method; because what happens then is that if someone incorrectly puts values in the .CSV file, then it treats the value just like zero without telling that something is wrong at least.
