---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# **File Handling and Persistent Data Storage**

# **Warm-Up**  

  1. Write a function `sum_list(numbers)` that calculates the sum of a list.
  2. Create a list of 3 numbers entered by the user and pass it to `sum_list()`.

# **Writing to a File**

**File Modes**:
- `"w"`: Write mode (overwrites the file).
- `"a"`: Append mode (adds to the end of the file).
- `"r"`: Read mode.

**Example: Writing to a File**:
```python
# Writing data to a file
with open("expenses.txt", "w") as file:
    file.write("Groceries: $50\n")
    file.write("Rent: $1000\n")
print("Data written to file.")
```

# **Activity**:
- Write a program to ask the user for their name and age, then save it to a file called `user_info.txt`.

**Solution**:
```python
name = input("Enter your name: ")
age = input("Enter your age: ")

with open("user_info.txt", "w") as file:
    file.write(f"Name: {name}\n")
    file.write(f"Age: {age}\n")

print("User info saved to user_info.txt.")
```

# Reading from a File**

**How to Read a File**:
- Use `"r"` mode to read the contents of a file.
- Use `read()` or `readlines()` to access the data.

**Example**:
```python
# Reading data from a file
with open("expenses.txt", "r") as file:
    contents = file.read()
    print(contents)
```

# **Activity**:
- Write a program to read and display the contents of `user_info.txt` created earlier.

# **Combining Lists and File Operations**

**Scenario**: Save and load a list of expenses.  

# **Example: Saving a List to a File**:
```python
expenses = [50, 100, 25]

with open("expenses.txt", "w") as file:
    for expense in expenses:
        file.write(f"{expense}\n")

print("Expenses saved to file.")
```

# **Example: Loading a List from a File**:
```python
expenses = []
with open("expenses.txt", "r") as file:
    for line in file:
        expenses.append(float(line.strip()))

print("Expenses loaded:", expenses)
```

# **Activity**:

1. Write a program that:
   - Asks the user to input 3 expenses.
   - Saves the expenses to a file.
   - Reads the expenses back from the file and calculates their total.

# 5. **Assignment**

**Objective**:
Create a program that manages a list of expenses, saves them to a file, and reloads them when the program restarts.

**Instructions**:
1. When the program starts, check if `expenses.txt` exists:
   - If it does, load existing expenses into a list.
   - If not, start with an empty list.
2. Allow the user to:
   - Add new expenses.
   - View all expenses.
   - Calculate and display the total and average of all expenses.
   - Save expenses to `expenses.txt` before exiting.

**Bonus**:
- Allow the user to delete an expense by its position in the list.

# **Reflection Questions**

1. Why is `os.path.exists()` used when loading expenses?
2. How does the program ensure expenses are saved persistently?
3. How could you extend this program to handle categories for expenses?
