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

# Lesson 0.3

# Warmup

Create a folder `lesson_3` and open it in VS Code

Write a function `show_name` that simply prints your name

## **Function call**
```python
show_name()
```

## **Example output**
```text
Mr. Lapinel
```

# Review: For Loops in Python

- Yesterday, we reviewed `if` statements.
- Today we'll talk about `for` loops.

# What is a `for` loop?
- A `for` loop is used to iterate over a sequence (e.g., list, tuple, string, range).
- It allows execution of a block of code multiple times.

Example:
```python
for i in range(5):
    print(i)
```
Output:
```text
0
1
2
3
4
```
# Looping Through a List
```python
fruits = ["apple", "banana", "cherry"]
for fruit in fruits:
    print(fruit)
```
Output:
```text
apple
banana
cherry
```

# Using `range()`
- `range(start, stop, step)` generates numbers from `start` to `stop-1`.

Example:
```python
for num in range(1, 10, 2):
    print(num)
```
Output:
```text
1
3
5
7
9
```
# Iterating Over a String
```python
word = "Python"
for letter in word:
    print(letter)
```
Output:
```text
P
y
t
h
o
n
```

# Nested `for` Loops
```python
for i in range(3):
    for j in range(2):
        print(f"i={i}, j={j}")
```
Output:
```text
i=0, j=0
i=0, j=1
i=1, j=0
i=1, j=1
i=2, j=0
i=2, j=1
```

# Looping with `enumerate()`
```python
names = ["Alice", "Bob", "Charlie"]
for index, name in enumerate(names):
    print(index, name)
```
Output:
```text
0 Alice
1 Bob
2 Charlie
```
# Looping with `zip()`
```python
names = ["Alice", "Bob"]
ages = [25, 30]
for name, age in zip(names, ages):
    print(f"{name} is {age} years old")
```
Output:
```text
Alice is 25 years old
Bob is 30 years old
```

# The `break` Statement
- Used to exit a loop early.
```python
for num in range(10):
    if num == 5:
        break
    print(num)
```
Output:
```text
0
1
2
3
4
```
# The `continue` Statement
- Skips the rest of the loop for the current iteration.
```python
for num in range(5):
    if num == 2:
        continue
    print(num)
```
Output:
```text
0
1
3
4
```
# The `else` Clause in Loops
- Runs when the loop finishes without encountering a `break`.
```python
for num in range(3):
    print(num)
else:
    print("Loop finished!")
```
Output:
```text
0
1
2
Loop finished!
```
# Summary
- `for` loops iterate over sequences.
- Use `range()` for numeric loops.
- Use `enumerate()` for indexing.
- Use `zip()` for parallel iteration.
- Control loops with `break` and `continue`.

# Agenda

Work on problems in `practice.py` (new file) together in class.

[practice.py](./files/practice.py)
