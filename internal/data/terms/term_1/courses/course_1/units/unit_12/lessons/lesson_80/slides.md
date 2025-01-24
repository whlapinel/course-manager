---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Lesson 10.9 Slides

# Review for Python I Final Exam (Teacher-Made)

# **Warm-Up Activity (10 minutes)**
Ask students to complete the following:
1. Write a program that asks the user for their name and age and then prints a message:
   - Example: `"Hello, John! You are 25 years old."`

Example Code:
```python
name = input("Enter your name: ")
age = input("Enter your age: ")
print(f"Hello, {name}! You are {age} years old.")
```

# Functions Refresher (15 minutes)**

**Why Use Functions?**
- Functions help organize code into reusable, modular blocks.
- They reduce repetition and make programs easier to debug.

# **Example**:
```python
def greet_user(name, age):
    return f"Hello, {name}! You are {age} years old."

name = input("Enter your name: ")
age = input("Enter your age: ")
print(greet_user(name, age))
```

# **Activity**:
- Write a function `add_numbers(a, b)` that takes two numbers as arguments and returns their sum.
- Ask students to call the function with different inputs and print the result.

# **Error Handling Basics**

**Why Handle Errors?**
- Prevent programs from crashing due to invalid user input or unexpected issues.

**Key Syntax**:
```python
try:
    num = int(input("Enter a number: "))
    print(f"You entered {num}")
except ValueError:
    print("That is not a valid number.")
```

# **Activity**:
- Write a program that asks the user to input two numbers and calculates their division.
- Use `try` and `except` to handle the following errors:
  1. Non-numeric input.
  2. Division by zero.

# **Using Basic Data Structures**

**Working with Lists**:
- Lists are used to store multiple values in one variable.
- Example:
```python
expenses = [10, 20, 30]
print(f"Total expenses: {sum(expenses)}")
print(f"Average expense: {sum(expenses) / len(expenses)}")
```

# **Activity**:
- Write a program that:
  1. Creates a list of 5 expenses entered by the user.
  2. Prints the total and average of the expenses.

# **Challenge Assignment**

**Assignment Objective**:
Create a small program that:
1. Accepts multiple expenses from the user.
2. Validates that all inputs are numbers.
3. Uses a function to calculate the total and average of the expenses.
4. Handles invalid input gracefully.

# **Challenge Instructions**:

1. **Write a function** `calculate_expenses(expenses)` that:
   - Accepts a list of numbers.
   - Returns the total and average.
2. Prompt the user to input 5 expenses.
3. Use `try` and `except` to validate input.
4. If all inputs are valid, call the function and display the results.

# **Reflection Questions**
1. What happens if the user enters invalid input? How is it handled?
2. Why is it important to use functions for calculations like total and average?
3. How would you modify the program to handle a different number of expenses?

# **What’s Next?**

- **Day 2**: Introduce file handling to save expenses to a file.
- **Day 3**: Combine all concepts into a project.
