# What exactly was wrong, and why did it only show up on some inputs and not others? Name the specific line/condition.



The program was giving Alice 15 minutes instead of 105 minutes when user had multiple sessions. 

When a user logged out, the duration of that session was assigned directly to res\[p].This replaced the user's previous total instead of adding the new session duration to it. 

so I changed the result so that first session initializes the users total and during later sessions result is added to the existing total 

This is the part I changed 

if(res\[p]===undefined) { 

&#x20;    res\[p]=d; 

&#x20;} else { 

&#x20;   res\[p]+=d; 

&#x20;}

# Paste your output for tricky\_input.txt. Does your fix handle this one correctly? If yes, why is it robust; if no, what class of input still breaks it?



The output for tricky\_input.txt is:

&#x20;Dave 20

&#x20;Eve 60



No, my fix does not handle this input correctly

The problem is that when Dave enters at 10 ,program stores 10 as his start time and again when it sees 30 as IN ,start time gets replaced with 30 .So the program calculates 50-30 =20 duration but the actual output should be 50-10=40



So repeated IN events should be ignored while a user is already logged in and this can be fixed by storing the st\[p] only when it is undefined 



if(ax==='IN'){

&#x20;  if(st\[p]===undefined){

&#x20;        st\[p]=parseInt(t);

&#x20;} 

}

