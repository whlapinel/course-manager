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

- Make a list of strings
- iterate through each item in the list using a for loop, printing each item.

# Agenda

- Certification exam preparation
- Overview of 2nd quarter
- Unit 6 Lesson 1

# Certification exam preparation

- New module posted: Certification Exam Preparation
- Python Institute Courses
  - Consists of 4 modules, with one quiz and one test each
  - And one summary exam
- 2 corresponding assignments per week
  - one prepare (for the quiz)
  - one rehearse (for the test)

This is in addition to other assignments

# Introduction to Python in Data Science

# What is Data Science?

Data Science is an interdisciplinary field that uses scientific methods, processes, algorithms, and systems to extract knowledge and insights from structured and unstructured data.

# Why Data Science?

- **Decision Making**: Powers informed decision-making with quantitative evidence.
- **Automation and Analytics**: Enhances automation and analytics with predictive capabilities.
- **Innovation**: Drives innovation in industries through pattern detection and actionable insights.

# The Role of Python in Data Science

Python is a popular programming language for data science due to:

- Its simplicity and readability, making it accessible to beginners.
- A rich ecosystem of libraries (pandas, numpy, matplotlib, scikit-learn) tailored for data analysis, machine learning, and data visualization.
- Strong community support and vast resources for learning and problem-solving.

# Key Python Libraries for Data Science

- **numpy**: Fundamental package for numerical computation.
- **pandas**: Essential for data manipulation and analysis.
- **matplotlib**: Primary plotting library.
- **scikit-learn**: Simple and efficient tools for predictive data analysis.

# Summary

Today, we've introduced the vast field of data science and where Python stands within it. In the next sessions, we will dive deeper into Python libraries that will aid us in performing effective data analysis.

# Introduction to numpy

**numpy** is one of the fundamental packages for numerical computing in Python. It provides an efficient interface to store and operate on dense data buffers.

# Why numpy?

- **Performance**: numpy arrays are stored more efficiently than Python lists and allow mathematical operations to be vectorized.
- **Functionality**: numpy provides a high-level syntax that is easy to understand and use, especially for linear algebra and matrix operations.

# Basics of numpy Arrays

```python
import numpy as np

# Create a numpy array from a Python list
arr = np.array([1, 2, 3, 4, 5])
print(arr)
```

# Array Operations

- **Element-wise operations**: Perform operations on each element individually
- **Aggregation**: Compute summaries such as sum, mean, max, min

```python
# Element-wise operations
print(arr + 2)

# Aggregation
print(np.sum(arr))
```

# Array Operations (cont'd)

- **Zeros Array**: Create an array filled with zeros.

```python
import numpy as np
zeros_array = np.zeros(10)
```

- **Array of Consecutive Numbers**:

```python
one_to_ten_array = np.arange(1, 11)
```

- **Array from Python List**:

```python
python_list = [5, 10, 15, 20, 25]
numpy_array_from_list = np.array(python_list)
```

# Basic Array Operations

- **Element-wise Multiplication**:

```python
multiplied_array = numpy_array_from_list * 2
```

- **Generating Random Numbers**:

```python
random_numbers = np.random.random(5)
```

# Basic Statistical Operations

- **Sum of Array Elements**:

```python
sum_of_elements = np.sum(one_to_ten_array)
```

- **Finding the Maximum Value**:
  
```python
max_number = np.max(random_numbers)
```

# Session Summary

Today, we introduced numpy and explored various methods to create arrays and perform basic operations. These tools are essential for any data analysis tasks we'll encounter in data science.

**Next Session**: We will dive deeper into more complex numpy functionalities and start looking at real data manipulations.
