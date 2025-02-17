import sqlite3


# Main function to execute the quiz
def main():
    # Create a database connection
    conn = create_connection("pokemon.db")

    # Create the Pokemon table
    create_table(conn)

    # Insert a new Pokemon
    insert_pokemon(conn, "Pikachu", "Electric")
    insert_pokemon(conn, "Charmander", "Fire")

    # Update an existing Pokemon's type
    update_pokemon_type(conn, "Pikachu", "Electric/Flying")

    # Query and display a specific Pokemon
    pikachu = query_pokemon_by_name(conn, "Pikachu")
    print(f"Queried Pokemon: {pikachu}")

    # Query and display all Pokémon
    pokemons = query_all_pokemon(conn)
    print(f"All Pokemon: {pokemons}")

    # Delete a Pokemon
    delete_pokemon(conn, "Charmander")

    # Close the connection
    conn.close()


# 1. Function to create a connection to the SQLite database
def create_connection(db_file: str) -> sqlite3.Connection:
    """
    Create a connection to the SQLite database.
    Hint: Use sqlite3.connect(db_file).
    """
    raise NotImplementedError


# 2. Function to create a table for storing Pokemon characters
def create_table(conn: sqlite3.Connection) -> None:
    """
    Create a table in the SQLite database for storing Pokemon.
    Hint: Use a CREATE TABLE SQL statement.
    """
    raise NotImplementedError


# 3. Function to insert a new Pokemon into the table
def insert_pokemon(conn, name: str, pokemon_type: str) -> None:
    """
    Insert a new Pokemon into the pokemon table.
    Hint: Use an INSERT INTO SQL statement.
    """
    raise NotImplementedError


# 4. Function to update a Pokemon's type
def update_pokemon_type(conn, name: str, new_type: str) -> None:
    """
    Update the type of a Pokemon by name.
    Hint: Use an UPDATE SQL statement.
    """
    raise NotImplementedError


# 5. Function to delete a Pokemon by name
def delete_pokemon(conn, name: str) -> None:
    """
    Delete a Pokemon by its name.
    Hint: Use a DELETE FROM SQL statement.
    """
    raise NotImplementedError


# 6. Function to query a specific Pokemon by name
def query_pokemon_by_name(conn, name: str):
    """
    Query a Pokemon by name.
    Hint: Use a SELECT SQL statement and cursor.fetchone() to find the Pokemon.
    """
    raise NotImplementedError


# 7. Function to query all Pokémon
def query_all_pokemon(conn):
    """
    Query all Pokemon from the pokemon table.
    Hint: Use a SELECT * SQL statement and cursor.fetchall() to fetch all rows.
    """
    raise NotImplementedError


# Call the main function to run the quiz
if __name__ == "__main__":
    main()

