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

- Phones in bags, bags against the wall
- Create the following and open it in VS Code

```text
/python
  /unit_5
    /lesson_2
```

# More on Lists / Intro to Tuples

# **List Slicing**
- **What is Slicing?**
  - Extracting a portion of a list.
- **Syntax**:
  ```python
  list[start:end:step]
  ```
- **Examples**:
  - Slice from index 1 to 3:
    ```python
    fruits = ["apple", "banana", "cherry", "date"]
    print(fruits[1:3])  # Output: ["banana", "cherry"]
    ```
  - Slice with a step of 2:
    ```python
    print(fruits[::2])  # Output: ["apple", "cherry"]
    ```

# **List Comprehension**
- **What is List Comprehension?**
  - A concise way to create or modify lists.
- **Basic Syntax**:
  ```python
  [expression for item in iterable]
  ```
- **Example**:  
  Create a list of squares from 0 to 4:
  ```python
  squares = [x**2 for x in range(5)]
  print(squares)  # Output: [0, 1, 4, 9, 16]
  ```

# **Common List Methods**
- **Sort the List**:
  ```python
  fruits = ["banana", "apple", "cherry"]
  fruits.sort()
  print(fruits)  # Output: ["apple", "banana", "cherry"]
  ```
- **Reverse the List**:
  ```python
  fruits.reverse()
  print(fruits)  # Output: ["cherry", "banana", "apple"]
  ```
- **Clear All Items**:
  ```python
  fruits.clear()
  print(fruits)  # Output: []
  ```

# **Nested Lists**
- **Lists inside Lists**:  
  A list can contain other lists.
  ```python
  nested_list = [[1, 2], [3, 4], [5, 6]]
  print(nested_list[0])  # Output: [1, 2]
  print(nested_list[0][1])  # Output: 2
  ```

# Intro to Tuples in Python

---

## What is a Tuple?
- A **tuple** is an **ordered** and **immutable** collection of elements.
- Similar to a list, but:
  - Defined with **parentheses** `()` instead of square brackets `[]`.
  - Cannot be modified after creation.

```python
my_tuple = (1, 2, 3)
print(my_tuple)
# Output: (1, 2, 3)
```

---

## Why Use Tuples?
✅ Tuples are **faster** than lists.  
✅ Useful for data that should **not change**.  
✅ Can be used as **keys** in dictionaries (lists cannot).  

```python
coordinates = (4, 5)
location = {coordinates: "Park"}
print(location)
# Output: {(4, 5): 'Park'}
```

---

## Creating Tuples
- With **parentheses**:
```python
my_tuple = (1, 2, 3)
```
- Without parentheses (tuple packing):
```python
my_tuple = 1, 2, 3
```
- Empty tuple:
```python
empty = ()
```
- Single element tuple (**comma required**):
```python
single = (5,)  # NOT (5)
```

---

## Accessing Tuple Elements
- By **index** (like a list):
```python
my_tuple = (10, 20, 30)
print(my_tuple[0])  # Output: 10
```
- Negative indexing:
```python
print(my_tuple[-1])  # Output: 30
```
- Slicing:
```python
print(my_tuple[1:])  # Output: (20, 30)
```

---

## Tuple Operations
- **Concatenation** (`+`):
```python
a = (1, 2)
b = (3, 4)
c = a + b
print(c)  # Output: (1, 2, 3, 4)
```
- **Repetition** (`*`):
```python
d = (5,) * 3
print(d)  # Output: (5, 5, 5)
```

---

## Tuple Unpacking
- Assign tuple values to variables:
```python
person = ("John", 30)
name, age = person
print(name)  # Output: John
print(age)   # Output: 30
```
- Using `_` for unused values:
```python
x, _, z = (1, 2, 3)
print(x, z)  # Output: 1 3
```

---

## Immutable Nature
- Tuples **cannot be modified** after creation:
```python
t = (1, 2, 3)
t[0] = 10  # Error: TypeError
```
- But you can **reassign**:
```python
t = (10, 20)
print(t)  # Output: (10, 20)
```

---

## Tuple Methods

| Method | Description |
|--------|-------------|
| `count(x)` | Returns the number of occurrences of `x` |
| `index(x)` | Returns the index of the first occurrence of `x` |

Example:
```python
t = (1, 2, 3, 2)
print(t.count(2))  # Output: 2
print(t.index(3))  # Output: 2
```

---

## Tuple vs List

| Feature | Tuple | List |
|---------|-------|------|
| **Mutable** | ❌ No | ✅ Yes |
| **Faster** | ✅ Yes | ❌ No |
| **Hashable** | ✅ Yes | ❌ No |
| **Syntax** | `()` | `[]` |

---

## When to Use Tuples
✔️ Use tuples when data should not change.  
✔️ Use lists when data needs to be modified.  
✔️ Use tuples as dictionary keys or in sets.  

---

## Challenge
Create a tuple containing your name, age, and favorite color.  
Unpack the tuple and print each value.  
Try to modify one value — what happens?  

---

## Summary
✅ Tuples are immutable.  
✅ Created with `()` or comma-separated values.  
✅ Support indexing, slicing, and unpacking.  
✅ Faster than lists and hashable.  
