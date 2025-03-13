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
    /lesson_3
```

# Python Generators

## Efficient Iteration with Generators

# What are Generators?

- **Generators** are a special type of iterator in Python.
- They allow you to iterate over data without storing everything in memory.
- Unlike lists, they produce items **on the fly**.

# Why Use Generators?

- **Memory efficiency**: They don’t hold the entire data in memory.
- **Lazy evaluation**: They compute values only when needed.
- Useful for handling **large data** or **infinite sequences**.

# Example of a Generator Function

```python
def my_generator():
    yield 1
    yield 2
    yield 3
```

- **`yield`** is used instead of `return` to produce values lazily.

# Using a Generator

```python
gen = my_generator()

print(next(gen))  # Output: 1
print(next(gen))  # Output: 2
print(next(gen))  # Output: 3
```

- Generators return an iterator object and produce values one at a time.

# Iterating Through a Generator

```python
for value in my_generator():
    print(value)
```

- Output:
  - 1
  - 2
  - 3

# Generator Expression

- A **generator expression** is similar to a list comprehension but more memory efficient.
  
```python
gen_exp = (x**2 for x in range(10))
```

- This will generate squares of numbers 0 to 9 lazily.

# Infinite Generators

- You can create generators that **never end**.

```python
def infinite_counter():
    n = 0
    while True:
        yield n
        n += 1
```

# Generator Use Cases

- **Processing large files** line by line:

  ```python
  def read_large_file(file):
      for line in open(file):
          yield line
  ```

- **Generating an infinite sequence**:

  ```python
  def fibonacci():
      a, b = 0, 1
      while True:
          yield a
          a, b = b, a + b
  ```

# Key Points

- Use `yield` to define a generator.
- Generators are memory-efficient.
- You can create infinite sequences with generators.
- Generator expressions offer concise syntax for simple generators.
