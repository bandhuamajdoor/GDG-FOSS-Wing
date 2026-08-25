Here is a quick breakdown of what went wrong in the Task 2 code and how I fixed it.

The program is basically supposed to read a log file full of 'IN' and 'OUT' timestamps and calculate the total amount of time each user spent logged in. 

But there was a pretty massive logical flaw in how it tracked the time. If a user logged in, logged out, and then logged in again later, the script would completely erase their first session. Instead of adding the new time to their total, it just blindly overwrote the variable with the newest session's time. 

I found this by running the script against `sample_input.txt` and noticing Alice only had 15 minutes total, even though looking at the raw text file, she clearly had a 90-minute session earlier on that just vanished from the final output. 

To fix it, I went to line 24 and changed `res[p] = d;` to `res[p] = (res[p] || 0) + d;`. 

This just adds a simple check: if the user already has some time logged, add the new duration to their running total instead of wiping the slate clean. If they don't have time logged yet, it defaults to 0 and adds the first session.