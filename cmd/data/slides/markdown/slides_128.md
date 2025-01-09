---
layout: none
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Lesson 10.1 Final Exam Review

# Exam dates

All blocks will take the exam on Tuesday, 1/9/25.

# Exam format

I will provide you a project that has a few things broken and you'll have to fix the broken parts.

In most cases I will tell you which parts are broken, but I may have some where you have to find which line or maybe even which file has incorrect or missing code.

It will involve OOP, Sqlite, and Web, as well as all of the Python-specific things you've learned.

# Review

We'll be doing some review over the next 3 days to prepare for the exam.

# Activity 1

Here’s the first practice activity designed to build up to the `complete_task` functionality. It includes a **refresher**, **mini-lessons**, and an **assignment/challenge**. This activity assumes students already have basic knowledge of Flask, routing, and templating.

# **Practice Activity 1: Introduction to Flask and SQLite Integration**

# Step 1: Set Up SQLite Database

**Objective**: Create a simple SQLite database and connect it to Flask.

**Code Example**:
```python
import sqlite3

# Create a database and table
conn = sqlite3.connect('example.db')
conn.execute('''
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    age INTEGER NOT NULL
)
''')
conn.close()
```

# Step 2: Fetch Data from SQLite

**Objective**: Fetch rows from SQLite and display them in the terminal.

**Code Example**:
```python
conn = sqlite3.connect('example.db')
cursor = conn.execute('SELECT * FROM users')
for row in cursor:
    print(row)
conn.close()
```

# Step 3: Connect SQLite to Flask

**Objective**: Dynamically fetch data in a Flask route and pass it to a template.

**Code Example**:
```python
from flask import Flask, render_template
import sqlite3

app = Flask(__name__)

@app.route('/')
def home():
    conn = sqlite3.connect('example.db')
    users = conn.execute('SELECT * FROM users').fetchall()
    conn.close()
    return render_template('index.html', users=users)

if __name__ == '__main__':
    app.run(debug=True)
```

# **Create `index.html`**

**Template Example** (`index.html`):

(You will need to modify this to suit our purposes)

```html
<!DOCTYPE html>
<html>
<head>
    <title>Users</title>
</head>
<body>
    <h1>Users List</h1>
    <ul>
        {% for user in users %}
        <li>{{ user[1] }} (Age: {{ user[2] }})</li>
        {% endfor %}
    </ul>
</body>
</html>
```

# **Assignment**

## **Objective**:

Create a small Flask app that:
1. Sets up a SQLite database with a `books` table (containing `id`, `title`, and `author`).
2. Displays all books on the homepage.
3. Dynamically passes data to an HTML template.

# **Step-by-Step Instructions**:

1. **Database Setup**:
   - Create a database `library.db` and add a `books` table with the following schema:
     ```sql
     CREATE TABLE books (
         id INTEGER PRIMARY KEY AUTOINCREMENT,
         title TEXT NOT NULL,
         author TEXT NOT NULL
     );
     ```
   - Insert at least 3 rows into the `books` table.

# **Instructions (Continued)**

2. **Flask App**:
   - Create a Flask route `/` to fetch all books and display them in a template.

3. **Template**:
   - Use a simple HTML template to display the list of books with their titles and authors.

# **Bonus**:

- Add another route `/book/<id>` that shows detailed information about a specific book based on its ID.

# **Submission Requirements**

Submit your project as a .zip file

# **Grading Rubric**

| **Criteria**                      | **Points** |
| --------------------------------- | ---------- |
| Flask app runs without errors     | 10         |
| SQLite database correctly created | 20         |
| Homepage displays all books       | 30         |
| Template dynamically renders data | 20         |
| Bonus: Detailed book route        | 20         |

# Next Class Preview
In the next class, students will:
- Learn how to handle user input via forms.
- Create a form to add new rows to the SQLite database.

