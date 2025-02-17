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

Create an array of 100 numbers and then find the mean and median using numpy mean and median functions.

# Session 3: Advanced numpy Functionalities

## Deep Dive into Multidimensional Arrays and Advanced Statistics

# Review of Previous Concepts

Here's a quick recap of what we covered in the last session:

- Indexing and slicing numpy arrays
- Reshaping arrays to change their dimensions
- Basic statistical operations: mean, median, and standard deviation

# Multidimensional Array Operations

Multidimensional arrays, or matrices, are crucial in numpy for representing and manipulating data in multiple dimensions.

## Example: Creating a 3D Array

```python
import numpy as np
three_d_array = np.array([[[1, 2], [3, 4]], [[5, 6], [7, 8]]])
print(three_d_array)
```

# Advanced Indexing Techniques

## Boolean Indexing

Use conditions to index arrays, which is useful for filtering data.

```python
data = np.array([1, 2, 3, 4, 5])
filter = data > 3
filtered_data = data[filter]
print(filtered_data)
```

# Advanced Statistical Functions

Numpy provides functions that go beyond basics, which are essential for deeper data analysis.

## Variance and Correlation

Calculate variance to understand data spread and correlation to find relationships between variables.

```python
data = np.random.rand(10)
print("Variance:", np.var(data))
x = np.random.rand(10)
y = 2 * x + np.random.normal(0, 0.1, 10)
print("Correlation coefficient:", np.corrcoef(x, y))
```

# Applying Functions Across Axes

When dealing with multidimensional arrays, you might want to apply functions along a specific axis.

## Summing Across Axes

```python
matrix = np.array([[1, 2, 3], [4, 5, 6]])
print("Sum across columns:", np.sum(matrix, axis=0))
```

# Session Summary

Today, we explored advanced functionalities in numpy, enhancing our ability to handle complex data manipulations and perform detailed statistical analysis.

- Reviewed and expanded on array operations
- Introduced advanced statistical functions for comprehensive data analysis

# **Introduction to Jupyter Notebooks**

# Slide 2: **What is a Jupyter Notebook?**

- Interactive environment for running and writing code
- Mixes code, text, images, and even plots in one place
- Commonly used in data science, machine learning, and research

# Slide 3: **Why Use Jupyter Notebooks?**

- **Immediate feedback**: Run code cells to see results instantly
- **Documentation**: Write explanations directly alongside your code
- **Reproducibility**: Other people can see your exact steps and replicate your analysis

# Slide 4: **Starting Jupyter Notebook**

1. Open your terminal
2. type `pip install jupyter` and press Enter (wait for installation to finish)
3. Type `jupyter notebook` and press Enter
4. This will open a notebook interface in your web browser

# Slide 5: **Navigating the Jupyter Interface**

- **Dashboard**: Lists your files and notebooks
- **Toolbar**: Contains tools like Save, Add Cells, etc.
- **Cells**: Where you write and run code or text

# Slide 6: **Types of Cells**

- **Code Cells**: For writing Python code (useful for Numpy!)
- **Markdown Cells**: For formatted text, headers, bullet points, and explanations

# Slide 7: **Writing and Running Code**

1. Select a code cell
2. Type your code (e.g., `import numpy as np`)
3. Press Shift + Enter to run the cell

# Slide 8: **Example: Simple Numpy Calculation**

- Write a code cell that:

  ```python
  import numpy as np
  data = np.array([1, 2, 3, 4])
  print(data * 2)
  ```

- Press Shift + Enter to see the output

# Slide 9: **Adding Text with Markdown Cells**

- Markdown cells let you add text, headings, lists, and more
- Example:

  ```markdown
  ## My Data Analysis
  This is a **bold** text, and this is an *italic* text.
  ```

# Slide 10: **Basic Shortcuts**

- **Shift + Enter**: Run cell
- **A**: Insert cell above
- **B**: Insert cell below
- **M**: Convert cell to Markdown
- **Y**: Convert cell to Code

# Slide 11: **Practice Exercise**

1. Create a new notebook
2. Import Numpy and create an array of numbers
3. Multiply each element by 3 and print the result
4. Add a Markdown cell to explain what you’re doing

# Slide 12: **Saving and Exporting Notebooks**

- Save your notebook frequently (`Ctrl + S` or click Save icon)
- Export as PDF, HTML, or Python script (File > Download as)

# Slide 13: **Summary**

- Jupyter Notebooks allow you to combine code, text, and visuals
- Great for data exploration and analysis
- Practice using code cells and Markdown for a smooth workflow