The code parses data from a .csv and expects atleast two columns in one line, the first column contains a name of a person and second column contains how much they paid, this data is stored in a map, also a net total expenditure is kept record of

The total is divided equally for the group using length of map and then stored as average amount every person needs to pay

this avg amount is subtracted from the amount every individual has paid till now and stored in the same map
This map with the subtracted amounts is returned

Now the people in the map are split into debitors (whose net difference after avg is subtracted is <0) and creditors(whose net difference after avg is subtracted is >0)
these two lists are sorted separately so as to compare the person with most owed to the person with most given

After comparing the debitor creditor and amount owed is stored separately to print

ADDED FEATURE:-
By allowing the code to read a third column if it exists, we allow to map expenditure with categories, now the first function returns 2 maps, one being the original and other stores category with expenditure, a fall back to Non specified is kept incase of no category related data

This map of categories is printed before the settlement data
