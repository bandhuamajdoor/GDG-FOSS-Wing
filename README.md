# GDG FOSS Wing Recruitment Tasks

Welcome. These tasks decide who joins the FOSS wing. We are **not** testing whether you can memorise git commands or copy a tutorial. We are testing whether you can read unfamiliar code, fix real problems, and explain your own work like an engineer.

> *"It works on my machine."*  
> — every candidate, right before we run it on ours

**On AI tools:** You may use AI (ChatGPT, Claude, Copilot, whatever) to *learn* — to understand a concept, look up syntax, or unstick yourself. That's how real engineers work. But every task has a checkpoint where you must demonstrate understanding in your own words or against a case AI won't have seen. **AI-written code you don't understand will fail these checkpoints, and we will notice.** Submit AI slop you can't explain and you're out. Use AI to get smarter and you're exactly who we want.

---

## Prerequisites & Workflow

You will need the following tools installed on your system:
- **Git** (Use Git Bash if you are on Windows)
- **Node.js** (Required for Task 2)
- **Go / Golang** (Required for Task 3)

> *"Wait, I don't know JavaScript or Go!"*  
> Good. FOSS is rarely a single-language ecosystem, and you won't always get to work in your comfort zone. You are going to have to read documentation, figure out how to run the code, and look up syntax. Google is your friend. Welcome to open source. ☕

**For every task:**
1. **Fork** this repository and **clone** your fork.
2. Copy the `contributors/_TEMPLATE` folder and rename it to your roll number (e.g., `contributors/IIT2024245`). All your work goes in here.
3. **Create a branch** `<roll-number>-task<N>` off `main` before starting (e.g. `IIT2024245-task1`).
4. Make your changes on that branch.
5. Add a row for yourself to **`CONTRIBUTORS.md`** (repo root) for this task. See the format in that file.
6. **Push** the branch to your fork.
7. Open a **Pull Request into `main` of this original repo** (not your fork).

**Rules:**
- **One PR per task.** Mixed or messy PRs won't be reviewed.
- **PRs without a `CONTRIBUTORS.md` row won't be merged.** That row is how we track who's completed what.
- Terminal screenshots only (`git log --oneline --graph`, etc.). No GitHub web-UI screenshots.
- Meaningful commit messages. No `"update"`, `"fix"`, `"changes"`, or `"final_final_v2"`.
- Plagiarism between submissions is checked. Do your own work.

---

## Task 1 — Git Competency Gate *(the only pure-git task)*

We keep exactly one git task, because branching and conflict resolution are genuine prerequisites for contributing anywhere. 

**File:** `contributors/<roll-number>/task1/profile.txt`

1. On your task branch, add your **favorite project you've built** as line 1. Commit: `"Add favorite project"`
2. Create a branch `alt-branch` off your task branch.
3. On `alt-branch`, add your **dream role** as line 2. Commit: `"Add dream role"`
4. Switch back to your task branch. Add your **home city** as line 2. Commit: `"Add home city"`
5. Merge `alt-branch` into your task branch — this conflicts on line 2. *(Yes, on purpose. The conflict is the assignment).*
6. Resolve so the file reads, in order: favorite project, home city, dream role.
7. Commit the resolution with a message naming the conflict.

**Screenshots (in your `task1/` folder):**
- `screenshot-1.png`: `git status` (or the conflict markers) *before* resolving.
- `screenshot-2.png`: `git log --oneline --graph` *after* resolution.

**Interactive checkpoint:** Answer the questions in your `task1/answers.md` template.

---

## Task 2 — Fix a Real Bug in the Sandbox

**Directory:** `sandbox/task2/`. It's a small, working-ish program with **one deliberately planted bug** that makes it produce wrong output on certain inputs. The code runs, looks fine, passes the easy cases, and betrays you in the demo.

**Your job:**
1. Run it. Reproduce the wrong behaviour.
2. Find the bug. Fix it with the **smallest change that's actually correct** (no rewriting the whole file to dodge understanding it).
3. Copy the fixed file into your `contributors/<roll-number>/task2/` folder and commit. Commit message must describe the *root cause*, not `"fix bug"`.

**Screenshots (in your `task2/` folder):**
- `screenshot-1.png`: terminal showing the **wrong** output before your fix.
- `screenshot-2.png`: terminal showing **correct** output after your fix.

**Interactive checkpoint:** Complete your `task2/writeup.md` template. 

---

## Task 3 — Read Unfamiliar Code & Extend It

**Directory:** `sandbox/task3/` contains a small program you did **not** write, with no comments and slightly terse naming. 

**Your job:**
1. Read it until you genuinely understand the flow. Do **not** rewrite or reformat it.
2. Add **one small feature** (specified in `sandbox/task3/FEATURE.md`) that requires you to hook into the existing logic.
3. Copy the extended program into your `contributors/<roll-number>/task3/` folder and commit with a message describing the feature.
4. Match the existing code's style. Fitting in > showing off.

**Interactive checkpoint:** Complete your `task3/understanding.md` template.

---

## Task 4 — Explain Task 2 to a Beginner *(AI-free)*

**File:** `contributors/<roll-number>/task4/README.md`

Write 10–15 lines explaining your **Task 2 bug fix** to someone who's never seen the code: what the program does, what the bug was, how you found it, how you fixed it. 

- **Write this yourself. No AI.** We want *your* voice. If your README opens with *"In today's fast-paced world of software development..."* we will know, and a single tear will roll down the reviewer's cheek. Just tell us what the bug was.
- Commit: `"Add Task 2 bug writeup"`

---

## How We Evaluate

We read PRs properly. Rough weighting:
- **Task 1** — pass/fail gate. Clean git history + coherent checkpoint answers = pass.
- **Task 2** — the biggest signal. Correct root-cause reasoning and honest robustness analysis matter more than a working patch.
- **Task 3** — do you actually understand code you didn't write, and can you extend it without breaking its shape?
- **Task 4** — can you communicate? Genuine, in-your-own-words explanation.

Use AI to learn. Don't use it to fake understanding. Good luck. 🚀

<br>
<sub>*P.S. — Somewhere in the sandbox there is a variable named `temp` that has been `temp` for three years. Try not to add a `temp2`.*</sub>