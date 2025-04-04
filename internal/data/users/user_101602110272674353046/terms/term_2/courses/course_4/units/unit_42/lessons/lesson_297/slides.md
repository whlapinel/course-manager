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
    /lesson_6
```

# String Manipulation in Python

## What are Strings?
- A **string** is a sequence of characters.
- Strings are **immutable** — they cannot be changed after creation.
- Strings can be accessed like lists using indexing and slicing.

```python
s = "Hello, world!"
print(s[0])    # Output: H
print(s[-1])   # Output: !
print(s[0:5])  # Output: Hello
```

# String Methods

## `upper()` and `lower()`

```python
s = "Hello"
print(s.upper())  # Output: HELLO
print(s.lower())  # Output: hello
```

# `strip()`

```python
s = "  Hello  "
print(s.strip())  # Output: "Hello"
```

# `replace()`

```python
s = "Hello, world!"
print(s.replace("world", "Python"))  # Output: Hello, Python!
```

# Splitting and Joining Strings

## `split()` — Converts a string into a list

```python
s = "apple,banana,orange"
fruits = s.split(",")
print(fruits)  # Output: ['apple', 'banana', 'orange']
```

# `join()` — Combines a list into a string

```python
words = ["Python", "is", "fun"]
result = " ".join(words)
print(result)  # Output: Python is fun
```

# Formatting Strings

## `f-strings` (Python 3.6+)

```python
name = "Alice"
age = 25
print(f"My name is {name} and I am {age} years old.")
```

# `format()` method

```python
print("My name is {} and I am {} years old.".format(name, age))
```

# Indexing and Slicing

```python
s = "Python"
print(s[0])    # Output: P
print(s[-1])   # Output: n
print(s[1:4])  # Output: yth
```

- `s[start:end]` → from `start` to `end-1`
- `s[:4]` → from beginning to 4th element
- `s[2:]` → from 2nd element to end

# Looping Through Strings

## Example 1: Character by character

```python
s = "Python"
for c in s:
    print(c)
```

# Example 2: With `enumerate()`

```python
for i, c in enumerate("Python"):
    print(f"Index {i}: {c}")
```

# String Checking

## `in` operator

```python
s = "Python"
print("th" in s)  # Output: True
print("xy" in s)  # Output: False
```

# `startswith()` and `endswith()`

```python
s = "Python"
print(s.startswith("Py"))  # Output: True
print(s.endswith("on"))    # Output: True
```

# Practical Example: Reverse a String

```python
s = "Python"
reversed_s = s[::-1]
print(reversed_s)  # Output: nohtyP
```

- `s[::-1]` → slice with step -1 to reverse the string

# Recap
✅ Strings are immutable  
✅ Access strings using indexing and slicing  
✅ Use `split()` and `join()` to convert between strings and lists  
✅ Format strings with `f-strings`  
✅ Use `in`, `startswith()`, and `endswith()` for checking  
