The objective of the program was to calculate the total time duration for which each user was active.

Each line of input in sample_input.txt contained information from the user in the form:- NAME IN/OUT TIME_STAMP where
IN/OUT indicates whether the user logged in or out of the session and
TIME_STAMP indicates the time at which the user logged in or out of the system

The code had a small error which returns correct login duration for a single session but when there are multiple sessions involved, it overwrites the stored duration of the previous session with the new one instead of adding them to return the total time

To fix this I changed the line "res[p]=d" to "res[p]=(res[p] || 0)+d" so that the result variable calculates the total time of all sessions instead of a single one

There was another bug present in the code which I realised later on that if a user inputs IN multiple times before an OUT the session duration is overwritten. Fixed the code from st[p]=undefined to 
if (st[p] === undefined) {
    st[p] = parseInt(t);
}
Now if a user inputs IN multiple times before an OUT the duration from the first IN to OUT is calculated.