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
def create_connection(db_file: str):
    """
    Create a connection to the SQLite database.
    Hint: Use sqlite3.connect(db_file).
    """
    conn = sqlite3.connect(db_file)
    return conn


# 2. Function to create a table for storing Pokemon characters
def create_table(conn: sqlite3.Connection) -> None:
    """
    Create a table in the SQLite database for storing Pokemon.
    Hint: Use a CREATE TABLE SQL statement.
    """
    sql = """
    CREATE TABLE IF NOT EXISTS pokemon (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        type TEXT NOT NULL
    );
    """
    conn.execute(sql)
    conn.commit()


# 3. Function to insert a new Pokemon into the table
def insert_pokemon(conn: sqlite3.Connection, name: str, pokemon_type: str) -> None:
    """
    Insert a new Pokemon into the pokemon table.
    Hint: Use an INSERT INTO SQL statement.
    """
    sql = """
    INSERT INTO pokemon (name, type)
    VALUES (?, ?);
    """
    conn.execute(sql, (name, pokemon_type))
    conn.commit()


# 4. Function to update a Pokemon's type
def update_pokemon_type(conn: sqlite3.Connection, name: str, new_type: str) -> None:
    """
    Update the type of a Pokemon by name.
    Hint: Use an UPDATE SQL statement.
    """
    sql = """
    UPDATE pokemon
    SET type = ?
    WHERE name = ?;
    """
    conn.execute(sql, (new_type, name))
    conn.commit()


# 5. Function to delete a Pokemon by name
def delete_pokemon(conn, name: str) -> None:
    """
    Delete a Pokemon by its name.
    Hint: Use a DELETE FROM SQL statement.
    """
    sql = """
    DELETE FROM pokemon
    WHERE name = ?;
    """
    conn.execute(sql, (name,))
    conn.commit()


# 6. Function to query a specific Pokemon by name
def query_pokemon_by_name(conn: sqlite3.Connection, name: str):
    """
    Query a Pokemon by name.
    Hint: Use a SELECT SQL statement to find the Pokemon.
    """
    sql = """
    SELECT * FROM pokemon
    WHERE name = ?;
    """
    cursor = conn.execute(sql, (name,))
    return cursor.fetchone()


# 7. Function to query all Pokémon
def query_all_pokemon(conn: sqlite3.Connection):
    """
    Query all Pokemon from the pokemon table.
    Hint: Use a SELECT * SQL statement to fetch all rows.
    """
    sql = """
    SELECT * FROM pokemon;
    """
    cursor = conn.execute(sql)
    return cursor.fetchall()


# Call the main function to run the quiz
if __name__ == "__main__":
    main()
