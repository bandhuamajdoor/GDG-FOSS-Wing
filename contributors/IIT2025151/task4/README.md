The program is basically calculating time duration of people,which means if some person has came in at a particular time and  went out  at some time ,then this program is calculating for how much time he was active there.

But,it  cannot handle when a person came in and went out for multiple times.So,the main bug was not able to tackle multiple session of a person.

When i ran a input which has multiple IN-OUT session of the same person,then the program was replacing the previous IN-OUT session with the latest session instead of  adding all the sessions for the same person.

I changed the program to keep track of the total session time for each person, so every completed IN-OUT session is added instead of overwriting the previous one.  

Later,when i ran an input containing multiple INs before an OUT,then my first IN time was overwritten with later ones ,which was a major bug .Since a person cannot enter again without first going out,so i modified the program to store an IN time only when no previous IN time is recorded. If one is already present, the new IN is ignored.