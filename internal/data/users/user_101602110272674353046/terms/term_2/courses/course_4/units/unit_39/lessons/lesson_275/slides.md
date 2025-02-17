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

- Make a new directory `lesson_2` in the `unit_2` directory and open it in VS Code.
- Create a new file `lesson_2.py` and prepare to run example code in today's lecture slides.

```text
/python
  /unit_2
    /lesson_2
      /lesson_2.py
```

# Introduction to `if` Statements in Python

# What is an `if` Statement?
- An `if` statement is used for **decision-making** in Python.
- It allows the program to execute certain code only if a condition is met.
- Example:

```python
age = 18
if age >= 18:
    print("You are an adult.")
```

# Basic Syntax of `if`
```python
if condition:
    # Code block executed if condition is True
```

Example:
```python
x = 10
if x > 5:
    print("x is greater than 5")
```

# Adding `else`
- The `else` block runs when the `if` condition is **False**.

```python
num = 3
if num % 2 == 0:
    print("Even number")
else:
    print("Odd number")
```

# Using `elif` (Else If)
- Allows checking multiple conditions.
- Only the **first matching condition** is executed.

```python
score = 75
if score >= 90:
    print("A grade")
elif score >= 80:
    print("B grade")
elif score >= 70:
    print("C grade")
else:
    print("Fail")
```

# Comparison Operators

| Operator | Meaning |
|----------|---------|
| `==` | Equal to |
| `!=` | Not equal to |
| `>` | Greater than |
| `<` | Less than |
| `>=` | Greater than or equal to |
| `<=` | Less than or equal to |

# Logical Operators
- Used to combine multiple conditions.

| Operator | Description |
|----------|-------------|
| `and` | Both conditions must be True |
| `or` | At least one condition must be True |
| `not` | Reverses the condition |

Example:
```python
age = 25
has_license = True
if age >= 18 and has_license:
    print("You can drive.")
```

# Nested `if` Statements
- `if` statements can be placed inside other `if` statements.

```python
num = 10
if num > 0:
    print("Positive number")
    if num % 2 == 0:
        print("Even number")
```

# `if` with User Input
- Example using `input()`:

```python
name = input("Enter your name: ")
if name:
    print(f"Hello, {name}!")
else:
    print("You didn't enter a name.")
```

# Practice Questions
1. Write an `if` statement to check if a number is positive.
2. Modify the above to also check if it's negative or zero.
3. Write an `if-elif-else` statement to grade a test score.
4. Use `and` and `or` to combine multiple conditions.
5. Get user input and check if it's greater than 10.

# Summary
- `if` statements control program flow based on conditions.
- `else` handles the False case.
- `elif` checks multiple conditions.
- Operators like `==`, `>`, and `and` help make decisions.

# Assignment 2.2

- Solve problems in [Python File](./files/assignment_1_2_2.py) and submit
  - If link doesn't work see "other files" on lesson page
- Complete CodeHS Assignments 2.13, 2.14
