# Task 0: import the library


# Create a connection to the SQLite database
def create_connection(db_name="movies.db"):
    """Create a database connection to the SQLite database"""
    # Task 1: Use sqlite3 to connect to the database.
    # Hint: Use sqlite3.connect(db_name) and handle potential connection errors using try-except.
    pass


# Create the movies table
def create_table(conn):
    """Create the movies table in the database"""
    # Task 2: Create a cursor from the connection and execute an SQL command to create the 'movies' table.
    # Hint: Use conn.cursor() to create a cursor object, and execute the CREATE TABLE statement with cursor.execute().
    # Remember to call conn.commit() to save changes.
    pass


# Function to add a new movie
def add_movie(conn):
    """Add a new movie to the collection"""
    # Task 3: Prompt the user to input movie details (title, genre, release year, director, rating).
    # Hint: Use the input() function to collect data from the user.

    # Task 4: Insert the collected data into the 'movies' table.
    # Hint: Use cursor.execute() with an INSERT INTO statement, and pass the values as a tuple.
    pass


# Function to update an existing movie
def update_movie(conn):
    """Update the details of an existing movie"""
    # Task 5: Ask the user for the ID of the movie they want to update.
    # Hint: Use input() to get the movie ID and check if it's a valid integer.

    # Task 6: Prompt the user to enter the new details (title, genre, release year, director, rating).
    # Hint: Allow the user to press Enter to skip fields they don't want to change.

    # Task 7: Execute an UPDATE SQL statement to modify the movie record.
    # Hint: Use cursor.execute() with an UPDATE statement and ensure you only update fields that were provided.
    pass


# Function to delete a movie
def delete_movie(conn):
    """Delete a movie from the collection"""
    # Task 8: Ask the user for the ID of the movie to be deleted.
    # Hint: Validate that the input is a valid integer.

    # Task 9: Execute a DELETE SQL statement to remove the movie from the table.
    # Hint: Use cursor.execute() with a DELETE FROM statement and commit the change.
    pass


# Function to search for movies
def search_movies(conn):
    """Search for movies based on title, director, or genre"""
    # Task 10: Ask the user which criterion they want to search by (title, director, or genre).
    # Hint: Use input() to collect this choice and validate the input.

    # Task 11: Execute a SELECT SQL statement based on the chosen criterion.
    # Hint: Use the cursor.execute() function with a SELECT statement that matches the user's choice.
    pass


# Function to display all movies sorted by rating
def display_movies_by_rating(conn):
    """Display all movies sorted by rating"""
    # Task 12: Execute an SQL query to retrieve all movies sorted by the 'rating' column in descending order.
    # Hint: Use the ORDER BY clause in your SELECT statement.

    # Task 13: Display the movies to the user in a formatted manner.
    # Hint: Use a loop to iterate over the results from cursor.fetchall() and print each movie's details.
    pass


# Function to display movies from a specific year
def display_movies_by_year(conn):
    """Display all movies from a specific year"""
    # Task 14: Ask the user for the year they want to filter by.
    # Hint: Validate that the input is a valid integer.

    # Task 15: Execute a SELECT SQL statement to retrieve movies from that specific year.
    # Hint: Use a WHERE clause in your SQL statement to filter by the 'release_year' column.

    # Task 16: Display the results in a user-friendly format.
    pass


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
            print("7. Exit")

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
                print("Exiting program.")
                break
            else:
                print("Invalid choice. Please try again.")

        # Close the database connection
        conn.close()
        print("Database connection closed.")


if __name__ == "__main__":
    main()
