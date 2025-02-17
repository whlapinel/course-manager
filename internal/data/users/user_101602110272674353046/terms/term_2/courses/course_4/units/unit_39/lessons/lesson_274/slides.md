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

- Make a new unit folder `unit_2` in your `python` directory.
- Make a new directory `lesson_1` in the `unit_2` directory and open it in VS Code.
- Create a new file `lesson_1.py` and prepare to run example code in today's lecture slides.

```text
/python
  /unit_0
  /unit_1
  /unit_2
    /lesson_1
      /lesson_1.py
```

# Introduction to Boolean Expressions

- **Boolean expressions** evaluate to either `True` or `False`.
- They are fundamental in decision-making in Python.
- Used in **conditionals**, **loops**, and **logic operations**.

# Boolean Values

- Python has two Boolean values:
  - `True`
  - `False`
- These values are of type `bool`.

```python
print(type(True))  # <class 'bool'>
print(type(False)) # <class 'bool'>
```

# Comparison Operators

- Used to compare values and return `True` or `False`.

| Operator | Meaning | Example |
|----------|---------|---------|
| `==` | Equal to | `5 == 5` → `True` |
| `!=` | Not equal to | `5 != 3` → `True` |
| `>` | Greater than | `10 > 3` → `True` |
| `<` | Less than | `2 < 5` → `True` |
| `>=` | Greater than or equal to | `5 >= 5` → `True` |
| `<=` | Less than or equal to | `4 <= 3` → `False` |

# Logical Operators

- Combine Boolean expressions using:
  - `and` → True if both conditions are true
  - `or` → True if at least one condition is true
  - `not` → Reverses the Boolean value

```python
x = 10
y = 5

print(x > 5 and y < 10)  # True
print(x < 5 or y > 10)   # False
print(not (x > y))       # False
```

# Boolean in Conditionals

- Used in `if` statements to control program flow.

```python
age = 18

if age >= 18:
    print("You are an adult.")
else:
    print("You are a minor.")
```

# Boolean in Loops

- Used to control loop execution.

```python
x = 5
while x > 0:
    print(x)
    x -= 1  # Decrease x
```

# Truthy & Falsy Values

- Some values behave as `True` or `False` in Boolean expressions.

| Truthy Values | Falsy Values |
|--------------|-------------|
| Non-zero numbers | `0` |
| Non-empty strings | `""` (empty string) |
| Non-empty lists/tuples/dicts | `[]`, `{}` (empty collections) |

```python
if []:
    print("This is True")  # Won't run
else:
    print("This is False") # Will run
```

# Practice Questions

1. What will this code print?
   ```python
   print(3 > 2 and 5 < 10)
   ```
   - A) `True`
   - B) `False`

2. Which of the following is **Falsy**?
   - A) `"hello"`
   - B) `0`
   - C) `[1, 2, 3]`
   - D) `None`

# Summary

- **Boolean expressions** evaluate to `True` or `False`.
- **Comparison operators**: `==`, `!=`, `<`, `>`, `<=`, `>=`
- **Logical operators**: `and`, `or`, `not`
- **Truthy & Falsy values** affect conditions.
- Used in **if statements** and **loops**.

# Assignment 2.1: Practice with Boolean Expressions

- [Assignment 2.1](./files/assignment1_2_1.py)
- Submit the Python file in Canvas Assignment 2.1
- Complete CodeHS Assignments 2.11 and 2.12 and mark complete in Canvas Assignment 2.1


