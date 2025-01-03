---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Warmup 10/15/24

- Create the following:
  - Create 1 class
  - Define the constructor (`__init__`)
  - 1 other method
  - create an object of your class
  - call your method

# Agenda 10/15/24

- Closures and assignment 5.4

# Announcements 10/15/24

- I will be absent on Thursday and Friday 10/17 and 10/18.

# Looking ahead

- Unit 5 Perform: Monday 10/21 (Alien attack project due)
- Midterm: Monday 10/28 (Multiple choice test and project)

# Closures in Python

# What is a Closure?

- A closure is a function object that remembers values in enclosing scopes even if they are not present in memory.
- Closures provide a way to retain access to variables after the outer function has finished execution.

# Components of a Closure

1. **Nested Function**: A function defined inside another function.
2. **Enclosing Variables**: Variables that are captured from the outer function.
3. **Return of Nested Function**: The inner function is returned from the outer function.

# Example of a Closure

```python
def outer_func(msg):
    def inner_func():
        print(msg)
    return inner_func

closure = outer_func('Hello, World!')
closure()  # Output: Hello, World!
```

# How Closures Work

- In the example, `msg` is an enclosed variable.
- The `inner_func` retains the value of `msg` after `outer_func` finishes.
- When `closure()` is called, it prints the value of `msg`.

# Changing closure state

**New:** to modify an outer variable you must use the `nonlocal` keyword like this:

```python
def divider(start: int):
    def divide():
        nonlocal start # need this to be able to modify the variable 'start' which is declared in the outer function
        start //= 2 # assign to start half of itself ('//' means integer division)
        return start

    return divide # return the inner function

divide = divider(1000)
print(divide())
print(divide())

divide2 = divider(200)
print(divide2())
print(divide2())
```

# Use Cases of Closures

- **Data hiding**: Closures can encapsulate data.
- **Factory functions**: Closures can generate functions with different behaviors based on arguments.
- **Callback functions**: Closures are useful in event-driven programming.

# Important Points

- Closures are created by nested functions.
- Enclosed variables are captured by the inner function.
- Functions returned by closures can access those captured variables even after the outer function has finished.

# Limitations

- Closures can sometimes lead to unintuitive behavior, especially with mutable variables.
- They can make code harder to debug due to the hidden state.

# Conclusion

- Closures are a powerful tool in Python for managing state in a clean, functional style.
- They provide a way to encapsulate data and logic in a single object, making the code modular and reusable.
