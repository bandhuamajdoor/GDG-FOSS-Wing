The program was giving Alice 15 minutes instead of 105 minutes when  user had multiple sessions.



When a user logged out, the duration of that session was assigned directly to res\[p].This replaced the user's previous total instead of adding the new session duration to it.



so I changed the result  so that first session initializes the users

total and during later sessions result is added to the existing total 



This is the part I changed
if(res\[p]===undefined) {

   res\[p]=d;

} else {

   res\[p]+=d;

}



And in the tricky input 



Dave IN 10

Dave IN 30

Dave OUT 50



when Dave enters at 10 ,program stores 10 as his start time and again when it sees 30 as IN ,start time gets replaced with 30 .So the program calculates 50-30 =20 duration 



So repeated IN events should be ignored while a user is already logged in and this can be fixed by storing the st\[p] only when it is undefined


if(ax==='IN'){

   if(st\[p]===undefined){

       st\[p]=parseInt(t);

   }

}






