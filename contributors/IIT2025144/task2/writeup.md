# What exactly was wrong, and why did it only show up on some inputs and not others? Name the specific line/condition.
he problem was res[p] = d as every new session replaced the user's previous time instead of adding to it, so it only broke when someone had multiple sessions. I changed it to: res[p] = (res[p] || 0) + d; to keep a running total.

# Paste your output for tricky_input.txt. Does your fix handle this one correctly? If yes, why is it robust; if no, what class of input still breaks it?
Dave: 20 mins
Eve: 60 mins
Yes, it works correctly because each session gets added to the user's total, so it handles both single-session and multiple-session users.
