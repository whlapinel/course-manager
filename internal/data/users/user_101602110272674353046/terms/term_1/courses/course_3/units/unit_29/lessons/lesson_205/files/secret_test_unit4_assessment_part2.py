import os
import sqlite3
import pytest
from unit4.unit4_assessment_part2_solution import (
    create_connection,
    create_table,
    insert_pokemon,
    update_pokemon_type,
    delete_pokemon,
    query_pokemon_by_name,
    query_all_pokemon,
)

# Test database name
test_db = "test_pokemon.db"


# Clean up the test database before and after running the tests
@pytest.fixture(autouse=True)
def cleanup():
    # Remove the test database if it exists
    if os.path.exists(test_db):
        os.remove(test_db)
    yield
    if os.path.exists(test_db):
        os.remove(test_db)


def test_create_connection():
    # Test if the connection to the database is successful
    conn = create_connection(test_db)
    assert isinstance(conn, sqlite3.Connection), "Expected a sqlite3.Connection object"
    conn.close()


def test_create_table():
    # Test if the table is created successfully
    conn = create_connection(test_db)
    create_table(conn)

    # Verify that the table exists
    cursor = conn.execute(
        "SELECT name FROM sqlite_master WHERE type='table' AND name='pokemon';"
    )
    table_exists = cursor.fetchone()
    assert table_exists is not None, "Table 'pokemon' should exist after creation"
    conn.close()


def test_insert_pokemon():
    # Test if a Pokemon is inserted successfully
    conn = create_connection(test_db)
    create_table(conn)

    insert_pokemon(conn, "Bulbasaur", "Grass/Poison")

    # Verify that the Pokemon was inserted
    cursor = conn.execute("SELECT * FROM pokemon WHERE name='Bulbasaur';")
    result = cursor.fetchone()
    assert result is not None, "Expected Bulbasaur to be inserted into the database"
    assert result[1] == "Bulbasaur", f"Expected 'Bulbasaur', got '{result[1]}'"
    assert result[2] == "Grass/Poison", f"Expected 'Grass/Poison', got '{result[2]}'"
    conn.close()


def test_update_pokemon_type():
    # Test if a Pokemon's type is updated successfully
    conn = create_connection(test_db)
    create_table(conn)
    insert_pokemon(conn, "Pikachu", "Electric")

    update_pokemon_type(conn, "Pikachu", "Electric/Flying")

    # Verify that the type was updated
    cursor = conn.execute("SELECT type FROM pokemon WHERE name='Pikachu';")
    result = cursor.fetchone()
    assert result is not None, "Expected Pikachu to exist in the database"
    assert (
        result[0] == "Electric/Flying"
    ), f"Expected 'Electric/Flying', got '{result[0]}'"
    conn.close()


def test_delete_pokemon():
    # Test if a Pokemon is deleted successfully
    conn = create_connection(test_db)
    create_table(conn)
    insert_pokemon(conn, "Charmander", "Fire")

    delete_pokemon(conn, "Charmander")

    # Verify that the Pokemon was deleted
    cursor = conn.execute("SELECT * FROM pokemon WHERE name='Charmander';")
    result = cursor.fetchone()
    assert result is None, "Expected Charmander to be deleted from the database"
    conn.close()


def test_query_pokemon_by_name():
    # Test querying a specific Pokemon by name
    conn = create_connection(test_db)
    create_table(conn)
    insert_pokemon(conn, "Squirtle", "Water")

    squirtle = query_pokemon_by_name(conn, "Squirtle")

    assert squirtle is not None, "Expected to find Squirtle in the database"
    assert squirtle[1] == "Squirtle", f"Expected 'Squirtle', got '{squirtle[1]}'"
    assert squirtle[2] == "Water", f"Expected 'Water', got '{squirtle[2]}'"
    conn.close()


def test_query_all_pokemon():
    # Test querying all Pokémon in the database
    conn = create_connection(test_db)
    create_table(conn)
    insert_pokemon(conn, "Bulbasaur", "Grass/Poison")
    insert_pokemon(conn, "Charmander", "Fire")

    pokemons = query_all_pokemon(conn)

    assert len(pokemons) == 2, f"Expected 2 Pokémon, got {len(pokemons)}"
    assert (
        pokemons[0][1] == "Bulbasaur"
    ), f"Expected 'Bulbasaur', got '{pokemons[0][1]}'"
    assert (
        pokemons[1][1] == "Charmander"
    ), f"Expected 'Charmander', got '{pokemons[1][1]}'"
    conn.close()
