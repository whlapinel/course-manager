### **Assignment: Build a Simple Task Manager**

### **Overview**
In this project, you will create a small Flask-based **Task Manager** application using SQLite for data storage and Object-Oriented Programming (OOP) principles. The application will allow users to:
1. View a list of tasks.
2. Add a new task.
3. Mark a task as completed.

### **Objectives**
1. Use Flask to create a basic web application.
2. Interact with an SQLite database for storing and retrieving tasks.
3. Apply OOP principles by creating a `Task` class for managing task-related operations.

### **Requirements**
1. **Task Class**:
   - Create a `Task` class to represent individual tasks.
   - Include methods for adding, retrieving, and marking tasks as completed.

2. **Flask Routes**:
   - A homepage (`/`) that displays all tasks.
   - A form for adding new tasks.
   - A link or button to mark a task as completed.

3. **SQLite Database**:
   - Store task data in an SQLite database.
   - Use SQLAlchemy or `sqlite3` for database interaction.

4. **Templates**:
   - Create basic HTML templates for rendering the task list and form.

### **Step-by-Step Instructions**

#### 1. **Set Up the Project**
- Create the following folder structure:

   ```text
   project/
       app.py
       templates/
           base.html
           index.html
       static/
           styles.css
       tasks.db
   ```

#### 2. **Task Class**
Define a `Task` class in `app.py`:

```python
import sqlite3

class Task:
    def __init__(self, db_name="tasks.db"):
        self.db_name = db_name
        self._create_table()

    def _create_table(self):
        with sqlite3.connect(self.db_name) as conn:
            conn.execute('''CREATE TABLE IF NOT EXISTS tasks (
                                id INTEGER PRIMARY KEY AUTOINCREMENT,
                                description TEXT NOT NULL,
                                completed BOOLEAN NOT NULL DEFAULT 0
                            )''')

    def add_task(self, description):
        with sqlite3.connect(self.db_name) as conn:
            conn.execute("INSERT INTO tasks (description, completed) VALUES (?, ?)", (description, False))

    def get_tasks(self):
        with sqlite3.connect(self.db_name) as conn:
            return conn.execute("SELECT * FROM tasks").fetchall()

    def complete_task(self, task_id):
        with sqlite3.connect(self.db_name) as conn:
            conn.execute("UPDATE tasks SET completed = 1 WHERE id = ?", (task_id,))
```

#### 3. **Flask App**

Set up Flask routes in `app.py`:

```python
from flask import Flask, render_template, request, redirect, url_for
from task import Task

app = Flask(__name__)
task_manager = Task()

@app.route('/')
def index():
    tasks = task_manager.get_tasks()
    return render_template('index.html', tasks=tasks)

@app.route('/add', methods=['POST'])
def add_task():
    description = request.form.get('description')
    if description:
        task_manager.add_task(description)
    return redirect(url_for('index'))

@app.route('/complete/<int:task_id>')
def complete_task(task_id):
    task_manager.complete_task(task_id)
    return redirect(url_for('index'))

if __name__ == '__main__':
    app.run(debug=True)
```

#### 4. **Templates**

**`base.html`**:

```html
<!DOCTYPE html>
<html>
<head>
    <title>Task Manager</title>
    <link rel="stylesheet" href="{{ url_for('static', filename='styles.css') }}">
</head>
<body>
    <header>
        <h1>Task Manager</h1>
    </header>
    <main>
        {% block content %}{% endblock %}
    </main>
</body>
</html>
```

**`index.html`**:

```html
{% extends 'base.html' %}
{% block content %}
<h2>Tasks</h2>
<ul>
    {% for task in tasks %}
    <li>
        <span style="text-decoration: {{ 'line-through' if task[2] else 'none' }}">{{ task[1] }}</span>
        {% if not task[2] %}
        <a href="{{ url_for('complete_task', task_id=task[0]) }}">Mark Complete</a>
        {% endif %}
    </li>
    {% endfor %}
</ul>
<h3>Add Task</h3>
<form action="{{ url_for('add_task') }}" method="POST">
    <input type="text" name="description" placeholder="Task description" required>
    <button type="submit">Add</button>
</form>
{% endblock %}
```

### **Submission Instructions**

1. Verify your app works:
   - Add tasks.
   - Mark tasks as completed.
   - See changes reflected dynamically.
2. Zip your `project/` folder and submit it.

### **Grading Rubric**

| **Criteria**                  | **Points** |
| ----------------------------- | ---------- |
| Flask app structure           | 10         |
| Working `Task` class          | 20         |
| Dynamic task list display     | 30         |
| Task addition functionality   | 20         |
| Task completion functionality | 20         |
