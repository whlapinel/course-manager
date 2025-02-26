---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Introduction to SQLite in Python

## What is SQLite?
- A lightweight, self-contained SQL database engine.
- Does not require a separate server process.
- Great for small to medium applications.

# Why Use SQLite with Python?
- Comes built-in with Python (`sqlite3` module).
- Easy to set up and use.
- Ideal for prototyping and lightweight applications.

# Setting Up SQLite in Python
**Step 1:** Import SQLite module
```python
import sqlite3
```

**Step 2:** Connect to a database (creates one if it doesn’t exist)
```python
conn = sqlite3.connect("example.db")
```

**Step 3:** Create a cursor to execute SQL commands
```python
cursor = conn.cursor()
```

# Creating a Table
```python
cursor.execute('''CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT,
    age INTEGER
)''')
conn.commit()
```

# Inserting Data
```python
cursor.execute("INSERT INTO users (name, age) VALUES (?, ?)", ("Alice", 25))
conn.commit()
```

# Retrieving Data
```python
cursor.execute("SELECT * FROM users")
rows = cursor.fetchall()
for row in rows:
    print(row)
```

# Updating Data
```python
cursor.execute("UPDATE users SET age = ? WHERE name = ?", (26, "Alice"))
conn.commit()
```

# Deleting Data
```python
cursor.execute("DELETE FROM users WHERE name = ?", ("Alice",))
conn.commit()
```

# Closing the Connection
```python
conn.close()
```

# Summary
- Use `sqlite3.connect()` to connect to a database.
- Use `cursor.execute()` to run SQL commands.
- Always `commit()` changes and `close()` the connection when done.
- SQLite is a great starting point for working with databases in Python!
