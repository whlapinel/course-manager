import sqlite3


class Task:
    db_name = "tasks.db"

    @staticmethod
    def _create_table():
        with sqlite3.connect(Task.db_name) as conn:
            pass
            # Bug(1): Implement this method using the query
            # CREATE TABLE IF NOT EXISTS tasks (
            # id INTEGER PRIMARY KEY,
            # description TEXT NOT NULL,
            # completed BOOLEAN NOT NULL
            # )

    @staticmethod
    def add_task(description):
        with sqlite3.connect(Task.db_name) as conn:
            pass
            # Bug(2): Implement this method using the query 'INSERT INTO tasks (description, completed) VALUES (?, ?)'

    @staticmethod
    def get_tasks():
        with sqlite3.connect(Task.db_name) as conn:
            pass
            # Bug(3): Implement this method using the query 'SELECT * FROM tasks'

    @staticmethod
    def complete_task(task_id):
        with sqlite3.connect(Task.db_name) as conn:
            pass
            # Bug(4): Implement this method using the query 'UPDATE tasks SET completed = 1 WHERE id = ?'
