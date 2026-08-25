I had a merge conflict because both branches changed line 2 of profile.txt differently. One branch added my dream role and the other added my home city.
 So, Git could not decide which change to keep. I manually resolved the conflict by removing the conflict markers and keeping the required information in the correct order.

Then I used git add and committed the merge.

The previous commit is still in the Git log because it was made on alt-branch before the merge. Merging the branch does not remove the old commits.

