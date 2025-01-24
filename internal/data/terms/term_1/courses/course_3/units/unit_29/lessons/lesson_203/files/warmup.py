import sqlite3


def create_table(conn: sqlite3.Connection):
    cursor = conn.cursor()
    cursor.execute(
        """
    CREATE TABLE IF NOT EXISTS cars (
    id INTEGER PRIMARY KEY,
    make TEXT NOT NULL,
    model TEXT NOT NULL
    )
    """
    )
    conn.commit()


def insert_car(conn: sqlite3.Connection, make: str, model: str):
    cursor = conn.cursor()
    cursor.execute(
        """
    INSERT INTO cars (make, model) VALUES(?, ?)
    """,
        (make, model),
    )
    conn.commit()


def query_car(conn: sqlite3.Connection, make: str, model: str) -> list:
    cursor = conn.cursor()
    cursor.execute("SELECT * from cars WHERE make = ? and model = ?", (make, model))
    cars = cursor.fetchall()
    return cars


conn = sqlite3.connect("cars.db")
create_table(conn)
insert_car(conn, "Checker", "Marathon")
cars = query_car(conn, "Checker", "Marathon")
print(len(cars))
for car in cars:
    print("car: ", car)
conn.close()
