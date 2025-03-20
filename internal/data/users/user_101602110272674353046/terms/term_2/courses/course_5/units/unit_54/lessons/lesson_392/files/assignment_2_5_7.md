# Unit 5 Project

## **"Movie Collection Manager"**  
**Type:** CLI Application  
**Topics:** OOP, SQLite, CRUD Operations  

---

### **Project Overview**
In this project, students will build a command-line interface (CLI) to manage a collection of movies. They'll be provided with starter code that sets up the database schema and the basic structure of the application. Students will implement methods to add, update, delete, and display movies, using object-oriented programming and SQLite.

---

### **Starter Code Outline**
**Files Provided:**
1. `movie.py` – Defines the `Movie` class  
2. `database.py` – Handles SQLite connection and setup  
3. `main.py` – CLI interface  

---

### **Starter Code Example**
#### `movie.py`
```python
class Movie:
    def __init__(self, title, director, year):
        self.title = title
        self.director = director
        self.year = year

    def __str__(self):
        return f"{self.title} ({self.year}) directed by {self.director}"
```

#### `database.py`
```python
import sqlite3

def connect():
    conn = sqlite3.connect("movies.db")
    cursor = conn.cursor()
    cursor.execute(
        """
        CREATE TABLE IF NOT EXISTS movies (
            id INTEGER PRIMARY KEY,
            title TEXT,
            director TEXT,
            year INTEGER
        )
        """
    )
    conn.commit()
    conn.close()

def add_movie(title, director, year):
    # TODO: Implement this
    pass

def get_movies():
    # TODO: Implement this
    pass

def delete_movie(movie_id):
    # TODO: Implement this
    pass
```

#### `main.py`
```python
from database import connect, add_movie, get_movies, delete_movie
from movie import Movie

def menu():
    print("\nMovie Collection Manager")
    print("1. Add Movie")
    print("2. List Movies")
    print("3. Delete Movie")
    print("4. Quit")

def add_movie_prompt():
    title = input("Enter title: ")
    director = input("Enter director: ")
    year = int(input("Enter year: "))
    # TODO: Call add_movie here

def list_movies():
    # TODO: Call get_movies and display them

def delete_movie_prompt():
    movie_id = int(input("Enter movie ID to delete: "))
    # TODO: Call delete_movie here

def main():
    connect()
    while True:
        menu()
        choice = input("Choose an option: ")
        if choice == "1":
            add_movie_prompt()
        elif choice == "2":
            list_movies()
        elif choice == "3":
            delete_movie_prompt()
        elif choice == "4":
            break
        else:
            print("Invalid option")

if __name__ == "__main__":
    main()
```

---

### **What Students Must Implement**
✅ In `database.py`:  
- Complete `add_movie()` to insert a movie into the database  
- Complete `get_movies()` to list all movies  
- Complete `delete_movie()` to delete a movie by ID  

✅ In `main.py`:  
- Complete `add_movie_prompt()` to call `add_movie()`  
- Complete `list_movies()` to call `get_movies()`  
- Complete `delete_movie_prompt()` to call `delete_movie()`  

---

### **Extra Credit Ideas**
- Add a search feature (search by title or director)  
- Implement data validation (e.g., prevent adding duplicates)  
- Allow sorting by title, director, or year  

---

### **Grading Criteria**

| Task | Points |
|-------|--------|
| Implement `add_movie()` | 20 |
| Implement `get_movies()` | 20 |
| Implement `delete_movie()` | 20 |
| CLI works without crashing | 10 |
| Code is well-organized and uses OOP principles | 10 |
| Extra credit | +5–10 |
