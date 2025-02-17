---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Lesson 1.1

# Warmup

- Create a new directory `unit_1` in your `python` course directory in onedrive.
- Make a new directory `lesson_1` in the `unit_1` directory

```text
/python
  /unit_1
    /lesson_1
```

# Variables & Data Types in Python

## What are Variables?  
- A **variable** stores a value in memory.  
- Variables are created when assigned a value:  

```python
x = 10  # Assigning an integer
name = "Alice"  # Assigning a string
pi = 3.14  # Assigning a float
```

# Rules for Variable Names  

✅ Can contain letters, numbers, and underscores (_).  
✅ Must start with a letter or underscore.  
❌ Cannot start with a number.  
❌ Cannot use Python keywords (e.g., `if`, `else`, `while`).  

```python
valid_name = "Python"
_invalid = 42  # Valid but not recommended
3cats = "Error"  # ❌ Invalid: Cannot start with a number
```

# Data Types in Python

Python has several built-in data types:  

| Type | Example | Description |
|------|---------|-------------|
| `int` | `5, -10, 100` | Whole numbers |
| `float` | `3.14, -0.5, 2.0` | Decimal numbers |
| `str` | `"Hello", 'Python'` | Text values |
| `bool` | `True, False` | Boolean (yes/no) |

```python
age = 25      # int
price = 19.99 # float
is_happy = True  # bool
text = "Hello!"  # str
```

# Checking Data Types  

Use `type()` to check a variable's data type.  

```python
x = 10
print(type(x))  # <class 'int'>

y = 3.14
print(type(y))  # <class 'float'>

z = "Python"
print(type(z))  # <class 'str'>
```

# Type Conversion (Casting)  

You can convert between types using:  

```python
x = "10"
y = int(x)  # Converts string "10" to integer 10

a = 3.5
b = str(a)  # Converts float 3.5 to string "3.5"
```

Common conversions:  
- `int(x)` → Converts `x` to an integer  
- `float(x)` → Converts `x` to a float  
- `str(x)` → Converts `x` to a string  

# Quick Quiz 🎯  
1️⃣ What is the type of `x` in `x = 42.0`?  
2️⃣ What will `type("123")` return?  
3️⃣ Convert the string `"99"` into an integer.  

# Summary  
- Variables store values and can change.  
- Data types include **int, float, str, and bool**.  
- Use `type()` to check a variable’s type.  
- Convert data types using `int()`, `float()`, and `str()`.  

# Assignment 1.1

For Canvas Assignment 1.1:

- Complete and submit [Assignment 1.1](./files/assignment1_1.py)
- Complete CodeHS 2.5, 2.6 and mark complete via the text box in Canvas