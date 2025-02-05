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

- Make a new directory `lesson_5` in the `unit_1` directory and open it in VS Code
- Create a new file `lesson_5.py` and prepare to run example code from the slides.

```text
/python
  /unit_1
    /lesson_5
      /lesson_5.py
```

# Lesson 1.5: Dictionaries & Comprehensions

# What Are Dictionaries?
- **Dictionaries** are collections of key-value pairs.
- Keys must be unique and immutable.
- Defined using curly braces `{}`.
- Example:

```python
student = {"name": "Alice", "age": 20, "grade": "A"}
```

# Dictionary Operations
- Access values: `student["name"]` → `"Alice"`
- Modify values: `student["age"] = 21`
- Add new key-value pair: `student["major"] = "CS"`
- Remove key: `del student["grade"]`
- Get all keys: `student.keys()`
- Get all values: `student.values()`
- Check if key exists: `'name' in student`

# Looping Through Dictionaries
```python
for key, value in student.items():
    print(f"{key}: {value}")
```

# Dictionary Methods
```python
# Merging dictionaries
d1 = {"a": 1, "b": 2}
d2 = {"c": 3, "d": 4}
d1.update(d2)  # {"a": 1, "b": 2, "c": 3, "d": 4}

# Getting value safely
print(student.get("name", "Unknown"))  # Returns 'Alice'
```

# What Are Comprehensions?
- **Comprehensions** provide a concise way to create collections.
- Used for lists, sets, and dictionaries.
- Faster and more readable than loops.

# List Comprehension
```python
numbers = [1, 2, 3, 4, 5]
squares = [x**2 for x in numbers]
print(squares)  # [1, 4, 9, 16, 25]
```

# Dictionary Comprehension
```python
nums = [1, 2, 3, 4]
square_dict = {x: x**2 for x in nums}
print(square_dict)  # {1: 1, 2: 4, 3: 9, 4: 16}
```

# Set Comprehension
```python
nums = [1, 2, 2, 3, 4, 4]
unique_squares = {x**2 for x in nums}
print(unique_squares)  # {1, 4, 9, 16}
```

# Filtering with Comprehensions
```python
numbers = range(10)
even_squares = {x: x**2 for x in numbers if x % 2 == 0}
print(even_squares)  # {0: 0, 2: 4, 4: 16, 6: 36, 8: 64}
```

# Practice Questions
1. What is the difference between lists and dictionaries?
2. How do you safely get a value from a dictionary?
3. Write a dictionary comprehension that maps numbers 1-5 to their cubes.
4. Create a list comprehension that filters out odd numbers from `range(10)`.
5. Convert a dictionary’s keys to uppercase using a comprehension.

# Summary
- **Dictionaries**: Key-value pairs, efficient lookups.
- **Comprehensions**: Concise, efficient way to generate collections.
- Know when to use them for readability and performance!

# Assignment 1.5

