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

# Advanced Data Retrieval with SQLite in Python

## Why Improve Data Retrieval?
- Makes database queries more efficient and readable.
- Prevents common security risks.
- Enhances usability with dictionary-style access.

# Using `row_factory` for Dictionary Access
### Default Behavior
- SQLite queries return rows as tuples:
  ```python
  cursor.execute("SELECT * FROM users")
  row = cursor.fetchone()
  print(row[0], row[1])  # Access via index
  ```
- This can be hard to read and maintain.

# Enabling Dictionary-Like Access
- Use `sqlite3.Row` for named column access:
  ```python
  conn.row_factory = sqlite3.Row
  ```
  ```python
  cursor.execute("SELECT * FROM users")
  for row in cursor.fetchall():
      print(row["name"], row["age"])  # Access via column name
  ```

# Why Use `row_factory`?
✅ More readable and maintainable code
✅ No need to remember column index positions
✅ Works seamlessly with loops and dictionaries

# Using Parameterized Queries for Security
### SQL Injection Risk ⚠️
```python
name = input("Enter name: ")
cursor.execute(f"SELECT * FROM users WHERE name = '{name}'")  # 🚨 Vulnerable!
```

# Safe Query Using Parameters ✅
```python
cursor.execute("SELECT * FROM users WHERE name = ?", (name,))
```
✅ Prevents SQL injection
✅ Ensures data is properly escaped

# Using `with` for Connection Management
### Traditional Way
```python
conn = sqlite3.connect("example.db")
cursor = conn.cursor()
# Execute queries...
conn.close()
```

### Better Approach with `with` ✅
```python
with sqlite3.connect("example.db") as conn:
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM users")
    print(cursor.fetchall())
```
✅ Ensures automatic closing of connection
✅ Reduces manual error handling

# Handling Errors Gracefully
### Catching SQLite Exceptions
```python
import sqlite3
try:
    conn = sqlite3.connect("example.db")
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM non_existing_table")
except sqlite3.Error as e:
    print("Database error:", e)
finally:
    conn.close()
```
✅ Prevents crashes
✅ Helps diagnose issues

# Fetching Data Efficiently
### Three Ways to Retrieve Data
```python
cursor.execute("SELECT * FROM users")
```
1️⃣ **fetchone()** - Get a single row:
   ```python
   row = cursor.fetchone()
   print(row)
   ```

2️⃣ **fetchall()** - Get all rows:
   ```python
   rows = cursor.fetchall()
   for row in rows:
       print(row)
   ```

3️⃣ **fetchmany(n)** - Get `n` rows at a time:
   ```python
   rows = cursor.fetchmany(2)  # Get first 2 rows
   ```

# Summary
✅ Use `row_factory` for better readability.
✅ Always use parameterized queries to prevent SQL injection.
✅ Manage connections with `with` to avoid forgetting `close()`.
✅ Use `try-except` blocks for error handling.
✅ Choose `fetchone()`, `fetchall()`, or `fetchmany(n)` as needed.
