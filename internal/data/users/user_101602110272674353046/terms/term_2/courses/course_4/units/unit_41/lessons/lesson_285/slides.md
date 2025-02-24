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
- Create a new directory `unit_4` and within that, another directory `lesson_1`
- Open your `lesson_1` directory in VS Code 
- Open today's lesson on the course website.

# Introduction to Functions in Python  

- What is a function and why use them?  
- Anatomy of a Python function  
- Defining and calling functions  
- Parameters vs. Arguments  
- Return values  
- Scope: Local vs. Global  
- Best practices & exercises

# **What is a Function?**  
- A **function** is a reusable block of code that performs a specific task.  
- Helps in organizing code and reducing repetition.  
- Example: Instead of writing the same code over and over, you define it once and call it when needed.

# **Why Use Functions?**  
- **Reusability:** Write once, use many times.  
- **Modularity:** Break code into manageable pieces.  
- **Maintainability:** Easier to update and debug.  
- **Clarity:** Improves readability by abstracting complex logic.

# **Anatomy of a Python Function**  
- **Definition:** Use the `def` keyword.  
- **Syntax:**  
  ```python
  def function_name(parameters):
      """Optional docstring explaining the function."""
      # Code block
      return value  # optional
  ```
- **Components:**  
  - **Name:** Identifier for the function.  
  - **Parameters:** Variables passed to the function.  
  - **Body:** The code that executes when the function is called.  
  - **Return Statement:** Outputs a value (optional).

# **Defining and Calling a Function**  
- **Example:**
  ```python
  def greet(name):
      """Prints a greeting to the given name."""
      print("Hello, " + name + "!")
  
  # Calling the function
  greet("Alice")
  ```
- **Explanation:**  
  - The function `greet` is defined to take one parameter (`name`).  
  - When called with `"Alice"`, it prints: *Hello, Alice!*

# **Parameters vs. Arguments**  
- **Parameters:**  
  - Variables in the function definition.  
  - They act as placeholders for the values the function will use.
- **Arguments:**  
  - Actual values passed to the function when calling it.
- **Example:**  
  - In `def greet(name):`, `name` is a parameter.  
  - In `greet("Alice")`, `"Alice"` is the argument.

# **Return Values**  
- **Purpose:** Functions can send back a result using the `return` statement.  
- **Example:**
  ```python
  def add(a, b):
      return a + b
  
  result = add(3, 4)  # result will be 7
  ```
- **Key Points:**  
  - Not all functions need to return a value.  
  - If no return statement is provided, the function returns `None` by default.

# **Scope: Local vs. Global Variables**  
- **Local Variables:**  
  - Declared inside a function.  
  - Accessible only within that function.
- **Global Variables:**  
  - Declared outside all functions.  
  - Accessible throughout the program.

# Scope continued
- **Example:**
  ```python
  x = 10  # Global variable
  
  def example():
      y = 5  # Local variable
      print(x, y)
  
  example()  # Outputs: 10 5
  ```

# **Best Practices & Tips**  
- **Naming:**  
  - Use clear and descriptive names (e.g., `calculate_total` instead of `ct`).  
- **Documentation:**  
  - Include docstrings to explain what the function does.  
- **Modularity:**  
  - Keep functions focused on a single task.  
- **Testing:**  
  - Test functions individually to ensure they work as expected.

# **Practice Exercise**  
- **Task:** Write a function that takes two numbers as input and returns their product.  
- **Guidelines:**  
  - Define a function `multiply(a, b)`.  
  - Test the function with different sets of numbers.
- **Discussion:**  
  - How might this function be used in a larger program?

# Assignment 4.1

- [Practice problems](./files/assignment_4_1_1.html)
