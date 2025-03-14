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
- Create the following and open it in VS Code

```text
/python
  /unit_5
    /lesson_4
```

# Closures in Python

## Understanding Closures and How They Work

# What is a Closure?
- A **closure** is a function object that remembers values in the enclosing lexical scope even when the scope is no longer available.
- Closures allow functions to:
  - Retain state between calls
  - Encapsulate logic and data

# **Example:**  
```python
def outer(x):
    def inner(y):
        return x + y
    return inner

closure = outer(10)
print(closure(5))  # Output: 15
```
- `inner` remembers the value of `x` even after `outer` has returned.

# How Closures Work
1. A nested function is defined within another function.
2. The inner function references variables from the outer function.
3. The outer function returns the inner function.
4. The inner function **retains access** to the outer function's scope.

# **Example:**  
```python
def multiplier(factor):
    def multiply_by(x):
        return x * factor
    return multiply_by

double = multiplier(2)
print(double(10))  # Output: 20
```

# Why Use Closures?
✅ Encapsulation  
✅ State Retention  
✅ Avoids Global Variables  
✅ Useful for Factory Functions

# **Example: Counter**  
```python
def make_counter():
    count = 0
    def counter():
        nonlocal count
        count += 1
        return count
    return counter

counter = make_counter()
print(counter())  # Output: 1
print(counter())  # Output: 2
```

# Comparison with Other Approaches
## 1. Global Variables (Less Encapsulation)
```python
count = 0

def counter():
    global count
    count += 1
    return count

print(counter())  # Output: 1
print(counter())  # Output: 2
```
- **Downside:** Pollutes global namespace  
- **Closures:** Keep state private and encapsulated

# 2. Using a Class (More Boilerplate)
```python
class Counter:
    def __init__(self):
        self.count = 0

    def __call__(self):
        self.count += 1
        return self.count

counter = Counter()
print(counter())  # Output: 1
print(counter())  # Output: 2
```
- **Downside:** More code to maintain state  
- **Closures:** Less code, easier to use

# When to Use Closures
✅ When you want to retain state between function calls  
✅ When you want to avoid global variables  
✅ When you need a factory function  
✅ When you need lightweight objects without a class  

# When NOT to Use Closures
❌ When state becomes too complex (use a class instead)  
❌ When multiple levels of nesting make the code hard to follow  
❌ When the function needs to be serialized or copied  

# Summary
- Closures capture and retain state from their enclosing scope.  
- They are simpler and more lightweight than classes for many tasks.  
- Great for small, stateful, reusable functions.  

**Example:**  
```python
def adder(x):
    def add(y):
        return x + y
    return add

add_five = adder(5)
print(add_five(3))  # Output: 8
```
