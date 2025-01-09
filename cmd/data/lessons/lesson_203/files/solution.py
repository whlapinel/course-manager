import sqlite3


# Create a connection to the SQLite database
def create_connection(db_name="movies.db"):
    """Create a database connection to the SQLite database"""
    try:
        conn = sqlite3.connect(db_name)
        print("Connected to database successfully.")
        return conn
    except sqlite3.Error as e:
        print(f"Error connecting to database: {e}")
        return None


# Create the movies table
def create_table(conn):
    """Create the movies table in the database"""
    try:
        cursor = conn.cursor()
        cursor.execute(
            """
            CREATE TABLE IF NOT EXISTS movies (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                title TEXT NOT NULL,
                genre TEXT,
                release_year INTEGER,
                director TEXT,
                rating REAL
            )
        """
        )
        conn.commit()
        print("Table created successfully.")
    except sqlite3.Error as e:
        print(f"Error creating table: {e}")


# Function to add a new movie
def add_movie(conn):
    """Add a new movie to the collection"""
    title = input("Enter the movie title: ")
    genre = input("Enter the genre: ")
    release_year = int(input("Enter the release year: "))
    director = input("Enter the director's name: ")
    rating = float(input("Enter the movie rating (0.0 - 10.0): "))

    try:
        cursor = conn.cursor()
        cursor.execute(
            """
            INSERT INTO movies (title, genre, release_year, director, rating)
            VALUES (?, ?, ?, ?, ?)
        """,
            (title, genre, release_year, director, rating),
        )
        conn.commit()
        print(f"Movie '{title}' added successfully.")
    except sqlite3.Error as e:
        print(f"Error adding movie: {e}")


# Function to update an existing movie
def update_movie(conn):
    """Update the details of an existing movie"""
    try:
        movie_id = int(input("Enter the ID of the movie to update: "))
        cursor = conn.cursor()
        cursor.execute("SELECT * FROM movies WHERE id = ?", (movie_id,))
        movie = cursor.fetchone()

        if not movie:
            print("No movie found with that ID.")
            return

        new_title = input(f"Enter new title ({movie[1]}): ") or movie[1]
        new_genre = input(f"Enter new genre ({movie[2]}): ") or movie[2]
        new_release_year = input(f"Enter new release year ({movie[3]}): ") or movie[3]
        new_director = input(f"Enter new director ({movie[4]}): ") or movie[4]
        new_rating = input(f"Enter new rating ({movie[5]}): ") or movie[5]

        cursor.execute(
            """
            UPDATE movies
            SET title = ?, genre = ?, release_year = ?, director = ?, rating = ?
            WHERE id = ?
        """,
            (
                new_title,
                new_genre,
                int(new_release_year),
                new_director,
                float(new_rating),
                movie_id,
            ),
        )
        conn.commit()
        print("Movie updated successfully.")
    except sqlite3.Error as e:
        print(f"Error updating movie: {e}")


# Function to delete a movie
def delete_movie(conn):
    """Delete a movie from the collection"""
    try:
        movie_id = int(input("Enter the ID of the movie to delete: "))
        cursor = conn.cursor()
        cursor.execute("DELETE FROM movies WHERE id = ?", (movie_id,))
        conn.commit()
        print("Movie deleted successfully.")
    except sqlite3.Error as e:
        print(f"Error deleting movie: {e}")


# Function to search for movies
def search_movies(conn):
    """Search for movies based on title, director, or genre"""
    criterion = input("Search by title (t), director (d), or genre (g): ").lower()
    search_term = input("Enter the search term: ")

    try:
        cursor = conn.cursor()
        if criterion == "t":
            cursor.execute(
                "SELECT * FROM movies WHERE title LIKE ?", ("%" + search_term + "%",)
            )
        elif criterion == "d":
            cursor.execute(
                "SELECT * FROM movies WHERE director LIKE ?", ("%" + search_term + "%",)
            )
        elif criterion == "g":
            cursor.execute(
                "SELECT * FROM movies WHERE genre LIKE ?", ("%" + search_term + "%",)
            )
        else:
            print("Invalid search criterion.")
            return

        results = cursor.fetchall()
        if results:
            for row in results:
                print(row)
        else:
            print("No movies found.")
    except sqlite3.Error as e:
        print(f"Error searching movies: {e}")


# Function to display all movies sorted by rating
def display_movies_by_rating(conn):
    """Display all movies sorted by rating"""
    try:
        cursor = conn.cursor()
        cursor.execute("SELECT * FROM movies ORDER BY rating DESC")
        movies = cursor.fetchall()

        for movie in movies:
            print(movie)
    except sqlite3.Error as e:
        print(f"Error displaying movies by rating: {e}")


# Function to display movies from a specific year
def display_movies_by_year(conn):
    """Display all movies from a specific year"""
    try:
        year = int(input("Enter the release year: "))
        cursor = conn.cursor()
        cursor.execute("SELECT * FROM movies WHERE release_year = ?", (year,))
        movies = cursor.fetchall()

        for movie in movies:
            print(movie)
    except sqlite3.Error as e:
        print(f"Error displaying movies by year: {e}")


# Extra Credit Task: Get the highest-rated movie
def get_highest_rated_movie(conn):
    """Retrieve and display the highest-rated movie(s)"""
    try:
        cursor = conn.cursor()
        cursor.execute(
            "SELECT * FROM movies WHERE rating = (SELECT MAX(rating) FROM movies)"
        )
        highest_rated_movies = cursor.fetchall()

        if highest_rated_movies:
            print("Highest Rated Movie(s):")
            for movie in highest_rated_movies:
                print(movie)
        else:
            print("No movies in the database.")
    except sqlite3.Error as e:
        print(f"Error retrieving highest-rated movie: {e}")


def main():
    conn = create_connection()
    if conn:
        create_table(conn)

        while True:
            print("\nMovie Collection Manager")
            print("1. Add a new movie")
            print("2. Update a movie")
            print("3. Delete a movie")
            print("4. Search for a movie")
            print("5. Display all movies sorted by rating")
            print("6. Display all movies from a specific year")
            print("7. Display highest-rated movie (Extra Credit)")
            print("8. Exit")

            choice = input("Enter your choice: ")

            if choice == "1":
                add_movie(conn)
            elif choice == "2":
                update_movie(conn)
            elif choice == "3":
                delete_movie(conn)
            elif choice == "4":
                search_movies(conn)
            elif choice == "5":
                display_movies_by_rating(conn)
            elif choice == "6":
                display_movies_by_year(conn)
            elif choice == "7":
                get_highest_rated_movie(conn)
            elif choice == "8":
                print("Exiting program.")
                break
            else:
                print("Invalid choice. Please try again.")

        conn.close()
        print("Database connection closed.")


if __name__ == "__main__":
    main()
