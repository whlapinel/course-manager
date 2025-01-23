import sqlite3


def main():
    conn = create_connection("cool_db")
    create_table(conn)
    insert_pokemon(conn, "Charizard", "Water")
    insert_pokemon(conn, "Pikachu", "Electric")
    update_pokemon_type(conn, "Charizard", "Fire")
    # delete_pokemon(conn, "Charizard")
    pokemon = query_pokemon_by_name(conn, "Charizard")
    print("Name: ", pokemon[1])
    pokemons = query_all_pokemon(conn)
    print("All pokemon:")
    for poke in pokemons:
        print("Name: ", poke[1])


def create_connection(db_name: str) -> sqlite3.Connection:
    conn = sqlite3.connect(db_name)
    return conn


def create_table(conn: sqlite3.Connection):
    sql = """
        create TABLE if not exists pokemon
        (id INTEGER PRIMARY KEY,
        name TEXT,
        type TEXT)
    """
    conn.execute(sql)
    conn.commit()


def insert_pokemon(conn: sqlite3.Connection, name: str, pokemon_type: str) -> None:
    sql = """
    INSERT INTO pokemon (name, type) VALUES (?, ?)
    """
    cursor = conn.cursor()
    cursor.execute(sql, (name, pokemon_type))
    conn.commit()


def update_pokemon_type(conn: sqlite3.Connection, name: str, new_type: str) -> None:
    sql = """
    UPDATE pokemon SET type = ? WHERE name = ?
    """
    cursor = conn.cursor()
    cursor.execute(sql, (new_type, name))
    conn.commit()


def delete_pokemon(conn: sqlite3.Connection, name: str) -> None:
    sql = """
    DELETE FROM pokemon WHERE name = ?
"""
    cursor = conn.cursor()
    cursor.execute(sql, (name))
    conn.commit()


def query_pokemon_by_name(conn: sqlite3.Connection, name: str):
    sql = """
    SELECT * FROM pokemon WHERE name = ?
"""
    cursor = conn.cursor()
    cursor.execute(sql, (name,))
    conn.commit()
    pokemon = cursor.fetchone()
    print(pokemon)
    return pokemon


def query_all_pokemon(conn: sqlite3.Connection):
    sql = """
    SELECT * FROM pokemon
"""
    cursor = conn.cursor()
    cursor.execute(sql)
    conn.commit()
    rows = cursor.fetchall()
    return rows



if __name__ == "__main__":
    main()
