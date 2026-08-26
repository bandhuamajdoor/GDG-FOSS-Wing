The code supposedly tracks the durations of shifts for individuals as per IN and OUT timing data

The bug was that instead of adding upto total duration, the duration of different shifts kept resetting the total duration
Another issue was the tricky case where the IN timings were being updated even before a valid OUT was found

The bug was found by comparing expected and generated values from the code since the durations didnt add up

The first issue is fixed by initialising the total time with 0 and accumulating valid durations in it
The second issue was fixed by adding condition so as to only update IN time when an OUT had already been registered for the previous IN time

Additional question's answer:
Acc. to readme, expected work time of Alice was 105 mins which based on code should have been calculated as (100-10)+(215-200) making it 105, however the output receibed was only 15 which is the later half of the sum, hence the bug of time being reset instead of being accumulated is discovered

the code is expected to output total work duration even if the inputs are tricky
In tricky test case the code should have worked such that daves time worked be (50-10) 40 however the output was actually (50-30) being 20
Hence another vulnerability of updating IN time before a valid OUT time is found was discovered


