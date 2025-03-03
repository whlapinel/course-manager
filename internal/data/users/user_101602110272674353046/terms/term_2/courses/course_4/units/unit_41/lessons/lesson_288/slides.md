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

- Phones in bags, bags against the wall
- Create and open `lesson_5` in VS Code

```text
/python
  /unit_4
    /lesson_5
```

# Python Modules: Introduction & Usage

## 📦 What is a Module?

A **module** is a Python file (`.py`) that contains **functions, classes, or variables**.

✅ Helps organize code  
✅ Promotes **reusability**  
✅ Makes programs **modular**  

# 🛠 Creating a Module

1️⃣ Create a new `.py` file, e.g., `math_utils.py`  
2️⃣ Define some functions inside it:

```python
# math_utils.py
def add(a, b):
    return a + b
```

3️⃣ Now, this file can be used in other Python scripts!

# 🔄 Importing a Module

## Method 1: Import the Whole Module
```python
import math_utils
print(math_utils.add(2, 3))  # Output: 5
```

✅ Keeps namespace clear  
❌ Must reference module name when calling functions

# 🎯 Importing Specific Functions

## Method 2: Import a Function Directly
```python
from math_utils import add
print(add(2, 3))  # Output: 5
```

✅ No need to prefix function with module name  
❌ May cause conflicts if function names overlap  

# 🌟 Importing with an Alias

## Method 3: Using `as` to Rename
```python
import math_utils as mu
print(mu.add(2, 3))  # Output: 5
```

✅ Shorter & more readable  
✅ Useful for long module names  

# ⚠️ `import *` (Wildcard Import)

## Method 4: Import Everything
```python
from math_utils import *
print(add(2, 3))  # Output: 5
```
🚨 **Avoid using `*` (wildcard imports)**:
- Can cause **name conflicts**  
- Makes debugging **harder**  
- Hard to know where functions come from  

# 🔄 Module Scope & `__name__`

Inside a module, `__name__` holds its name.  
If a script is **run directly**, `__name__ == "__main__"`.

```python
# script.py
print(__name__)
```

## Running it:
```sh
$ python script.py
__main__
```
✅ Helps create reusable scripts by differentiating **imported modules** vs **executed scripts**.

# 🏗 Built-in & Third-Party Modules

Python has many built-in modules!  
Try these:
```python
import math
print(math.sqrt(16))  # Output: 4.0

import random
print(random.randint(1, 10))  # Random number between 1-10
```
💡 Third-party modules (e.g., `numpy`, `pandas`) must be **installed** using `pip install`.

# 🎯 Recap

✅ **What is a module?** A `.py` file containing functions, classes, or variables.  
✅ **Ways to import:**  
- `import module`  
- `from module import func`  
- `import module as alias`  
✅ **Best practices:** Avoid `import *`, use `__name__ == "__main__"`  

# 🎯 Your Turn! 💡

✅ **Task 1:** Create a module (`greetings.py`) with a function `hello(name)` that prints `"Hello, <name>!"`.  
✅ **Task 2:** Import and call `hello("Richard")` from another script.  
✅ **Task 3:** Try different import methods.
