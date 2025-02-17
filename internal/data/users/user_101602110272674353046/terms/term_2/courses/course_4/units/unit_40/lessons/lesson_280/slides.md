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

# Introduction to Lists and using `for` Loops with Lists

## What is a List?
- A **list** is a collection of items stored in a single variable.
- Lists are **ordered**, **mutable**, and can hold different types of data.
- Defined using square brackets `[]`.

**Example:**
```python
fruits = ["apple", "banana", "cherry"]
print(fruits)
```
**Output:**
```text
['apple', 'banana', 'cherry']
```

# Accessing List Elements
- Each item in a list has an **index**, starting from `0`.
- Access elements using square brackets `[]`.

**Example:**
```python
fruits = ["apple", "banana", "cherry"]
print(fruits[0])
```
**Output:**
```text
apple
```

# Iterating Through a List with a `for` Loop
- `for` loops allow us to go through each item in a list.

**Example:**
```python
fruits = ["apple", "banana", "cherry"]
for fruit in fruits:
    print(fruit)
```
**Output:**
```text
apple
banana
cherry
```

# Using `range()` with Lists
- We can use `range()` and `len()` to iterate by index.

**Example:**
```python
fruits = ["apple", "banana", "cherry"]
for i in range(len(fruits)):
    print(i, fruits[i])
```
**Output:**
```text
0 apple
1 banana
2 cherry
```

# Modifying Lists in a `for` Loop
- We can update elements while looping.

**Example:**
```python
numbers = [1, 2, 3, 4]
for i in range(len(numbers)):
    numbers[i] *= 2
print(numbers)
```
**Output:**
```text
[2, 4, 6, 8]
```

# Using `for` with `if` Conditions in Lists
- `for` loops can filter items.

**Example:**
```python
numbers = [10, 15, 20, 25]
for num in numbers:
    if num % 2 == 0:
        print(num, "is even")
```
**Output:**
```text
10 is even
20 is even
```

# Summary
- Lists store multiple values in a single variable.
- Use `for` loops to iterate through list elements.
- Use `range(len(list))` to loop through list indices.
- Lists can be modified inside loops.

# Practice Questions
1. Create a list of 5 colors and print each one using a `for` loop.
2. Iterate through a list of numbers and print only the even ones.
3. Multiply all numbers in a list by 3 using a loop.

# Assignment 3.4

- CodeHS 3.5 & 3.6
- Submit solutions to [problems](./files/assignment_1_3_4.html) in the form of a python file
