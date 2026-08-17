# Feature Request: Categories

We need to track *what* people are spending money on, not just how much. 

**The Goal:**
Update the code to support an optional third column in the CSV for the expense category (`Name,Amount,Category`). 

Before printing the settlement transactions, the program should now also print a subtotal of how much was spent in each category across the entire group.

**Example input line:**
`Alice,15.50,Food`

**Constraints:**
1. Hook into the existing parsing logic.
2. Do not break or rewrite the `settle()` logic.
3. Keep the same coding style.