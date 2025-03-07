# **Python 2: Unit 4 Exam**  

The **Time Traveler’s Archive** CLI allows users to log historical events they’ve “visited” and manage their archive. But it has some bugs! 🪳

## **📜 Instructions**  
Your task is to **debug** and **fix** the following Python program. When fixed, it should:  
- Store time travel records in an SQLite3 database.  
- Allow users to add, view, update, and delete records.  
- Handle errors and incorrect inputs properly.  
- Ensure the database is correctly set up and persists data.  

---

### **🚨 Buggy Code Below 🚨**
```python
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
    cursor.execute("""
        CREATE TABLE IF NOT EXISTS travels (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            year TEXT,  -- BUG: Should be INTEGER?
            location TEXT,
            description TEXT
        )
    """)
    conn.commit()

# ✍️ Add a New Time Travel Log (BUG: Missing commit)
def add_entry(conn, year, location, description):
    try:
        cursor = conn.cursor()
        cursor.execute("INSERT INTO travels (year, location, description) VALUES (?, ?, ?)", 
                       (year, location, description))
        print("✅ Travel entry added!")
    except sqlite3.Error as e:
        print(f"Error adding entry: {e}")

# 📜 List All Time Travel Entries (BUG: Fetching from a non-existent table)
def list_entries(conn):
    try:
        cursor = conn.cursor()
        cursor.execute("SELECT * FROM time_travels")  # BUG: Table name incorrect
        travels = cursor.fetchall()

        if not travels:
            print("No travel logs found.")
        for travel in travels:
            print(f"ID: {travel['id']} | Year: {travel['year']} | Location: {travel['location']} | Note: {travel['description']}")
    except sqlite3.Error as e:
        print(f"Error fetching entries: {e}")

# ✏️ Update a Travel Log Entry (BUG: 'id' not referenced properly)
def update_entry(conn, travel_id, new_year, new_location, new_description):
    try:
        cursor = conn.cursor()
        cursor.execute("""
            UPDATE travels SET year = ?, location = ?, description = ?
            WHERE travel_id = ?  -- BUG: Wrong column name
        """, (new_year, new_location, new_description, travel_id))
        conn.commit()
        print("✅ Travel entry updated!")
    except sqlite3.Error as e:
        print(f"Error updating entry: {e}")

# ❌ Delete a Travel Log Entry (BUG: Missing parameter)
def delete_entry(conn):
    try:
        cursor = conn.cursor()
        travel_id = input("Enter ID to delete: ")  # BUG: Should pass travel_id as an argument
        cursor.execute("DELETE FROM travels WHERE id = ?", (travel_id,))
        conn.commit()
        print("✅ Travel entry deleted!")
    except sqlite3.Error as e:
        print(f"Error deleting entry: {e}")

# 📜 CLI Menu (BUG: Database connection closes inside the loop)
def menu():
    conn = create_connection()
    create_table(conn)  # BUG: This wasn't being called before

    while True:
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
            add_entry(conn, year, location, description)  # BUG: This function is missing a commit

        elif choice == "2":
            list_entries(conn)

        elif choice == "3":
            travel_id = int(input("ID of entry to update: "))
            new_year = input("New year: ")
            new_location = input("New location: ")
            new_description = input("New description: ")
            update_entry(conn, travel_id, new_year, new_location, new_description)

        elif choice == "4":
            travel_id = input("ID of entry to delete: ")  # BUG: Pass travel_id as an argument
            delete_entry(conn, travel_id)

        elif choice == "5":
            print("Goodbye, time traveler!")
            conn.close()  # Correct placement
            break

        else:
            print("Invalid choice! Please try again.")

menu()
```

---

## **🛠 Expected Fixes**  

Students should:
1. **Call** `create_table(conn)` inside `menu()` to ensure the table is created.  
2. **Fix the incorrect table name** in `list_entries()`.  
3. **Change the column type** of `year` from `TEXT` to `INTEGER` in `create_table()`.  
4. **Add `conn.commit()`** to `add_entry()` so changes persist.  
5. **Fix the incorrect `WHERE` clause** in `update_entry()`, replacing `travel_id` with `id`.  
6. **Pass `travel_id` as an argument** to `delete_entry()` and remove redundant `input()`.  
7. **Ensure database connection remains open** during the menu loop but closes properly at exit.  

---

### **✅ Expected Output When Fixed**
```
⏳ Time Traveler’s Archive ⏳
1. Add Time Travel Entry
2. View Travel Logs
3. Update Travel Entry
4. Delete Travel Entry
5. Exit

Enter your choice: 1
Year of travel: 1776
Travel destination: Philadelphia
Brief description: Witnessed the signing of the Declaration of Independence.
✅ Travel entry added!

Enter your choice: 2
ID: 1 | Year: 1776 | Location: Philadelphia | Note: Witnessed the signing of the Declaration of Independence
```

---

### **📜 Grading Criteria**  

| **Task**                          | **Points** |
|------------------------------------|-----------|
| Calls `create_table(conn)`         | 5         |
| Fixes `list_entries()` table name  | 5         |
| Changes `year` type to INTEGER     | 5         |
| Adds `conn.commit()` to `add_entry()` | 5         |
| Fixes `WHERE` clause in `update_entry()` | 5 |
| Passes `travel_id` properly in `delete_entry()` | 5 |
| Handles database connection properly | 5 |
| Overall program fully functional   | 10        |
