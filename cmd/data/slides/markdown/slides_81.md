---
marp: true
---

<!-- headingDivider: 1 -->

# Warmup

Write a function `abbreviated` that accepts a string and returns a new string that consists of the last 5 characters.

## Example

```python
modified_string = abbreviated('my favorite class is Python')
print(modified_string) # ython
```

# Agenda

## Today

- Begin Unit 4 (Standards 3.01 through 3.04)
- (U4 Test is October 8th)
- Lesson on File operations
- Assignment: Assignment 4.1

# Glance-ahead

## This week

- 3.01 Build programs that perform file input and output operations.
- 3.02 Develop programs using built-in functions to open/close files.

## Next week

- 3.03 Construct programs using appropriate Python libraries.
- 3.04 Build programs using the sqlite3 and/or the sqlalchemy libraries.

# Announcements

- Certification exam permission slips
- Github repo is now private so that I can put solutions and assessments in the same folders. I will post all materials in Canvas from now on and will try to go back and add them as well.

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

# CSV files and Python

CSV stands for "comma separated values". It is a very common format for files containing data.

## Writing to CSV file

```python
import csv

data = [
    ["Name", "Age", "City"],
    ["Alice", 30, "New York"],
    ["Bob", 25, "San Francisco"]
]

with open('example.csv', mode='w', newline='') as file:
    csv_writer = csv.writer(file)
    csv_writer.writerows(data)  # Write all rows at once
```

Run this code and take a look at the file created in your directory. It's literally just comma-separated values, you could make one yourself using a simple word processor like notepad!

# Reading from CSV file

```python
import csv

with open('example.csv', mode='r') as file:
    csv_reader = csv.reader(file)
    for row in csv_reader:
        print(row)  # Each row is a list of strings
```

**Task:** Read each line from a file and print it

# Exercise: Create and Read a File

1. Write a Python script that:
   - Opens a file named `students.txt` in write mode.
   - Writes 5 student names to the file (one per line).

2. Then:
   - Open `students.txt` in read mode.
   - Print each student name from the file.

# Solution

**Code:**

```python
# Step 1: Writing to the file
with open("students.txt", "w") as file:
    file.writelines(["Alice\n", "Bob\n", "Charlie\n", "Diana\n", "Eve\n"])

# Step 2: Reading from the file
with open("students.txt", "r") as file:
    for line in file:
        print(line.strip())
```

# Assignment: 4.1

Due today, if you don't finish in class, do for homework
