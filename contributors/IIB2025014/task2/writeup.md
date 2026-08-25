The wrong output for sample input was because the function updates the total time 
instead of adding to it at every valid OUT. 

The fix here is initializing the total time for a person with 0 and adding the valid durations to it.


The tricky input shows a vulnerability of multiple IN before a valid OUT,
The original script updates intime at every IN

The fix adds a check based on the fact that intimes are deleted after the duration is recorded 
This makes sure the intime is updated only when an Intime doesnt pre exist