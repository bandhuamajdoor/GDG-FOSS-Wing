The program was supposed to calculate the time duration of every user for which user was active 



Every row of input contains 2 forms of information one is IN or OUT where IN describes that user has entered the session and OUT describes user has ended the session



Here the program has a small bug which runs correctly for a single session but becomes problematic when there are multiple sessions as the result being stored is not counting the time duration of previous sessions

So I figured out this by running sample\_input.txt.



And to fix this I changed the result variable calculation to accumulate previous time durations also 

From res\[p]=d to res\[p]+=d



But there was also another bug which i identified with trickyinput.txt that when there are multiple IN's before an OUT , the start time is getting replaced with the newer one

So to fix this we can use the IN time only when it is still undefined else skip that

