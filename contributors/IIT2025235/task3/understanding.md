# Task 3 - Expense Splitter

## What the program does

The program reads a CSV file containing people's expenses and calculates the minimum number of transactions required to settle the expenses fairly.

The original CSV format was:

Name,Amount

The feature request adds support for an optional third column:

Name,Amount,Category

## What I changed

I extended the existing `parseTx()` function to support the optional category column.

I added a `categories` map to store the total amount spent in each category.

For every expense containing a category, its amount is added to the corresponding category total.

For example:

Alice,50.00,Food
Charlie,30.00,Food

produces:

Food: $80.00

I also sorted the category names alphabetically before printing them so that the output remains consistent.

## Settlement Logic

I did not modify or rewrite the existing `settle()` function.

The program first calculates and prints the category subtotals and then uses the existing `settle()` function to calculate the settlement transactions.

## Testing

I tested the new feature using a CSV containing multiple categories and multiple expenses in the same category.

The test input was:

Alice,50.00,Food
Bob,10.00,Travel
Charlie,30.00,Food
Alice,20.00,Food
Bob,10.00,Travel

The output was:

Food: $100.00
Travel: $20.00
Bob owes Alice $20.00
Charlie owes Alice $10.00

The category totals are correct:

Food = 50 + 30 + 20 = $100.00
Travel = 10 + 10 = $20.00

I also tested the original expenses.csv, which uses the old two-column format.

The output remained:

Bob owes Alice $20.00
Charlie owes Alice $10.00

This confirms that the new category feature works without breaking the existing settlement functionality.

## Summary

The program now supports an optional third CSV column for expense categories, calculates and displays category subtotals before the settlement transactions, and preserves the original `settle()` logic.