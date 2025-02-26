---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Introduction to SQL

## What is SQL?
- SQL stands for **Structured Query Language**
- It is used to interact with databases
- Allows storing, retrieving, updating, and deleting data
- Common databases using SQL: MySQL, PostgreSQL, SQLite, SQL Server

# Why Learn SQL?
- Used in **data analysis, web development, machine learning, and more**
- Essential for managing structured data efficiently
- Enables powerful data querying and manipulation

# Getting Started with SQL
**Practice SQL Online:** [sqliteonline.com](https://sqliteonline.com/)
- Open the website
- Choose SQLite as your database
- Try running your first query!

# First Query: Creating a Table
```sql
CREATE TABLE students (
    id INTEGER PRIMARY KEY,
    name TEXT,
    age INTEGER
);
```

# Challenge #1:
Create a table called `books` with the following columns:
- `id` (INTEGER, Primary Key)
- `title` (TEXT)
- `author` (TEXT)
- `year_published` (INTEGER)

# Inserting Data
We add data using the `INSERT` statement:
```sql
INSERT INTO students (name, age) VALUES ('Alice', 21);
INSERT INTO students (name, age) VALUES ('Bob', 22);
```

# Challenge #2:
Insert three books into your `books` table with different titles, authors, and years.

# Retrieving Data
Use `SELECT` to view data:
```sql
SELECT * FROM students;
```

To view only specific columns:
```sql
SELECT name FROM students;
```

# Challenge #3:
Retrieve only the `title` and `author` columns from your `books` table.

# Filtering Data with WHERE
Use `WHERE` to filter results:
```sql
SELECT * FROM students WHERE age > 21;
```

# Challenge #4:
Retrieve all books published after 2000.

# Updating Data
Modify existing data using `UPDATE`:
```sql
UPDATE students SET age = 23 WHERE name = 'Alice';
```

# Challenge #5:
Change the author of one book in your `books` table.

# Deleting Data
Remove records using `DELETE`:
```sql
DELETE FROM students WHERE name = 'Bob';
```

# Challenge #6:
Delete a book from your `books` table where the year published is before 1990.

# Summary
- SQL helps manage databases
- `CREATE TABLE` to define structure
- `INSERT INTO` to add data
- `SELECT` to retrieve data
- `WHERE` for filtering
- `UPDATE` to modify data
- `DELETE` to remove data
