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

- Make a new directory `lesson_4` in the `unit_1` directory and open it in VS Code
- Create a new file `lesson_4.py` and prepare to run example code from the slides.

```text
/python
  /unit_1
    /lesson_4
      /lesson_4.py
```
  
# Lesson 1.4: Lists & Tuples

# What Are Lists?
- **Lists** are ordered, mutable collections.
- Defined using square brackets: `[]`.
- Can store different data types.
- Example:

```python
fruits = ["apple", "banana", "cherry"]
```

# List Operations
- Access elements: `fruits[0]` → `"apple"`
- Modify elements: `fruits[1] = "blueberry"`
- Add elements: `fruits.append("orange")`
- Remove elements: `fruits.remove("banana")`
- Get length: `len(fruits)`

# What Are Tuples?
- **Tuples** are ordered, immutable collections.
- Defined using parentheses: `()`.
- Example:

```python
colors = ("red", "green", "blue")
```

# Tuple Operations
- Access elements: `colors[1]` → `"green"`
- Cannot modify: `colors[1] = "yellow"` ❌ (Error)
- Tuple unpacking:

```python
r, g, b = colors
print(g)  # "green"
```

# Lists vs. Tuples

| Feature  | List  | Tuple  |
|----------|------|-------|
| Mutable? | ✅ Yes | ❌ No |
| Ordered? | ✅ Yes | ✅ Yes |
| Syntax   | `[]`  | `()`  |
| Performance | Slower | Faster |

# When to Use Lists vs. Tuples
- Use **lists** when data needs to change.
- Use **tuples** for fixed, read-only data.
- Tuples are faster and use less memory.

# Common List Methods
```python
nums = [3, 1, 4, 1, 5]
nums.sort()  # [1, 1, 3, 4, 5]
nums.reverse()  # [5, 4, 3, 1, 1]
nums.pop()  # Removes last element
```

# Converting Between Lists & Tuples
```python
# List to tuple
tuple_version = tuple([1, 2, 3])

# Tuple to list
list_version = list((4, 5, 6))
```

# Practice Questions
1. What’s the difference between lists and tuples?
2. How do you access the second element in a list?
3. Can you change a tuple’s values? Why or why not?
4. What method adds an element to a list?
5. Convert this list to a tuple: `numbers = [10, 20, 30]`

# Summary
- **Lists**: Mutable, ordered, uses `[]`
- **Tuples**: Immutable, ordered, uses `()`
- Lists are flexible, tuples are memory-efficient
- Know when to use each!
