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
    /lesson_4
```
# Housekeeping

- Permission forms

# Python Data Structures: Sets

## What is a Set?
- A set is an **unordered** collection of **unique** elements.  
- Defined using `{}` or `set()`:
```python
# Creating a set
s = {1, 2, 3}
print(s)  # Output: {1, 2, 3}

# Creating an empty set
empty_set = set()
```

# Why Use Sets?
✅ Fast membership testing (`in`)  
✅ Automatically removes duplicates  
✅ Useful for mathematical operations (union, intersection, difference)  

# Basic Set Operations
- **Adding elements**:
```python
s = {1, 2, 3}
s.add(4)
print(s)  # Output: {1, 2, 3, 4}
```

# Basic Set Operations (cont'd)
- **Removing elements**:
```python
s.remove(2)  # Raises error if element not found
s.discard(2) # No error if element not found
```

# Basic Set Operations (cont'd)
# Set Operations
- **Union (`|`)** – Combine sets:
```python
a = {1, 2, 3}
b = {3, 4, 5}
print(a | b)  # Output: {1, 2, 3, 4, 5}
```

- **Intersection (`&`)** – Common elements:
```python
print(a & b)  # Output: {3}
```

- **Difference (`-`)** – Elements in one set but not the other:
```python
print(a - b)  # Output: {1, 2}
```

- **Symmetric difference (`^`)** – Elements in either set, but not both:
```python
print(a ^ b)  # Output: {1, 2, 4, 5}
```

# Membership Testing
✅ Fast lookup with `in`:
```python
s = {1, 2, 3}
print(2 in s)  # Output: True
print(4 in s)  # Output: False
```

# Set Methods
- **.update()** – Add elements from another set:
```python
a = {1, 2}
b = {3, 4}
a.update(b)
print(a)  # Output: {1, 2, 3, 4}
```

# Set Methods (cont'd)
- **.clear()** – Remove all elements:
```python
s = {1, 2, 3}
s.clear()
print(s)  # Output: set()
```

# Set Methods (cont'd)
- **.copy()** – Return a shallow copy:
```python
s = {1, 2, 3}
copy = s.copy()
```

# Removing Duplicates with Sets
- Converting a list to a set removes duplicates:
```python
items = [1, 2, 2, 3, 4, 4, 5]
unique_items = set(items)
print(unique_items)  # Output: {1, 2, 3, 4, 5}
```

# Frozen Sets
- **Immutable sets**:
```python
fs = frozenset([1, 2, 3])
fs.add(4)  # Error – frozenset is immutable!
```

# Exercise
✅ Write a function `unique_elements(lst)` that takes a list and returns a set of unique elements.  

✅ Write a function `common_elements(set1, set2)` that returns a set of common elements between two sets.  

✅ Write a function `set_difference(set1, set2)` that returns elements in set1 but not in set2.  

# Summary
✅ Sets store unique elements.  
✅ Sets allow fast lookup and support union, intersection, and difference operations.  
✅ Sets are mutable; frozen sets are immutable.  
✅ Use sets to remove duplicates from lists.  
