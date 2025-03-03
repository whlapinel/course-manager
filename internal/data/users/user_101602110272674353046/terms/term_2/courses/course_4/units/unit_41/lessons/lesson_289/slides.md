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
- Create and open `lesson_6` in VS Code

```text
/python
  /unit_4
    /lesson_6
```

# Introduction to Python Variable Scope

## What is Variable Scope?
- **Scope** refers to where a variable can be accessed in a program.
- Python follows a hierarchy for determining variable scope.
- Knowing scope helps avoid errors and unintended behavior.

# Types of Scope in Python
1. **Local Scope** – Inside a function.
2. **Enclosing Scope** – In nested functions.
3. **Global Scope** – Defined outside functions.
4. **Built-in Scope** – Provided by Python.

This is known as the **LEGB Rule**.

# Understanding Local Scope
- Variables declared inside a function exist only within that function.
- They cannot be accessed outside.

```python
def greet():
    message = "Hello!"  # Local variable
    print(message)

greet()
print(message)  # Raises NameError
```

# Shadowing: When Names Collide
- **Shadowing** happens when a variable in a local scope has the same name as a variable in an outer scope.
- The local variable **hides** the outer variable.

```python
x = 5  # Global variable

def example():
    x = 10  # Local variable (shadows global x)
    print(x)  # Prints 10

example()
print(x)  # Prints 5 (original x remains unchanged)
```

# Common Mistakes with Scope
## 1. Using a local variable before assignment
```python
def error_example():
    print(x)  # Raises UnboundLocalError
    x = 5
```

# 2. Unexpected shadowing
```python
name = "Alice"

def show_name():
    name = "Bob"  # Shadows global 'name'
    print(name)

show_name()  # Prints "Bob"
print(name)  # Prints "Alice"
```

# 3. Modifying global variables without `global`
```python
count = 0

def increase():
    global count  # Without this, it would raise an error
    count += 1
```

# Best Practices
- Avoid modifying global variables inside functions.
- Use function arguments to pass data instead of relying on global variables.
- Be careful with variable names to prevent shadowing.
- Use `nonlocal` inside nested functions to modify enclosing variables.
