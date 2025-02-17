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

# **Warmup**

- Make a new directory `unit_2` in your `python` directory
- Make a new directory `lesson_1` in the `unit_2` directory and open it in VS Code
- Create a new file `lesson_1.py` and prepare to run example code from the slides.

```text
/python
  /unit_2
    /lesson_1
      /lesson_1.py
```


# Lesson 2.1: Nested loops and complex conditionals

# What Are Nested Loops?

- A **nested loop** is a loop inside another loop.
- The inner loop executes **fully** for each iteration of the outer loop.
- Used when working with multi-dimensional data structures.

**Example:**
```python
for i in range(3):
    for j in range(2):
        print(f"i: {i}, j: {j}")
```

# Common Use Cases

## 1. Working with 2D Lists (Matrices)
```python
matrix = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
for row in matrix:
    for element in row:
        print(element, end=' ')
    print()
```

# 2. Generating Multiplication Tables
```python
for i in range(1, 6):
    for j in range(1, 6):
        print(f"{i} x {j} = {i*j}", end='\t')
    print()
```

# 3. Combinations & Permutations
```python
for x in ["A", "B", "C"]:
    for y in [1, 2, 3]:
        print(x, y)
```

# Performance Considerations

## 1. **Time Complexity**
- Nested loops can lead to **O(n²)** or worse time complexity.
- Example: Finding duplicates in a list.
```python
nums = [1, 2, 3, 2, 4, 3]
for i in range(len(nums)):
    for j in range(i+1, len(nums)):
        if nums[i] == nums[j]:
            print(f"Duplicate: {nums[i]}")
```

# 2. **Memory Usage**
- Storing results inside a nested loop can use a lot of memory.
- Example: Storing all pairs.
```python
pairs = [(x, y) for x in range(100) for y in range(100)]
```

# Things to Watch Out For

## 1. **Unnecessary Nested Loops**
- Can often be replaced with built-in functions or optimized approaches.
```python
# Slow
for i in range(len(lst)):
    for j in range(len(lst[i])):
        process(lst[i][j])

# Better
for row in lst:
    for item in row:
        process(item)
```

# 2. **Breaking Out of Nested Loops**
- Use `break` or `return` carefully to exit loops early.
```python
for i in range(5):
    for j in range(5):
        if i == 2 and j == 3:
            break  # Exits inner loop only
```

# 3. **Using itertools for Efficiency**
```python
from itertools import product
for x, y in product(range(3), range(3)):
    print(x, y)
```

# Summary

- Nested loops allow iteration over multi-dimensional data.
- Be mindful of performance, especially **O(n²) complexity**.
- Use **optimized approaches** when possible.
- Consider **`itertools.product()`** for generating combinations.

# Complex Conditionals in Python

## What Are Complex Conditionals?

- Conditionals involve **if, elif, else** statements.
- Complex conditionals use **multiple conditions** combined with logical operators.
- Essential for **decision-making** in programs.

# **Example:**
```python
age = 25
if age >= 18 and age < 65:
    print("Adult")
elif age >= 65:
    print("Senior")
else:
    print("Minor")
```

# Common Use Cases

## 1. Validating User Input
```python
user_input = input("Enter a number: ")
if user_input.isdigit() and int(user_input) > 0:
    print("Valid input")
else:
    print("Invalid input")
```

# 2. Checking Multiple Conditions
```python
time = 14
if time >= 9 and time <= 17:
    print("Business hours")
else:
    print("Closed")
```

# 3. Nested Conditionals
```python
def classify_number(num):
    if num > 0:
        if num % 2 == 0:
            print("Positive even number")
        else:
            print("Positive odd number")
    else:
        print("Non-positive number")
```

# Things to Watch Out For

## 1. **Logical Operator Precedence**
- `and` has higher precedence than `or`.
```python
x = True or False and False  # Evaluates to True
```
- Use parentheses for clarity:
```python
x = (True or False) and False  # Evaluates to False
```

# 2. **Boolean Short-Circuiting**
- `and` stops at the first False condition.
- `or` stops at the first True condition.
```python
x = 0
y = 10 or x / 0  # No division error
```

# 3. **Overuse of Nested Conditionals**
- Can make code **hard to read**.
- Use **logical operators** to simplify.
```python
# Before
if age >= 18:
    if age < 65:
        print("Adult")

# After
if 18 <= age < 65:
    print("Adult")
```

# Best Practices

1. **Use parentheses** for readability in complex expressions.
2. **Simplify nested conditionals** with logical operators.
3. **Avoid redundant checks** (e.g., `if x == True`, just use `if x`).
4. **Consider using match-case** (Python 3.10+).

```python
def process_status(code):
    match code:
        case 200:
            print("Success")
        case 400:
            print("Bad Request")
        case 500:
            print("Server Error")
        case _:
            print("Unknown Code")
```

# Summary

- Complex conditionals allow **advanced decision-making**.
- Be mindful of **operator precedence and short-circuiting**.
- Simplify **nested conditionals** for better readability.
- Use **match-case** where applicable (Python 3.10+).
