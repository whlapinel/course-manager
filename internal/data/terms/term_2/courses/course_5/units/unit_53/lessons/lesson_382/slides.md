---
marp: true
---

<!-- headingDivider: 1 -->

# Creating the connection

```Python
def create_connection(db_name="movies.db"):
    """Create a database connection to the SQLite database"""
    try:
        # insert your code here
    except sqlite3.Error as e:
        print(f"Error connecting to database: {e}")
        return None
```

# Creating the table

```python
def create_table(conn):
    """Create the movies table in the database"""
    try:
        cursor = conn.cursor()
        cursor.execute(
            """
            CREATE TABLE IF NOT EXISTS movies (
                id INTEGER PRIMARY KEY,
            )
        """
        # add the other columns needed for the application
        # note: if you specify the id as INTEGER PRIMARY KEY like above, you don't need to 
        # include it in your INSERT statements and it will automatically be filled 
        # with the next highest id value in the table
        )
        conn.commit()
        print("Table created successfully.")
    except sqlite3.Error as e:
        print(f"Error creating table: {e}")
```

# Add Movie

```python
# Function to add a new movie
def add_movie(conn):
    """Add a new movie to the collection"""

    title = input("Enter the movie title: ")
    # take user input and store into variables similarly for other fields
    # be sure to typecast to integer or float when you need a number!

    try:
        # cursor = conn.cursor()
        # cursor.execute(
            # INSERT INTO ... (put your SQL here)
        # )
        # conn.commit()
        print(f"Movie '{title}' added successfully.")
    except sqlite3.Error as e:
        print(f"Error adding movie: {e}")
```

# Updating a movie

```python
def update_movie(conn):
    """Update the details of an existing movie"""
    try:
        movie_id = int(input("Enter the ID of the movie to update: "))
        # cursor = conn.cursor()
        # cursor.execute("SELECT * FROM ... (put your SQL to get the row for that id here))
        # movie = cursor.fetchone() (this gives you the row as a list, so movie is a list 
        # where each item is the value for a column. First item is id, second is title, etc)

        if not movie:
            print("No movie found with that ID.")
            return

        new_title = input(f"Enter new title ({movie[1]}): ") or movie[1] 
        # the part after 'or' movies[1] (the original title) is assigned to new_title if the 
        # user hits enter i.e. (input() returns an empty string '')
        # take updates similarly for other fields

        cursor.execute(
            # UPDATE ... (put your SQL to update all fields with new values)
        )
        conn.commit()
        print("Movie updated successfully.")
    except sqlite3.Error as e:
        print(f"Error updating movie: {e}")
```

# Deleting a movie

```python
def delete_movie(conn):
    # this is simple enough that researching and following previous examples should be enough
```

# Search for movie

```python
# Function to search for movies
def search_movies(conn):
    """Search for movies based on title, director, or genre"""
    # store user input into variable for search criterion e.g. "Search by title (t), director (d), or genre (g): "
    # store user input into variable for search term e.g. "Enter the search term: "

    try:
        cursor = conn.cursor()
        if criterion == "t": # criterion is variable storing the user-entered search criterion assigned above
            cursor.execute(
                "SELECT * FROM movies WHERE title LIKE ?", ("%" + search_term + "%",) 
                # search_term is storing the user-entered search term assigned above
            )
        # implement other search criteria similarly
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
```

# Display by rating

```python
# Function to display all movies sorted by rating
def display_movies_by_rating(conn):
    """Display all movies sorted by rating"""
    try:
        cursor = conn.cursor()
        # cursor.execute() and pass in SELECT statement using ORDER BY clause
        movies = cursor.fetchall() 
        # fetchall() gives you a list of lists, where each list is a movie where items are column values

        for movie in movies: # loop through list of lists
            print(movie)
    except sqlite3.Error as e:
        print(f"Error displaying movies by rating: {e}")
```

# Display movies by year

```python
def display_movies_by_year(conn):
    """Display all movies from a specific year"""
    try:
        year = int(input("Enter the release year: "))
        cursor = conn.cursor()
        # cursor.execute( ... your SQL goes here ... ) 
        # use cursor.fetchall() and store into variable movies
        # loop through movies and print each movie
    except sqlite3.Error as e:
        print(f"Error displaying movies by year: {e}")
```

# Get highest rated movie

```python
# Extra Credit Task: Get the highest-rated movie
def get_highest_rated_movie(conn):
    #
```
