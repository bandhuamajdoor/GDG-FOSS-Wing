The program in Task2 tracks the duration for which a user was logged IN in a particular portal and displays the duration outputs for each user accordingly. The input data supplies the user, action(IN/OUT) and the time data.

During testing, I found that instead of calculating the total duration for which a user was logged in, the code resets to calculate only the last duration for which the user was active.
(Ex: Alice logged in between 20-50 mins and then she again logged in between 80-100 mins, then the code only displays 20 mins durations instead of 50 mins).

This was resolved by keeping an accumulator in the program which tracks total duration for each user by adding all his/her valid durations.

Another issue faced was that the code showed incorrect results for a user who logs IN although he is already logged IN or logs OUT although he is already logged OUT.(maybe technical glitch)

For such cases, the code shows INVALID input for such users only. A st object is introduced to track the current duration for a user and deleted after the session is over. If a user logs IN consecutively, then the str object is already initialized the 2nd time and it throws error. Also if the user logs OUT consecutively, then the str object is already not initialized the 2nd time and it throws error.