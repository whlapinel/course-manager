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

- Make a new directory `lesson_2` in the `unit_2` directory and open it in VS Code
- Create a new file `lesson_2.py` and prepare to run example code from the slides.

```text
/python
  /unit_2
    /lesson_2
      /lesson_2.py
```

# Introduction to Recursion

## What is Recursion?

- A function that calls itself
- Used to solve problems by breaking them into smaller subproblems
- Has a **base case** to stop recursion
- Has a **recursive case** to call itself with a smaller input

# Example 1: Counting Down

```python
def countdown(n):
    if n <= 0:  # Base case
        print("Blast off!")
    else:
        print(n)
        countdown(n - 1)  # Recursive case

countdown(5)
```

**Output:**
```text
5
4
3
2
1
Blast off!
```

# Identifying Base and Recursive Cases

## Example: Counting Down

- **Base Case:** `if n <= 0: print("Blast off!")`
- **Recursive Case:** `countdown(n - 1)`

# Example 2: Factorial

```python
def factorial(n):
    if n == 0:  # Base case
        return 1
    else:
        return n * factorial(n - 1)  # Recursive case

print(factorial(5))
```

**Output:**
```text
120
```

# Identifying Base and Recursive Cases

## Example: Factorial

- **Base Case:** `if n == 0: return 1`
- **Recursive Case:** `return n * factorial(n - 1)`

# Example 3: Exploring Nested Directories

```python
import os

def list_files(directory):
    for item in os.listdir(directory):
        path = os.path.join(directory, item)
        if os.path.isdir(path):
            list_files(path)  # Recursive case
        else:
            print(path)

list_files("./some_directory")
```

# Identifying Base and Recursive Cases

## Example: Exploring Nested Directories

- **Base Case:** When there are no more subdirectories to explore
- **Recursive Case:** `list_files(path)` if `path` is a directory

# Example 4: Reverse a String

```python
def reverse_string(s):
    if len(s) == 0:  # Base case
        return ""
    else:
        return s[-1] + reverse_string(s[:-1])  # Recursive case

print(reverse_string("hello"))
```

**Output:**
```text
olleh
```

# Identifying Base and Recursive Cases

## Example: Reverse a String

- **Base Case:** `if len(s) == 0: return ""`
- **Recursive Case:** `return s[-1] + reverse_string(s[:-1])`

# Example 5: Nested List Flattening

```python
def flatten_list(nested_list):
    flat = []
    for item in nested_list:
        if isinstance(item, list):
            flat.extend(flatten_list(item))  # Recursive case
        else:
            flat.append(item)
    return flat

print(flatten_list([1, [2, [3, 4], 5], 6]))
```

**Output:**
```text
[1, 2, 3, 4, 5, 6]
```

# Identifying Base and Recursive Cases

## Example: Nested List Flattening

- **Base Case:** When an item is not a list, append it to `flat`
- **Recursive Case:** `flat.extend(flatten_list(item))` if `item` is a list

# Recursion vs. Iteration

| Recursion | Iteration |
|-----------|----------|
| Uses function calls | Uses loops |
| Can be elegant and concise | Can be more efficient |
| Needs a base case | Uses conditionals |
| Uses more memory (stack calls) | Uses less memory |

# When to Use Recursion?

- When a problem is naturally recursive (e.g., tree traversal, directory searching)
- When dividing a problem into smaller subproblems is easier
- When an iterative solution is too complex

# Recap

- Recursion is a function that calls itself
- Every recursive function needs a **base case**
- Useful for problems like **counting down, directory traversal, string reversal, and list flattening**
- Recursion can be less efficient than iteration
