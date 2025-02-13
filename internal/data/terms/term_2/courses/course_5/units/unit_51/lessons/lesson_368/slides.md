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

# Warmup

- Go to [VS Code Keyboard Shortcuts](https://code.visualstudio.com/shortcuts/keyboard-shortcuts-windows.pdf)
  - Identify 3 cool keyboard shortcuts and try them out
  - Be prepared to share your favorite!

# Introduction to Python Modules

# What is a Module?
- A module is a file containing Python code (functions, classes, variables).
- It allows code reusability and organization.
- Modules can be built-in or user-defined.

# Importing a Module
- Use the `import` statement to load a module.
- Example:
  ```python
  import math
  print(math.sqrt(16))  # Output: 4.0
  ```

# Importing Specific Functions
- Use `from ... import ...` to import only what you need.
- Example:
  ```python
  from math import sqrt
  print(sqrt(25))  # Output: 5.0
  ```

# Importing with an Alias
- Use `as` to rename a module for convenience.
- Example:
  ```python
  import numpy as np
  array = np.array([1, 2, 3])
  ```

# Importing Everything
- Use `from module import *` to import all functions and variables (not recommended).
- Example:
  ```python
  from math import *
  print(sin(0), cos(0))
  ```
- This can cause naming conflicts.

# Creating Your Own Module
- Create a `.py` file with functions and variables.
- Example (`mymodule.py`):
  ```python
  def greet(name):
      return f"Hello, {name}!"
  ```
- Use it in another script:
  ```python
  import mymodule
  print(mymodule.greet("Alice"))
  ```

# Built-in Modules: datetime
- The `datetime` module helps manage dates and times.
- Example:
  ```python
  import datetime
  now = datetime.datetime.now()
  print("Current date and time:", now)
  ```

# Extracting Date & Time
- Get specific parts of a date:
  ```python
  today = datetime.date.today()
  print(today.year, today.month, today.day)
  ```
- Formatting dates:
  ```python
  formatted = now.strftime("%Y-%m-%d %H:%M:%S")
  print(formatted)
  ```

# Using `__name__ == "__main__"`
- In Python, `__name__` is a special variable that stores the module name.
- When a script is imported, `__name__` is set to the module name.
- When a script is run directly, `__name__` is set to `"__main__"`.
- Example:
  ```python
  def greet():
      print("Hello from the script!")
  
  if __name__ == "__main__":
      greet()
  ```
- This ensures code runs only when executed as a script, not when imported as a module.

# Why Use `__name__ == "__main__"`?
- Helps distinguish between running a script directly vs. importing it.
- Prevents unintended execution when a module is imported.
- Encourages modular code and reusability.

# Example: Running vs. Importing
**my_script.py:**
  ```python
  def message():
      print("This is a message.")
  
  if __name__ == "__main__":
      message()
  ```

**another_script.py:**
  ```python
  import my_script
  ```
- When `my_script.py` runs directly, it prints the message.
- When imported, it does not print anything automatically.

# Conclusion
- Modules help organize and reuse code.
- Import modules in different ways.
- Create your own module for custom functionality.
- The `datetime` module is useful for handling dates and times.
- `__name__ == "__main__"` helps control script execution.

# Assignment 2.4

- Complete [practice problems](./files/assignment_2_2_4.html) and submit as python file
- Enroll in [Python Essentials 2 Course](https://pythoninstitute.org/python-essentials-2);
- Use your personal email address
- When you create your account make sure your name matches Powerschool
- Module 1 Quiz will be due end of class Friday.
