The code supposedly tracks the durations of shifts for individuals as per IN and OUT timing data

The bug was that instead of adding upto total duration, the duration of different shifts kept resetting the total duration
Another issue was the tricky case where the IN timings were being updated even before a valid OUT was found

The bug was found by comparing expected and generated values from the code since the durations didnt add up

The first issue is fixed by initialising the total time with 0 and accumulating valid durations in it
The second issue was fixed by adding condition so as to only update IN time when an OUT had already been registered for the previous IN time

And in today's fast-paced world of software development-

