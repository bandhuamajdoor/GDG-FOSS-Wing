# 🐛 Task 2 - Bug Fix Writeup

## 🔵 What the Program Does

The program reads a space-separated log file containing `IN` and `OUT` events for users and calculates the **total number of minutes each user was active**.

---

## 🟠 Bug Reproduction

I ran the program using:

node sandbox/task2/process_logs.js sandbox/task2/sample_input.txt

### ❌ Output Before Fix

Alice: 15 mins
Bob: 45 mins
Charlie: 30 mins

### ✅ Expected Output

Alice: 105 mins
Bob: 45 mins
Charlie: 30 mins

Alice had **2 active sessions**:

- `IN 10` to `OUT 100` = **90 minutes**
- `IN 200` to `OUT 215` = **15 minutes**

Therefore:

**90 + 15 = 105 minutes**

---

## 🔴 Root Cause

The program had the following line:

res[p] = d;

This was **overwriting the previous total** with the duration of the current active session.

For Alice:

- First session → `90` minutes
- Second session → `15` minutes
- The program replaced `90` with `15`
- Therefore, it incorrectly returned **15 minutes**

The problem was that the program was **not accumulating the durations of multiple sessions**.

---

## 🟣 Fix

I changed:

res[p] = d;

to:

res[p] = (res[p] || 0) + d;

### How the Fix Works

If an older value of `res[p]` exists from a previous session, the new duration `d` is added to it.

If no previous value exists, `res[p] || 0` gives `0`, so the current duration is added to `0`.

For Alice:

First session: **90 minutes**

Second session: **15 minutes**

Total:

**90 + 15 = 105 minutes**

---

## 🟢 Verification

After applying the fix, I ran the program again:

node sandbox/task2/process_logs.js sandbox/task2/sample_input.txt

The output was:

Alice: 105 mins
Bob: 45 mins
Charlie: 30 mins

I also tested the program with `tricky_input.txt` after the fix to verify that the change did not break the existing behavior.

---

## ⭐ Summary

The bug was caused by **overwriting the previous session total instead of accumulating it**.

The fix was a **one-line change**:

res[p] = (res[p] || 0) + d;

This correctly adds the duration of every completed session to the user's total.