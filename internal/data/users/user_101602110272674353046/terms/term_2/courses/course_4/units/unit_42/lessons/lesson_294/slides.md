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

# Introduction to Dictionaries in Python

## What is a Dictionary?
- A **dictionary** is a built-in Python data structure.
- Stores data in **key-value pairs**.
- Similar to a real-world dictionary:
  - **Key** = the word
  - **Value** = the definition

# Creating a Dictionary
```python
# Empty dictionary
my_dict = {}

# Dictionary with initial values
person = {
    "name": "Alice",
    "age": 30,
    "city": "New York"
}
```
- Keys must be **unique**.
- Keys can be **strings, numbers**, or **tuples** (immutable types).

# Accessing Values
```python
person = {"name": "Alice", "age": 30}
print(person["name"]) # Output: Alice
```
- Use **square brackets** to access values by key.
- Accessing a non-existent key raises a **KeyError**.

# Adding and Updating Values
```python
person["job"] = "Engineer" # Adding a new key-value pair
person["age"] = 31         # Updating an existing value
print(person)
```
- Dictionaries are **mutable** (can be modified).

# Removing Values
```python
person.pop("age")     # Removes a key and returns its value
del person["job"]     # Removes a key completely
print(person)
```
- `pop()` returns the value.
- `del` removes the key without returning the value.

# Dictionary Methods

| Method | Description |
|--------|-------------|
| `keys()` | Returns a view object with all keys |
| `values()` | Returns a view object with all values |
| `items()` | Returns key-value pairs as tuples |
| `get(key)` | Returns value or `None` if key not found |
| `update()` | Merges another dictionary |

# Example of Methods
```python
person = {"name": "Alice", "age": 30}

# Get keys
print(person.keys())

# Get values
print(person.values())

# Get key-value pairs
print(person.items())
```

# Looping Through a Dictionary
```python
for key, value in person.items():
    print(f"{key}: {value}")
```
- Use `.items()` to loop through key-value pairs.

# `get()` vs Direct Access
```python
print(person.get("name")) # Returns "Alice"
print(person.get("height")) # Returns None (no KeyError)
```
- `get()` is safer than direct access when the key might not exist.

# Example Program
```python
grades = {"Alice": 90, "Bob": 85, "Charlie": 92}

# Add a new entry
grades["David"] = 88

# Update an existing value
grades["Bob"] = 87

# Print all entries
for student, grade in grades.items():
    print(f"{student}: {grade}")
```
- Combines creating, adding, updating, and looping.

# Summary
✅ Stores key-value pairs  
✅ Fast access using keys  
✅ Mutable and dynamic  
✅ Provides useful methods (`get`, `items`, `keys`, `values`)  

# Practice Questions
1. Create a dictionary to store the names and ages of 3 people.
2. Add a new key-value pair to the dictionary.
3. Write a loop to print each key-value pair.
