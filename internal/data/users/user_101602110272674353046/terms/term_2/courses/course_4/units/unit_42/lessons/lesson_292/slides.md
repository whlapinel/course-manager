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
    /lesson_1
```

# **What is a List?**
- **Definition**: A list is a collection of ordered items.
- **List Example**:  
  ```python
  my_list = [1, 2, 3, 4, 5]
  ```
- **Key Characteristics**:
  - Ordered (items have an index)
  - Mutable (can be changed after creation)
  - Can store different data types (integers, strings, etc.)

# **Creating Lists**
- **Basic List Syntax**:
  ```python
  my_list = [element1, element2, element3]
  ```
- **Examples**:
  - A list of numbers:  
    ```python
    numbers = [1, 2, 3, 4]
    ```
  - A list of strings:  
    ```python
    fruits = ["apple", "banana", "cherry"]
    ```
  - A list of mixed data types:  
    ```python
    mixed = [1, "apple", 3.14, True]
    ```

# **Accessing List Elements**
- **Access by Index**:
  ```python
  fruits = ["apple", "banana", "cherry"]
  print(fruits[0])  # Output: apple
  ```
- **Indexing**:
  - Lists are zero-indexed, so the first item is at index 0.
- **Negative Indexing**:  
  Access items from the end of the list.
  ```python
  print(fruits[-1])  # Output: cherry
  ```

# **Modifying Lists**
- **Change an item**:
  ```python
  numbers = [1, 2, 3]
  numbers[1] = 99
  print(numbers)  # Output: [1, 99, 3]
  ```
- **Add items to a list**:
  - Using `append()`:
    ```python
    fruits.append("date")
    print(fruits)  # Output: ["apple", "banana", "cherry", "date"]
    ```
  - Using `insert()`:
    ```python
    fruits.insert(1, "kiwi")
    print(fruits)  # Output: ["apple", "kiwi", "banana", "cherry"]
    ```

# **Removing Items from a List**
- **Remove by Value**:
  ```python
  fruits.remove("banana")
  print(fruits)  # Output: ["apple", "kiwi", "cherry"]
  ```
- **Remove by Index**:
  ```python
  del fruits[1]
  print(fruits)  # Output: ["apple", "cherry"]
  ```
- **Pop an Item** (removes the last item):
  ```python
  popped_item = fruits.pop()
  print(popped_item)  # Output: cherry
  ```

# **List Length and Checking Membership**
- **Getting the Length of a List**:
  ```python
  length = len(fruits)
  print(length)  # Output: 2
  ```
- **Check if an item is in a list**:
  ```python
  if "apple" in fruits:
      print("Apple is in the list!")
  ```

# **Summary**
- **Key Points**:
  - Lists are ordered, mutable collections.
  - You can access, modify, and remove elements using indices.
  - Lists support various methods like `append()`, `remove()`, and `sort()`.
  - List comprehension allows you to create new lists concisely.
  - Nested lists are lists inside lists.
