what the program does
It just reads a log file of IN and OUT timestamps and calculates the total active time for each user.

the bug
If a user logs in, logs out, and then logs in again later, their earlier time gets completely wiped. The code was just overwriting the total time variable instead of adding to it.

how I found it
I ran the script on sample_input.txt and noticed Alice's time was way off. She clearly had a 90-minute session and a 15-minute session in the logs, but the program output only gave her 15 minutes. It was pretty obvious her first session was just being erased by the second one.

how I fixed it
I went to line 24 and stopped it from blindly overwriting the value. I changed res[p] = d; to res[p] = (res[p] || 0) + d;. Now it checks if they already have time logged, and if they do, it just adds the new session to their running total.