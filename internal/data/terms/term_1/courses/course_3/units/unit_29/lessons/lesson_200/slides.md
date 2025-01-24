---
marp: true
---
<!-- headingDivider: 1 -->

# Warmup

- Create a dictionary `grades` where keys are are student names and values are student grades.

**Example:**

```python
grades = {
    "Kurt": 86,
    "Russell": 67,
    "Leon": 78
}
```

- Write a function `grade_avg` that accepts a dictionary where keys are student names and values are student grades.

```python
avg = grade_avg(grades)
print(avg) # 70
```

<!-- this is a test note to see if it shows up as speaker's notes -->

# Agenda

- Lesson on using SQLite
- Assignment 4.2 (using SQLite)

# Announcements

- Permission slips
- Assignment 4.1 was due Sunday night, complete and submit ASAP if you haven't already

# Introduction to SQLite

- **What is SQLite?**
  - A lightweight, embedded database engine.
  - Doesn't require a separate server to run.
  - Commonly used for local storage in applications.
  - Comes built-in with Python!

# Why Use SQLite?

- **Advantages:**
  - Easy to set up (no configuration needed).
  - Perfect for small to medium-sized applications.
  - Cross-platform and widely supported.
  
- **Disadvantages:**
  - Not suitable for very large, complex databases or high-traffic web apps.
  
# Importing SQLite in Python

- SQLite is built into Python, so no installation is required!

```python
import sqlite3
```

- The `sqlite3` module allows interaction with SQLite databases.

# Creating a Database

- **Creating a connection:**

```python
import sqlite3

conn = sqlite3.connect('students.db')  # Creates a new database file
```

- If the database file `students.db` doesn’t exist, SQLite will create it.

# Creating a Table

- **SQL Query to create a table:**

```python
conn.execute('''CREATE TABLE if not exists students
                (id INT PRIMARY KEY NOT NULL,
                 name TEXT NOT NULL,
                 age INT NOT NULL);''')
```

- **Commit the changes** to save them:

```python
conn.commit()
```

# Inserting Data

- **Insert values into the table:**

```python
conn.execute("INSERT INTO students (id, name, age) VALUES (1, 'Alice', 21)")
conn.execute("INSERT INTO students (id, name, age) VALUES (2, 'Bob', 22)")
conn.commit()
```

- Don't forget to commit after inserting data!

# Querying Data

- **Fetching data from the table:**

```python
cursor = conn.execute("SELECT id, name, age from students")
for row in cursor:
    print(row)
```

- This will print all the rows from the `students` table.

# Updating Data

- **Updating a record:**

```python
conn.execute("UPDATE students set age = 23 where id = 1")
conn.commit()
```

- This updates the age of the student with `id = 1`.

# Deleting Data

- **Deleting a record:**

```python
conn.execute("DELETE from students where id = 2")
conn.commit()
```

- This removes the student with `id = 2` from the table.

# Closing the Connection

- **Always close the connection when you're done:**

```python
conn.close()
```

- This ensures the database is saved properly and resources are freed.

# Summary

- SQLite is simple and great for beginners.
- Basic steps:
  1. Create a connection.
  2. Create tables.
  3. Insert, update, and delete data.
  4. Query data.
  5. Close the connection.
  
- Practice by building small projects with SQLite!
