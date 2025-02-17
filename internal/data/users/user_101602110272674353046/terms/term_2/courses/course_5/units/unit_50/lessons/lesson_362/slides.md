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

# Lesson 1.3: Practice with 'if' statements and functions

# **Warmup**


- Make a new directory `lesson_3` in the `unit_1` directory and open it in VS Code
- Create a new file `lesson_3.py` and prepare to run example code from the slides.

```text
/python
  /unit_1
    /lesson_3
      /lesson_3.py
```
  

# Python Review: `if` Statements and Functions


# `if` Statements

- Used to execute code conditionally
- Syntax:
  ```python
  if condition:
      # Code runs if condition is True
  elif another_condition:
      # Code runs if another_condition is True
  else:
      # Code runs if no condition is True
  ```
- Example:
  ```python
  x = 10
  if x > 5:
      print("x is greater than 5")
  elif x == 5:
      print("x is exactly 5")
  else:
      print("x is less than 5")
  ```

# Boolean Expressions

- Comparisons:
  - `==` (equal)
  - `!=` (not equal)
  - `>` (greater than)
  - `<` (less than)
  - `>=` (greater than or equal to)
  - `<=` (less than or equal to)

# Boolean Expressions (continued)

- Logical Operators:
  - `and` (both conditions must be True)
  - `or` (at least one condition must be True)
  - `not` (negates a condition)
- Example:
  ```python
  age = 18
  if age >= 18 and age < 21:
      print("You can vote but not drink.")
  ```

# Functions

- Reusable blocks of code
- Syntax:
  ```python
  def function_name(parameters):
      # Code block
      return value  # (Optional)
  ```
- Example:
  ```python
  def greet(name):
      return "Hello, " + name + "!"
  
  print(greet("Alice"))
  ```

# Function Parameters

- Positional Arguments:
  ```python
  def add(a, b):
      return a + b
  
  print(add(3, 5))  # Outputs: 8
  ```
- Default Arguments:
  ```python
  def greet(name="Guest"):
      return "Hello, " + name + "!"
  
  print(greet())  # Outputs: Hello, Guest!
  ```
- Keyword Arguments:
  ```python
  def describe(color, shape):
      return f"A {color} {shape}"
  
  print(describe(shape="circle", color="blue"))
  ```
# `return` vs `print`

- `return` sends a value back to the caller
- `print` just displays output
- Example:
  ```python
  def square(x):
      return x * x
  
  result = square(4)  # Stores the value 16
  print(result)       # Displays 16
  ```

# Recap

- `if` statements allow conditional execution
- Functions allow reusable and modular code
- Boolean expressions control logic
- Functions can take parameters and return values

**Practice:**
1. Write an `if` statement to check if a number is positive, negative, or zero.
2. Define a function that takes two numbers and returns their product.

# Assignment 1.3

- See files

