---
layout: none
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Warmup

- Make a new directory `unit_3` in your `python` directory
- Make a new directory `lesson_1` in the `unit_3` directory and open it in VS Code
- Create a new file `lesson_3.py` and prepare to run example code from the slides.

```text
/python
  /unit_3
    /lesson_1
      /lesson_1.py
```

# Reminder

- Always back up your files to `.zip` file on google drive at the end of class!

# Introduction to While Loops in Python

## What is a While Loop?
A `while` loop in Python is used to repeatedly execute a block of code as long as a specified condition is `True`.

**Syntax:**
```python
while condition:
    # Code to execute
```

# Basic Example
```python
count = 0
while count < 5:
    print("Count is:", count)
    count += 1
```
# Output:
```text
Count is: 0
Count is: 1
Count is: 2
Count is: 3
Count is: 4
```

# Infinite Loop
A `while` loop can run indefinitely if the condition never becomes `False`.

```python
while True:
    print("This will run forever!")
```
Use **Ctrl + C** to stop execution.

# Using a Break Statement
The `break` statement is used to exit a loop prematurely.

```python
count = 0
while count < 10:
    if count == 5:
        break
    print(count)
    count += 1
```

# Output:
```text
0
1
2
3
4
```

# Using a Continue Statement
The `continue` statement skips the rest of the code in the loop and moves to the next iteration.

```python
count = 0
while count < 5:
    count += 1
    if count == 3:
        continue
    print(count)
```
## Output:
```text
1
2
4
5
```

# Else Clause with While Loop
A `while` loop can have an `else` block, which executes when the loop condition becomes `False` (unless `break` is used).

```python
count = 0
while count < 3:
    print(count)
    count += 1
else:
    print("Loop finished")
```

## Output:
```text
0
1
2
Loop finished
```

# Practical Example: User Input
```python
password = "secret"
user_input = ""
while user_input != password:
    user_input = input("Enter password: ")
print("Access granted!")
```

## Summary
✅ A `while` loop runs as long as its condition is `True`.
✅ Use `break` to exit early.
✅ Use `continue` to skip an iteration.
✅ `else` runs when the loop finishes normally.

# Assignment 3.1

- [Solve these problems](./files/assignment_1_3_1.py)
- Complete CodeHS Assignments 2.17 and 2.18
