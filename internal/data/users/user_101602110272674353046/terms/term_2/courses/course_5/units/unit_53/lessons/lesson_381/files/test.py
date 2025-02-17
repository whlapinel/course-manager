import sqlite3
import pytest
from solution import (
    create_connection,
    create_table,
    add_movie,
    update_movie,
    delete_movie,
    get_highest_rated_movie,
)


# Pytest fixture to set up a temporary database connection for each test
@pytest.fixture
def db_connection():
    """Fixture to create a new in-memory database for testing."""
    conn = create_connection(":memory:")  # Use an in-memory database
    create_table(conn)
    yield conn  # Provide the connection to the test
    conn.close()  # Cleanup after the test is done


def test_add_movie(db_connection):
    """Test adding a movie to the database."""
    cursor = db_connection.cursor()
    cursor.execute(
        "INSERT INTO movies (title, genre, release_year, director, rating) VALUES (?, ?, ?, ?, ?)",
        ("Inception", "Sci-Fi", 2010, "Christopher Nolan", 9.0),
    )
    db_connection.commit()

    cursor.execute("SELECT * FROM movies WHERE title = ?", ("Inception",))
    movie = cursor.fetchone()

    assert movie is not None
    assert movie[1] == "Inception"
    assert movie[4] == "Christopher Nolan"


def test_update_movie(db_connection, monkeypatch):
    """Test updating a movie's details."""
    cursor = db_connection.cursor()
    cursor.execute(
        "INSERT INTO movies (title, genre, release_year, director, rating) VALUES (?, ?, ?, ?, ?)",
        ("Interstellar", "Sci-Fi", 2014, "Christopher Nolan", 8.6),
    )
    db_connection.commit()

    # Use monkeypatch to simulate user input during the test
    movie_id = cursor.lastrowid  # Get the ID of the inserted movie
    monkeypatch.setattr(
        "builtins.input",
        lambda prompt: {
            "Enter the ID of the movie to update: ": str(movie_id),
            f"Enter new title (Interstellar): ": "",
            f"Enter new genre (Sci-Fi): ": "Adventure",
            f"Enter new release year (2014): ": "2014",
            f"Enter new director (Christopher Nolan): ": "Christopher Nolan",
            f"Enter new rating (8.6): ": "9.0",
        }[prompt],
    )

    update_movie(db_connection)

    cursor.execute("SELECT * FROM movies WHERE id = ?", (movie_id,))
    updated_movie = cursor.fetchone()

    assert updated_movie is not None
    assert updated_movie[2] == "Adventure"
    assert updated_movie[5] == 9.0


def test_delete_movie(db_connection, monkeypatch):
    """Test deleting a movie from the database."""
    cursor = db_connection.cursor()
    cursor.execute(
        "INSERT INTO movies (title, genre, release_year, director, rating) VALUES (?, ?, ?, ?, ?)",
        ("The Matrix", "Sci-Fi", 1999, "Wachowskis", 8.7),
    )
    db_connection.commit()

    movie_id = cursor.lastrowid  # Get the ID of the inserted movie

    # Use monkeypatch to simulate user input during the test
    monkeypatch.setattr("builtins.input", lambda prompt: str(movie_id))

    delete_movie(db_connection)

    cursor.execute("SELECT * FROM movies WHERE id = ?", (movie_id,))
    deleted_movie = cursor.fetchone()

    assert deleted_movie is None


def test_get_highest_rated_movie(db_connection, capsys):
    """Test retrieving the highest-rated movie(s)."""
    cursor = db_connection.cursor()
    cursor.execute(
        "INSERT INTO movies (title, genre, release_year, director, rating) VALUES (?, ?, ?, ?, ?)",
        ("The Dark Knight", "Action", 2008, "Christopher Nolan", 9.0),
    )
    cursor.execute(
        "INSERT INTO movies (title, genre, release_year, director, rating) VALUES (?, ?, ?, ?, ?)",
        ("Inception", "Sci-Fi", 2010, "Christopher Nolan", 9.0),
    )
    cursor.execute(
        "INSERT INTO movies (title, genre, release_year, director, rating) VALUES (?, ?, ?, ?, ?)",
        ("Dunkirk", "War", 2017, "Christopher Nolan", 8.5),
    )
    db_connection.commit()

    get_highest_rated_movie(db_connection)

    captured = capsys.readouterr()
    assert "The Dark Knight" in captured.out
    assert "Inception" in captured.out
    assert (
        "Dunkirk" not in captured.out
    )  # Ensure only the highest-rated ones are displayed
