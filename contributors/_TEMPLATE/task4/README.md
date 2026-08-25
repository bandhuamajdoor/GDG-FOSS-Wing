Write 10–15 lines explaining your Task 2 bug fix to someone who has never seen the code. Open your task4/README.md file, remove the template explantion and answer these four points:

The Context: What is the program supposed to do in plain English? The Problem: What was the bug actually doing? The Investigation: How did you figure out where the bug was hiding? The Solution: How did you fix it, and why does your fix work?

Write this yourself. No AI. We want your voice. If your README opens with "In today's fast-paced world of software development..." we will instantly know, and a single tear will roll down the reviewer's cheek. Just explain it like you are talking to a classmate.

Commit: "Add Task 2 bug writeup"

**The Context** -
This Code was used to log the time, spent by users in the system. So, If an user logs IN at ‘a’ minutes and logs OUT at ‘b’ minutes, so total time spent will be (b-a) minutes. Hence, we want it to keep adding the time spent by a user in the system if the user logs IN and OUT again.

**The Problem** -
If a user already existed and logged into system once again. After logging out, its time was not added to already existing time spent, instead it overwrote the existing time. There was another problem in which If a user logged IN once again without logging OUT, its old IN time was overwritten and hence it calculated time spent from new IN rather than OLD.

**Investigation** -
I used AI to understand code and its structure i.e how it started and logged the time spent. After dry running the code over sample_input.txt , it was found that it is not adding the time if user already existed, instead it overwrote the old time spent. Dry running it over tricky_input.txt helped me understand the second problem i.e overwriting of old IN, if user did not log OUT.

**Solution** -
I made following two changes in code:

1. res[p] = (res[p] || 0) + d; 
In this, now my code adds up new time spent in the old time already logged in. Also, if user is new i.e it is not logged IN earlier, then it adds the time spent in 0.

  2. if (st[p] === undefined)
            {
                st[p] = parseInt(t); 
            }
Now, if an user logs IN once again without logging OUT the last entry, It does not overwrite the old IN, old IN remains the original IN time.