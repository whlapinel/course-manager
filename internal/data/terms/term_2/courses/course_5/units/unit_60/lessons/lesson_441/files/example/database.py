import sqlite3


def setup_database():
    conn = sqlite3.connect("library.db")
    conn.execute(
        """
        CREATE TABLE IF NOT EXISTS books (
            id INTEGER PRIMARY KEY,
            title TEXT NOT NULL,
            author TEXT NOT NULL,
            year INTEGER NOT NULL
        )
    """
    )
    conn.commit()
    conn.close()


setup_database()
print("Database and books table created.")
