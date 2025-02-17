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

Create an numpy array consisting of numbers from 1 to 1000 and print the sum.

# Session 2: Advanced numpy Operations

## Indexing, Slicing, Reshaping, and Statistics

# Indexing and Slicing

Indexing and slicing allow you to access parts of arrays and matrices. 

```python
import numpy as np
arr = np.array([1, 2, 3, 4, 5])

# Single element
print(arr[2])

# Slicing
print(arr[1:4])
```

# Reshaping Arrays

Reshaping changes the arrangement of items so that you can increase or decrease the dimensionality.

```python
# Reshape from 1D to 2D
arr = np.arange(10)  # Creates an array with numbers from 0 to 9
reshaped_arr = arr.reshape(2, 5)
print(reshaped_arr)
```

# Introduction to Statistical Functions

numpy provides a wide range of statistical functions that can be applied directly to numpy arrays.

```python
# Statistics
data = np.array([1, 2, 3, 4, 5])

print("Mean:", np.mean(data))
print("Median:", np.median(data))
print("Standard Deviation:", np.std(data))
```

# Applying Functions Across Dimensions

You can apply functions along a specific axis of a multidimensional array.

```python
matrix = np.array([[1, 2, 3], [4, 5, 6]])

# Sum across rows
print("Sum of each column:", np.sum(matrix, axis=0))

# Sum across columns
print("Sum of each row:", np.sum(matrix, axis=1))
```

# Session Summary

Today we delved deeper into numpy's capabilities, learning how to manipulate array shapes and perform basic statistical analysis. These tools are fundamental for analyzing and understanding data in Python.

**Next Session**: We will begin our exploration of pandas, another powerful tool for data manipulation and analysis in Python.
