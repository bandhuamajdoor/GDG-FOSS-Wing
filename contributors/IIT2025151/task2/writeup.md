# Task 2 — Session Tracker Bug Fix

## What the program does

The program reads the IN and OUT records of different users and calculates the total time for which each user was active.

For example, if a person enters at 10 and leaves at 100, their active time is 90 minutes.

## What was the bug

The bug was that when a person had multiple sessions, the program was not adding the time of all the sessions.

The original code had:

res[p] = d;

This means that whenever a new session was completed, its duration was directly stored in `res[p]`. So if Alice had two sessions of 90 minutes and 15 minutes, the second value 15 replaced the first value 90. Therefore, the program gave Alice 15 minutes instead of 105 minutes.

## How I found the bug

First, I ran the program with `sample_input.txt` and compared the output with the expected output given in the README.

The program gave:

Alice: 15 mins

But the expected output was:

Alice: 105 mins

I noticed that Alice had two sessions. I then checked the code and found that `res[p] = d` was replacing the previous value instead of adding the new session time to it.

## How I fixed it

I changed:

res[p] = d;

to:

res[p] = (res[p] || 0) + d;

Here, if the person has no previous session, `res[p]` is undefined, so `0` is used. If the person already has a previous session, the previous value is used and the current duration is added to it.

So for Alice:

90 + 15 = 105 minutes

## Testing

After making the change, I ran the program again with `sample_input.txt`.

The output became:

Alice: 105 mins
Bob: 45 mins
Charlie: 30 mins

I also tested `tricky_input.txt`.

For example:

Dave IN 10
Dave IN 30
Dave OUT 50

The program gives Dave 20 minutes because the second `IN` replaces the first starting time, so the calculation becomes:

50 - 30 = 20 minutes.

The fix correctly handles users having multiple completed sessions by adding their session durations instead of overwriting the previous total.

## Root cause

The root cause was that the program was storing only the current session duration in `res[p]` instead of accumulating the current duration with the user's previous session duration.