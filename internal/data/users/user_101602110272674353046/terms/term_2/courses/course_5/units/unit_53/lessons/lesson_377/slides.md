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

1. Put your phone in your bag and put your bag against the wall
2. Create a new directory `unit_4` in your `python` directory
  - (We are skipping Unit 3)
3. Create a new directory `lesson_1` in your `unit_4` directory
```text
/python
  /unit_4
    /lesson_1
```

# Agenda

## Today

- Begin Unit 4 (Standards 3.01 through 3.04)
- Lesson on File operations
- Assignment: Assignment 4.1

# Glance-ahead

## This week

- 3.01 Build programs that perform file input and output operations.
- 3.02 Develop programs using built-in functions to open/close files.

## Next week

- 3.03 Construct programs using appropriate Python libraries.
- 3.04 Build programs using the sqlite3 and/or the sqlalchemy libraries.

# Reading and Writing to Files in Python

Get familiar with handling file operations: creating, reading, writing, and appending files.

# Opening and Writing to a File

In Python, you use the `open()` function to open a file. To write to a file, use the `'w'` mode.

**Code:**

```python
# Open a file in write mode
file = open("example.txt", "w")

# Write to the file
file.write("Hello, world!\n")
file.write("This is a test.")

# Close the file
file.close()
```

**Note:** This will **overwrite** any existing content in `example.txt`.

# Reading from a File

To read from a file, use the `'r'` mode.

**Code:**

```python
# Open the file in read mode
file = open("example.txt", "r")

# Read the content of the file
content = file.read()

# Print the content
print(content)

# Close the file
file.close()
```

**Note:** Always close the file after you're done!

# Using the `with` Statement

The `with` statement automatically handles file closing for you.

**Code:**

```python
# Using 'with' to open the file
with open("example.txt", "r") as file:
    content = file.read()
    print(content)
```

**Advantage:** No need to manually call `file.close()`; Python does it for you.

# Writing Multiple Lines to a File

You can write multiple lines to a file by looping or using a list.

**Code:**

```python
lines = ["Line 1\n", "Line 2\n", "Line 3\n"]

# Open the file in write mode
with open("example.txt", "w") as file:
    file.writelines(lines)
```

**Task:** Write a list of 3 favorite hobbies to a file.

# Appending to a File

To add content to an existing file without overwriting, use the `'a'` (append) mode.

**Code:**

```python
# Open the file in append mode
with open("example.txt", "a") as file:
    file.write("This is an additional line.\n")
```

**Task:** Add a line to an existing file, then read and print its content.

# Reading a File Line by Line

Use a loop to read a file line by line.

**Code:**

```python
# Open the file in read mode
with open("example.txt", "r") as file:
    for line in file:
        print(line.strip())  # Strip removes the newline character
```

# Assignment: 4.1

- [Python problems](./files/assignment4_1.py) (due today)
- Python Essentials 2 Module 2 Quiz (due Friday)

