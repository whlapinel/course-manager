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

# Testing & Unit 2 Review

# **Warmup**
		
# **Agenda**

- Testing in Python (we didn't get to this on Friday)
- Quiz Review
- Unit 2 Review Game

# **Introduction to Testing in Python**  

## Why Test Code?  
- Helps catch errors early  
- Makes debugging easier  
- Ensures code works as expected  

# **Using `assert` for Simple Testing**  

## What is `assert`?  
- An **assertion** checks if a condition is `True`.  
- If the condition is `False`, Python raises an `AssertionError`.  

### Example:  
```python
def add(a, b):
    return a + b

# Test using assert
assert add(2, 3) == 5  # ✅ Passes
assert add(2, 2) == 5  # ❌ Fails (AssertionError)
```

# Why Use `assert` Instead of `print`?  
✅ Stops execution if a test fails  
✅ Easier to find issues in large programs  

# **Introduction to `pytest`**  

## What is `pytest`?  
- A testing framework for Python  
- Runs multiple tests automatically  
- Provides better error messages  

## Installing `pytest`:  
```sh
pip install pytest
```

# **Writing a Simple `pytest` Test**  

## Example:  
Create a file named **`test_math.py`**  
```python
def add(a, b):
    return a + b

def test_add():
    assert add(2, 3) == 5
    assert add(-1, 1) == 0
```

# Running the Test:  
```sh
pytest test_math.py
```
✅ **Passes if all assertions hold**  
❌ **Fails if any assertion is false**  

# **Summary**  

✔ **Use `assert` for quick tests**  
✔ **Use `pytest` for better test automation**  
✔ **Run `pytest` to check multiple test cases**  

# Assignment 2.7

[Assignment 2.7](./files/practice_problems_2_2_7.md)
