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

- Place your phones in your bags and your bags against the wall
- Within your `unit_4` directory, create a new directory `lesson_2`
- Open your `lesson_2` directory in VS Code 
- Open today's lesson on the course website.

# Introduction to CSV Files in Python

## What is a CSV File?

- **CSV** stands for **Comma-Separated Values**.
- A plain text file that stores tabular data.
- Each line represents a row, and values are separated by commas.
- Example:

```csv
Name,Age,Grade
Alice,14,A
Bob,15,B
Charlie,13,A
```

# Why Use CSV Files?

- Simple format, easy to read and write.
- Compatible with spreadsheets and databases.
- Commonly used for data exchange between different applications.

# Reading a CSV File in Python

Using the built-in `csv` module:

```python
import csv

with open("students.csv", mode="r") as file:
    reader = csv.reader(file)
    for row in reader:
        print(row)  # Each row is a list
```

# Writing a CSV File

Using `csv.writer`:

```python
import csv

data = [["Name", "Age", "Grade"],
        ["Alice", 14, "A"],
        ["Bob", 15, "B"],
        ["Charlie", 13, "A"]]

with open("students.csv", mode="w", newline="") as file:
    writer = csv.writer(file)
    writer.writerows(data)
```

# Reading a CSV File as Dictionaries

Using `csv.DictReader`:

```python
import csv

with open("students.csv", mode="r") as file:
    reader = csv.DictReader(file)
    for row in reader:
        print(row["Name"], row["Age"], row["Grade"])
```

# Writing a CSV File from Dictionaries

Using `csv.DictWriter`:

```python
import csv

data = [
    {"Name": "Alice", "Age": 14, "Grade": "A"},
    {"Name": "Bob", "Age": 15, "Grade": "B"},
    {"Name": "Charlie", "Age": 13, "Grade": "A"}
]

with open("students.csv", mode="w", newline="") as file:
    fieldnames = ["Name", "Age", "Grade"]
    writer = csv.DictWriter(file, fieldnames=fieldnames)
    writer.writeheader()
    writer.writerows(data)
```

# Using Pandas to Work with CSV Files

## Reading CSV into a DataFrame

```python
import pandas as pd

df = pd.read_csv("students.csv")
print(df.head())
```

## Writing a DataFrame to CSV

```python
import pandas as pd

data = {"Name": ["Alice", "Bob", "Charlie"], "Age": [14, 15, 13], "Grade": ["A", "B", "A"]}
df = pd.DataFrame(data)
df.to_csv("students.csv", index=False)
```

# Summary

- CSV files store tabular data in plain text.
- Use Python’s `csv` module to read and write CSV files.
- `csv.reader` and `csv.writer` work with lists.
- `csv.DictReader` and `csv.DictWriter` work with dictionaries.
- `pandas` provides a convenient way to handle CSV files in data analysis.

# Assignment 4.2

- [Practice problems](./files/assignment_2_4_2.html)
