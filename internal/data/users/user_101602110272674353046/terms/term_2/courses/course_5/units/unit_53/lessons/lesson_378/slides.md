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
