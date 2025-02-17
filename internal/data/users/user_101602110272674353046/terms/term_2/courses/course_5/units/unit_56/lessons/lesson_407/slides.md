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

Make an array of 100 random integers and then use a filter to get all integers less than 10.

```python
import numpy as np
arr = np.random.randint(1, 101, size=(100))
# add code to filter the array here
```

# Session 4: Introduction to pandas

## Exploring DataFrames and Data Manipulation

# Introduction to pandas

Pandas is a powerful Python library for data manipulation and analysis. It provides data structures and functions designed to work efficiently with structured data.

## Key Features

- Handling of missing data
- Size mutability: columns can be inserted and deleted from DataFrame
- Powerful, flexible group by functionality to perform split-apply-combine operations on data sets

# Creating DataFrames

DataFrames are the primary pandas data structure for storing and manipulating two-dimensional data with labeled axes.

## Example: Creating a DataFrame

```python
import pandas as pd
data = {
    'Name': ['John', 'Anna', 'James', 'Linda'],
    'Age': [28, 22, 35, 32],
    'City': ['New York', 'Paris', 'Berlin', 'London']
}
df = pd.DataFrame(data)
print(df)
```

# Importing Data

One of the most common tasks in data science is importing data from external sources.

## Example: Importing from CSV

```python
df = pd.read_csv('data.csv')
print(df.head())  # Displays the first few rows of the DataFrame
```

Writing to a csv is equally easy:

```python
df.to_csv('my_file.csv')
```

# Basic Data Manipulation

Manipulating data is straightforward with pandas, which offers numerous built-in methods for slicing, transforming, and summarizing data.

## Example: Selecting Data

```python
# Select rows where Age is greater than 30
older_than_30 = df[df['Age'] > 30]
print(older_than_30)
```

# Advanced DataFrame Operations

Pandas provides methods for more complex data manipulations and transformations, including merging, grouping, and pivoting data.

## Example: Grouping Data

```python
grouped_data = df.groupby('City')
mean_ages = grouped_data['Age'].mean()
print(mean_ages)
```

# Session Summary

Today, we introduced pandas, learned how to create DataFrames, import data from a CSV file, and perform basic data manipulations. These capabilities are fundamental for any data scientist working with Python.

**Next Session**: We will dive deeper into pandas data manipulation, exploring more complex operations and how to handle larger data sets effectively.
