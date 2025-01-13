import sqlite3


class Task:
    db_name = "tasks.db"

    def __init__(self, db_name="tasks.db"):
        self.db_name = db_name
        self._create_table()

    @staticmethod
    def _create_table():
        with sqlite3.connect(Task.db_name) as conn:
            conn.execute(
                """CREATE TABLE IF NOT EXISTS tasks (
                                id INTEGER PRIMARY KEY AUTOINCREMENT,
                                description TEXT NOT NULL,
                                completed BOOLEAN NOT NULL DEFAULT 0
                            )"""
            )

    @staticmethod
    def add_task(description):
        with sqlite3.connect(Task.db_name) as conn:
            conn.execute(
                "INSERT INTO tasks (description, completed) VALUES (?, ?)",
                (description, False),
            )

    @staticmethod
    def get_tasks():
        with sqlite3.connect(Task.db_name) as conn:
            return conn.execute("SELECT * FROM tasks").fetchall()

    @staticmethod
    def complete_task(task_id):
        with sqlite3.connect(Task.db_name) as conn:
            conn.execute("UPDATE tasks SET completed = 1 WHERE id = ?", (task_id,))
