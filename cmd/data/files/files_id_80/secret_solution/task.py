import sqlite3


class Task:
    def __init__(self, db_name="tasks.db"):
        self.db_name = db_name
        self._create_table()

    def _create_table(self):
        with sqlite3.connect(self.db_name) as conn:
            conn.execute(
                """CREATE TABLE IF NOT EXISTS tasks (
                                id INTEGER PRIMARY KEY AUTOINCREMENT,
                                description TEXT NOT NULL,
                                completed BOOLEAN NOT NULL DEFAULT 0
                            )"""
            )

    def add_task(self, description):
        with sqlite3.connect(self.db_name) as conn:
            conn.execute(
                "INSERT INTO tasks (description, completed) VALUES (?, ?)",
                (description, False),
            )

    def get_tasks(self):
        with sqlite3.connect(self.db_name) as conn:
            return conn.execute("SELECT * FROM tasks").fetchall()

    def complete_task(self, task_id):
        with sqlite3.connect(self.db_name) as conn:
            conn.execute("UPDATE tasks SET completed = 1 WHERE id = ?", (task_id,))
