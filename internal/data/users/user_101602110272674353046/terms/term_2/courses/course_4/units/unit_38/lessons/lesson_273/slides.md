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

- Make a new directory `lesson_3` in the `unit_1` directory and open it in VS Code.
- Create a new file `lesson_3.py` and prepare to run example code in today's lecture slides.

```text
/python
  /unit_1
    /lesson_2
      /lesson_2.py
```

# Introduction to `print()` and `input()` in Python

# The `print()` Function

- Used to display text or output on the screen.
- Example:

```python
print("Hello, world!")
```

- This will print:

```
Hello, world!
```

# Printing Multiple Items

- You can print multiple items using commas:

```python
print("My name is", "Alice")
```

- Output:

```
My name is Alice
```

- Commas add a space automatically.

# The `input()` Function

- Used to get user input.
- Example:

```python
name = input("What is your name? ")
print("Hello,", name)
```

- If the user enters `Bob`, the output will be:

```
What is your name? Bob
Hello, Bob
```

# Storing User Input

- `input()` always returns a string.
- Example:

```python
age = input("How old are you? ")
print("You are", age, "years old.")
```

- Even if the user types a number, `age` is a string.

# Converting Input to a Number

- Use `int()` to convert input to an integer.
- Example:

```python
age = int(input("How old are you? "))
print("Next year, you will be", age + 1)
```

- If the user enters `10`, output will be:

```
How old are you? 10
Next year, you will be 11
```

# Recap

✅ `print()` displays output.
✅ `input()` gets user input.
✅ `input()` returns a string.
✅ Convert input to numbers if needed.

# Practice Time!

Try writing a Python program that:
1. Asks for your favorite color.
2. Prints a message using your input.

```python
color = input("What is your favorite color? ")
print("Wow!", color, "is a great color!")
```

# Assignment 1.3

- [Assignment 1.3](./files/assignment1_1_3.py) Practice with print(), input(), data types and operators
- CodeHS 2.9 and 2.10