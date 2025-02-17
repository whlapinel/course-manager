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

# Lesson 1.1

Here is a changed slide

# Warmup

- Create a new directory `unit_1` in your `python` course directory in onedrive.
- Make a new directory `lesson_1` in the `unit_1` directory

```text
/python
  /unit_1
    /lesson_1
```
  
# Agenda

- While loops
  
# Looking ahead

- Unit 1: Python Fundamentals Review officially starts today
- Unit 2 will start next Friday and this will be when we get into the "new" stuff (Python 2 standards)

# While Loops in Python

- Loop until a condition is false

# What is a While Loop?
- A `while` loop repeatedly executes a block of code **as long as a condition is true**.
- Useful when you don't know beforehand how many times the loop should run.

**Syntax:**
```python
while condition:
    # Code block
```

# Basic Example

```python
count = 0
while count < 5:
    print("Count is:", count)
    count += 1
```
**Output:**
```text
Count is: 0
Count is: 1
Count is: 2
Count is: 3
Count is: 4
```

# Infinite Loops
- If the condition never becomes `False`, the loop will run forever.
- Example of an **infinite loop**:

```python
while True:
    print("This will run forever!")
```

*Press Ctrl+C to stop execution.*

# Using `break` to Exit a Loop
- The `break` statement **immediately exits** the loop.

```python
count = 0
while count < 10:
    if count == 5:
        break
    print(count)
    count += 1
```
**Output:**
```text
0
1
2
3
4
```

# Using `continue` to Skip an Iteration
- The `continue` statement **skips the rest of the loop's body** and goes to the next iteration.

```python
count = 0
while count < 5:
    count += 1
    if count == 3:
        continue
    print(count)
```
**Output:**
```text
1
2
4
5
```
# Using `else` with `while`
- The `else` block runs **only if the loop was not terminated with `break`**.

```python
count = 0
while count < 3:
    print(count)
    count += 1
else:
    print("Loop finished without a break!")
```

**Output:**
```text
0
1
2
Loop finished without a break!
```

# Common Mistakes
## 1. Forgetting to Update the Condition
```python
count = 0
while count < 5:
    print(count)  # Infinite loop (count never increases)
```

## 2. Using `=` Instead of `==`
```python
while count = 5:  # SyntaxError: invalid syntax
    print(count)
```

# Summary
✅ Use `while` for loops when the number of iterations is **unknown** beforehand.
✅ Use `break` to **exit** a loop early.
✅ Use `continue` to **skip** an iteration.
✅ Use `else` to execute code **only if the loop completes**.
✅ Avoid **infinite loops** by ensuring the condition will eventually be `False`.

# Practice Questions
1. Write a `while` loop that prints numbers from 10 to 1.
2. Modify the loop to **skip number 5** using `continue`.
3. Modify the loop to **stop at 3** using `break`.

# Assignment 1.1: Practice With While Loops

[Assignment 1.1](./files/assignment1_1.py)
