
## 1. What exactly was wrong, and why did it only show up on some inputs? Name the specific line/condition.


The program was not correctly handling multiple sessions of the same person. After every OUT, the duration of the current session was stored in res[p]. This caused the previous session's total to be replaced instead of adding the new session duration to it.

The original code contain this:

res[p] = d;

which i have changed to

res[p] = (res[p] || 0) + d;

Now, the first session is stored as the user's total, and the duration of every later session is added to the existing total


## 2. Paste your output for tricky_input.txt. Does your fix handle this one correctly? If yes, why is it robust; if no, what class of input still breaks it?

The output for `tricky_input.txt` is:

Dave: 40 mins
Eve: 60 mins

Yes, my fix handles this input correctly. When Dave appears with another IN at 30, his first IN time of 10 is not overwritten. The second IN is ignored until his OUT at 50, so his active time is calculated as 50 - 10 = 40 minutes.

The fix also keeps the total of multiple completed sessions by adding each new session duration to the existing total.