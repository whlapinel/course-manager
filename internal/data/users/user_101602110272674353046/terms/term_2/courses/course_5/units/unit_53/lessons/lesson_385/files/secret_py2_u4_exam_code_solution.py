import sqlite3


# 🚀 Establish Database Connection
def create_connection():
    """Creates and returns a database connection."""
    try:
        conn = sqlite3.connect("time_travel.db")
        conn.row_factory = sqlite3.Row
        return conn
    except sqlite3.Error as e:
        print("Error connecting database:", e)


# 🔧 Create the 'travels' table (BUG: This function isn't called anywhere)
def create_table(conn):
    cursor = conn.cursor()
    cursor.execute(
        """
        CREATE TABLE IF NOT EXISTS travels (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            year INTEGER, -- BUG: should be integer
            location TEXT,
            description TEXT
        )
    """
    )
    conn.commit()


# ✍️ Add a New Time Travel Log (BUG: Missing commit)
def add_entry(conn: sqlite3.Connection, year, location, description):
    try:
        cursor = conn.cursor()
        cursor.execute(
            "INSERT INTO travels (year, location, description) VALUES (?, ?, ?)",
            (year, location, description),
        )
        print("✅ Travel entry added!")
        conn.commit()
    except sqlite3.Error as e:
        print(f"Error adding entry: {e}")


# 📜 List All Time Travel Entries (BUG: Fetching from a non-existent table)
def list_entries(conn: sqlite3.Connection):
    try:
        cursor = conn.cursor()
        cursor.execute("SELECT * FROM travels")  # BUG: Table name incorrect
        travels = cursor.fetchall()

        if not travels:
            print("No travel logs found.")
        for travel in travels:
            print(
                f"ID: {travel['id']} | Year: {travel['year']} | Location: {travel['location']} | Note: {travel['description']}"
            )
    except sqlite3.Error as e:
        print(f"Error fetching entries: {e}")


# ✏️ Update a Travel Log Entry (BUG: 'id' not referenced properly)
def update_entry(conn, travel_id, new_year, new_location, new_description):
    try:
        cursor = conn.cursor()
        cursor.execute(
            """
            UPDATE travels SET year = ?, location = ?, description = ?
            WHERE id = ?  -- BUG: Wrong column name
        """,
            (new_year, new_location, new_description, travel_id),
        )
        conn.commit()
        print("✅ Travel entry updated!")
    except sqlite3.Error as e:
        print(f"Error updating entry: {e}")


# ❌ Delete a Travel Log Entry (BUG: Missing parameter)
def delete_entry(conn, id):
    try:
        cursor = conn.cursor()
        # travel_id = input(
        #     "Enter ID to delete: "
        # )  # BUG: Should pass travel_id as an argument
        cursor.execute("DELETE FROM travels WHERE id = ?", (id,))
        conn.commit()
        print("✅ Travel entry deleted!")
    except sqlite3.Error as e:
        print(f"Error deleting entry: {e}")


# 📜 CLI Menu (BUG: Database connection closes inside the loop)
def menu():
    conn = create_connection()
    create_table(conn)  # BUG: This wasn't being called before

    while True:
        if not conn:
            break
        print("\n⏳ Time Traveler’s Archive ⏳")
        print("1. Add Time Travel Entry")
        print("2. View Travel Logs")
        print("3. Update Travel Entry")
        print("4. Delete Travel Entry")
        print("5. Exit")

        choice = input("Enter your choice: ")

        if choice == "1":
            year = input("Year of travel: ")
            location = input("Travel destination: ")
            description = input("Brief description: ")
            add_entry(
                conn, year, location, description
            )  # BUG: This function is missing a commit
            conn.commit()

        elif choice == "2":
            list_entries(conn)

        elif choice == "3":
            travel_id = int(input("ID of entry to update: "))
            new_year = input("New year: ")
            new_location = input("New location: ")
            new_description = input("New description: ")
            update_entry(conn, travel_id, new_year, new_location, new_description)

        elif choice == "4":
            travel_id = input(
                "ID of entry to delete: "
            )  # BUG: Pass travel_id as an argument
            delete_entry(conn, travel_id)

        elif choice == "5":
            print("Goodbye, time traveler!")
            conn.close()  # Correct placement
            break

        else:
            print("Invalid choice! Please try again.")


menu()
