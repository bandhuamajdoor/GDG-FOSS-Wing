The program reads user IN and OUT logs and calculates how long each person was active.
The bug was that it replaced the previous session time instead of adding to it.
I noticed this when Alice got 15 minutes instead of the expected 105.
I traced it to res[p] = d, which was overwriting the old value.
I changed it to res[p] = (res[p] || 0) + d.
Now each session is added to the user's total instead of replacing it.
I ran the sample again and got the expected output.