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

- Go to [VS Code Keyboard Shortcuts](https://code.visualstudio.com/shortcuts/keyboard-shortcuts-windows.pdf)
  - Identify 3 *new* keyboard shortcuts and try them out
  - Be prepared to share your favorite!

# Introduction to `for` Loops in Python

## What is a `for` Loop?
- A `for` loop is used to iterate over a sequence (like a string or a range of numbers).
- It allows you to execute a block of code multiple times.

**Example:**
```python
for char in "hello":
    print(char)
```
**Output:**
```text
h
e
l
l
o
```

# Iterating Over a Range
- The `range()` function generates a sequence of numbers.
- Commonly used with `for` loops to iterate a specific number of times.

**Example:**
```python
for i in range(5):
    print(i)
```
**Output:**
```text
0
1
2
3
4
```

# Looping Through Strings
- Strings are sequences of characters, so they can be iterated using a `for` loop.

**Example:**
```python
for char in "Python":
    print(char)
```
**Output:**
```text
P
y
t
h
o
n
```

# `for` Loop with `break` and `continue`
- `break` exits the loop early.
- `continue` skips the current iteration.

**Example:**
```python
for num in range(5):
    if num == 3:
        break
    print(num)
```
**Output:**
```text
0
1
2
```

# **Example with `continue`:**
```python
for num in range(5):
    if num == 2:
        continue
    print(num)
```
**Output:**
```text
0
1
3
4
```

# `for` with `else`
- The `else` block runs if the loop completes normally (without `break`).

**Example:**
```python
for i in range(3):
    print(i)
else:
    print("Loop finished")
```
**Output:**
```text
0
1
2
Loop finished
```

# Summary
- `for` loops are useful for iterating over sequences.
- The `range()` function helps define loop limits.
- `break` and `continue` control loop flow.
- `else` executes if no `break` occurs.

## Practice Questions
1. Write a `for` loop to print all even numbers from 1 to 10.
2. Iterate over the string "Python" and print each letter in uppercase.
3. Loop through a dictionary of countries and their capitals.

# Assignment 3.3

- CodeHS 3.3 and 3.4
- Submit solutions to [problems](./files/assignment_1_3_3.py)
  - See [hints](./files/hints_assignment_1_3_3.html)