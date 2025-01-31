---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Lesson 1.2

# Warmup

- Make a new directory `lesson_2` in the `unit_1` directory and open it in VS Code.
- Create a new file `lesson_2.py`

```text
/python
  /unit_1
    /lesson_2
      /lesson_2.py
```


# Python Arithmetic Operators

Learn the basic arithmetic operations in Python.

# What are Arithmetic Operators?

Arithmetic operators are used to perform mathematical operations on numbers.

Python supports the following arithmetic operators:
- Addition (`+`)
- Subtraction (`-`)
- Multiplication (`*`)
- Division (`/`)
- Floor Division (`//`)
- Modulus (`%`)
- Exponentiation (`**`)

# Addition (`+`)

Adds two numbers together.

```python
x = 5 + 3
print(x)  # Output: 8
```

# Subtraction (`-`)

Subtracts one number from another.

```python
x = 10 - 4
print(x)  # Output: 6
```

# Multiplication (`*`)

Multiplies two numbers.

```python
x = 6 * 7
print(x)  # Output: 42
```

# Division (`/`)

Divides one number by another. Always returns a float.

```python
x = 15 / 2
print(x)  # Output: 7.5
```

# Floor Division (`//`)

Divides one number by another and rounds down to the nearest integer.

```python
x = 15 // 2
print(x)  # Output: 7
```

# Modulus (`%`)

Returns the remainder of a division.

```python
x = 10 % 3
print(x)  # Output: 1
```

# Exponentiation (`**`)

Raises a number to the power of another number.

```python
x = 2 ** 3
print(x)  # Output: 8
```

# Operator Precedence

Python follows standard mathematical precedence:
1. `**` (Exponentiation)
2. `*`, `/`, `//`, `%` (Multiplication, Division, Modulus)
3. `+`, `-` (Addition, Subtraction)

Example:
```python
x = 2 + 3 * 4 ** 2
print(x)  # Output: 50
```

## Summary

- Python supports basic arithmetic operators: `+`, `-`, `*`, `/`, `//`, `%`, `**`
- Operator precedence determines order of operations
- Use parentheses `()` to control precedence when needed


