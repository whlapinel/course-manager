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

# **Lesson 10.2: Combining SQLite and Flask**

# **Warm-Up**

Ask students:
1. What is the difference between Flask routes for `GET` and `POST` requests?
2. What is the purpose of OOP in Python?
3. How do we use SQLite in a Python project?

# **Creating the `books` Table in SQLite**

**Goal**: Set up the database and define a `books` table.

**Schema**:
- `id`: Primary key, auto-incremented.
- `title`: The title of the book.
- `author`: The author of the book.
- `year`: The year the book was published.

# **Code**:
```python
import sqlite3

def setup_database():
    conn = sqlite3.connect('library.db')
    conn.execute('''
        CREATE TABLE IF NOT EXISTS books (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT NOT NULL,
            author TEXT NOT NULL,
            year INTEGER NOT NULL
        )
    ''')
    conn.close()

setup_database()
print("Database and books table created.")
```

# **The `Book` Class**

**Code**:
```python
class Book:
    def __init__(self, title, author, year):
        self.title = title
        self.author = author
        self.year = year

    @staticmethod
    def add_book(title, author, year):
        conn = sqlite3.connect('library.db')
        conn.execute('INSERT INTO books (title, author, year) VALUES (?, ?, ?)', 
                     (title, author, year))
        conn.commit()
        conn.close()
        print(f"Book '{title}' by {author} added.")

    @staticmethod
    def get_all_books():
        conn = sqlite3.connect('library.db')
        cursor = conn.execute('SELECT * FROM books')
        books = cursor.fetchall()
        conn.close()
        return books
```

# **Explanation**:
- `add_book`: Adds a new book to the database.
- `get_all_books`: Retrieves all books from the database.

**Activity**:
1. Add 2–3 books using the `Book.add_book()` method.
2. Retrieve and print all books using `Book.get_all_books()`.

**Example**:
```python
Book.add_book("1984", "George Orwell", 1949)
Book.add_book("To Kill a Mockingbird", "Harper Lee", 1960)
print(Book.get_all_books())
```

# **Display Books in Flask**

**Goal**: Create a Flask route to display all books dynamically.

**Code**:
```python
from flask import Flask, render_template

app = Flask(__name__)

@app.route('/')
def home():
    books = Book.get_all_books()
    return render_template('index.html', books=books)

if __name__ == '__main__':
    app.run(debug=True)
```

# **Template (`index.html`)**:
```html
<!DOCTYPE html>
<html>
<head>
    <title>Library</title>
</head>
<body>
    <h1>Library Books</h1>
    <ul>
        {% for book in books %}
        <li>{{ book[1] }} by {{ book[2] }} ({{ book[3] }})</li>
        {% endfor %}
    </ul>
</body>
</html>
```

# **Activity**:
- Run the app and verify that all books are displayed on the homepage.

# **Adding a Book via a Form**

**Goal**: Add functionality to let users add new books using a form.

**Route and Form**:
```python
from flask import request, redirect, url_for

@app.route('/add', methods=['GET', 'POST'])
def add_book():
    if request.method == 'POST':
        title = request.form['title']
        author = request.form['author']
        year = request.form['year']

        Book.add_book(title, author, int(year))
        return redirect(url_for('home'))

    return render_template('add.html')
```

# **Template (`add.html`)**:
```html
<!DOCTYPE html>
<html>
<head>
    <title>Add a Book</title>
</head>
<body>
    <h1>Add a New Book</h1>
    <form method="POST" action="{{ url_for('add_book') }}">
        <label for="title">Title:</label>
        <input type="text" id="title" name="title" required><br>

        <label for="author">Author:</label>
        <input type="text" id="author" name="author" required><br>

        <label for="year">Year:</label>
        <input type="number" id="year" name="year" required><br>

        <button type="submit">Add Book</button>
    </form>
</body>
</html>
```

# **Activity**:
- Add 2 new books through the form and verify that they appear on the homepage.

# **Challenge Assignment**

**Objective**:
Add functionality to delete a book by its ID.

**Instructions**:
# 1. Add a method `delete_book(book_id)` in the `Book` class*:
   ```python
   @staticmethod
   def delete_book(book_id):
       conn = sqlite3.connect('library.db')
       conn.execute('DELETE FROM books WHERE id = ?', (book_id,))
       conn.commit()
       conn.close()
   ```

# 2. Add a delete button next to each book in `index.html`:
   ```html
   <li>
       {{ book[1] }} by {{ book[2] }} ({{ book[3] }})
       <form method="POST" action="{{ url_for('delete_book', book_id=book[0]) }}" style="display:inline;">
           <button type="submit">Delete</button>
       </form>
   </li>
   ```

# 3. Add a `/delete/<book_id>` route to handle the delete request:
   ```python
   @app.route('/delete/<int:book_id>', methods=['POST'])
   def delete_book(book_id):
       Book.delete_book(book_id)
       return redirect(url_for('home'))
   ```

# **Reflection Questions**

1. Why do we use a `Book` class to manage books instead of writing all database logic directly in the Flask routes?
2. What would happen if you tried to delete a book that doesn’t exist in the database?
3. How could you improve the `add_book` form to prevent duplicate book entries?
