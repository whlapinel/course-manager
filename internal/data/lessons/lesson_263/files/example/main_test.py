import database
from book import Book


def main():
    database.setup_database()
    Book.add_book("1984", "George Orwell", 1949)
    Book.add_book("To Kill a Mockingbird", "Harper Lee", 1960)
    print(Book.get_all_books())


if __name__ == "__main__":
    main()
