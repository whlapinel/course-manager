import sqlite3


class Book:
    def __init__(self, title, author, year):
        self.title = title
        self.author = author
        self.year = year

    @staticmethod
    def add_book(title, author, year):
        conn = sqlite3.connect("library.db")
        conn.execute(
            "INSERT INTO books (title, author, year) VALUES (?, ?, ?)",
            (title, author, year),
        )
        conn.commit()
        conn.close()
        print(f"Book '{title}' by {author} added.")

    @staticmethod
    def get_all_books():
        conn = sqlite3.connect("library.db")
        cursor = conn.execute("SELECT * FROM books")
        books = cursor.fetchall()
        conn.commit()
        conn.close()
        return books

    @staticmethod
    def delete_book(book_id):
        conn = sqlite3.connect("library.db")
        conn.execute("DELETE FROM books WHERE id = ?", (book_id,))
        conn.commit()
        conn.close()
