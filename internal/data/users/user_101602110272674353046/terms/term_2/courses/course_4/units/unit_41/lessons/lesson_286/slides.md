---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Warmup

- Place your phones in your bags and your bags against the wall
- Within your `unit_4` directory, create a new directory `lesson_2`
- Open your `lesson_2` directory in VS Code 
- Open today's lesson on the course website.

# Raising and Handling Exceptions in Python

## What are Exceptions?

- An exception is an error detected during execution.
- Interrupts normal program flow.
- Examples:
  - `ZeroDivisionError`
  - `IndexError`
  - `KeyError`
  - `TypeError`

# Raising Exceptions

- Use `raise` to generate an exception.

```python
raise ValueError("Invalid value provided!")
```

- Can be used to enforce constraints in functions.

```python
def set_age(age):
    if age < 0:
        raise ValueError("Age cannot be negative!")
    return age
```

# Handling Exceptions

- Use `try` and `except` to catch exceptions.

```python
try:
    x = int("abc")
except ValueError as e:
    print("Error:", e)
```

- Prevents program from crashing.

# Multiple Exception Handling

- Handle multiple exceptions separately.

```python
try:
    num = int(input("Enter a number: "))
    result = 10 / num
except ValueError:
    print("Invalid input! Please enter a number.")
except ZeroDivisionError:
    print("Cannot divide by zero!")
```

# Using `else` and `finally`

- `else`: Runs if no exception occurs.
- `finally`: Runs no matter what.

```python
try:
    file = open("data.txt", "r")
    content = file.read()
except FileNotFoundError:
    print("File not found!")
else:
    print("File read successfully!")
finally:
    file.close()
```

# Creating Custom Exceptions

- Define custom exceptions by subclassing `Exception`.

```python
class CustomError(Exception):
    pass

raise CustomError("This is a custom error!")
```

# Summary

- Use `raise` to trigger exceptions.
- Use `try-except` to handle exceptions.
- Handle multiple exceptions as needed.
- Use `else` and `finally` for better control.
- Create custom exceptions for specific use cases.

# Assignment 4.2

- [Practice problems](./files/assignment_1_4_2.html)
