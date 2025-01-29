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

# Lesson 0.2

# Warmup

Create a folder `lesson_2` and open it in VS Code

Create a new python file and insert the following code:

```python
name = "Alice"
age = 12
```

Add a line to print the name and age.

# Review: If statements in Python

## What is an If Statement?
- **Definition**: An if statement is used to execute a block of code if a condition is `True`.
- **Why It's Important**:
  - Adds decision-making to your programs.
  - Allows conditional execution of code based on logic.
- **Syntax**:
  ```python
  if condition:
      # Code to execute if condition is True
  ```

# Basic Example
```python
age = 18
if age >= 18:
    print("You are an adult!")
```
**Output**:
```text
You are an adult!
```
- **Key Points**:
  - The `condition` (`age >= 18`) evaluates to `True`.
  - The indented block is executed.

# Else Clause
- **Purpose**: The `else` block executes when the condition in `if` is `False`.
- **Syntax**:
  ```python
  if condition:
      # Code if condition is True
  else:
      # Code if condition is False
  ```
- **Example**:
  ```python
  temperature = 30
  if temperature > 20:
      print("It's warm.")
  else:
      print("It's cold.")
  ```
**Output**:
```text
It's warm.
```

# Elif Clause
- **Purpose**: The `elif` (short for "else if") allows checking multiple conditions.
- **Syntax**:
  ```python
  if condition1:
      # Code if condition1 is True
  elif condition2:
      # Code if condition2 is True
  else:
      # Code if none of the above are True
  ```
# **Example**:
  ```python
  score = 85
  if score >= 90:
      print("Grade: A")
  elif score >= 80:
      print("Grade: B")
  else:
      print("Grade: C")
  ```
**Output**:
```text
Grade: B
```

# Nested If Statements
- **Definition**: An `if` statement inside another `if`.
- **Example**:
  ```python
  num = 10
  if num > 0:
      if num % 2 == 0:
          print("Positive and even.")
      else:
          print("Positive and odd.")
  ```
**Output**:
```text
Positive and even.
```

# Common Mistakes
1. **Indentation Errors**:
   - Python uses indentation to define code blocks.
   - Ensure consistent use of spaces or tabs.
   - **Bad Example**:
     ```python
     if True:
     print("Missing indentation")
     ```
2. **Using `=` Instead of `==`**:
   - `=` is assignment; `==` is comparison.
   - **Example**:
     ```python
     if x = 5:  # Incorrect
     if x == 5:  # Correct
     ```

# Practice Problems
1. **Problem 1**: Write a program that checks if a number is positive, negative, or zero.
2. **Problem 2**: Check if a user-entered year is a leap year.
3. **Problem 3**: Determine if a number is divisible by both 3 and 5.

# Recap
- **Key Concepts**:
  - `if`, `elif`, and `else` statements allow decision-making.
  - Conditions must evaluate to `True` or `False`.
  - Indentation is critical for structuring your code.
- **Next Steps**: Apply these concepts in larger programs to handle more complex logic.

# Agenda

Continue working on `practice.py` (see yesterday's files) together in class.

[practice.py](../357/files/practice.py)
