---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Building a Clash Royale-Themed CLI App
## Using SQLite3 in Python

## 🏰 Project Overview

### **Clash Royale Card Manager**
- Manage a database of Clash Royale cards.
- Add, update, delete, and list cards.
- Built using Python and SQLite3.

# 🛠 Setting Up the Database

```python
import sqlite3

def create_connection():
    """Creates and returns a database connection."""
    try:
        conn = sqlite3.connect("clash_royale.db")
        conn.row_factory = sqlite3.Row  # Enables dictionary-like access to rows
        return conn
    except sqlite3.Error as e:
        print(f"Database error: {e}")
        return None
```

# 🎴 Adding Cards (With Error Handling)

```python
def add_card(conn, name, elixir_cost, rarity):
    try:
        cursor = conn.cursor()
        cursor.execute("INSERT INTO cards (name, elixir_cost, rarity) VALUES (?, ?, ?)",
                       (name, elixir_cost, rarity))
        conn.commit()
        print("✅ Card added successfully!")
    except sqlite3.Error as e:
        print(f"Error adding card: {e}")
```

Usage:
```python
conn = create_connection()
if conn:
    add_card(conn, "P.E.K.K.A", 7, "Epic")
    conn.close()
```

# 📜 Viewing Cards (With Row Factory)

```python
def list_cards(conn):
    try:
        cursor = conn.cursor()
        cursor.execute("SELECT * FROM cards")
        cards = cursor.fetchall()
        
        for card in cards:
            print(f"ID: {card['id']}, Name: {card['name']}, Elixir: {card['elixir_cost']}, Rarity: {card['rarity']}")
    except sqlite3.Error as e:
        print(f"Error fetching cards: {e}")
```

Usage:
```python
conn = create_connection()
if conn:
    list_cards(conn)
    conn.close()
```

# 🔄 Updating a Card (With Error Handling)

```python
def update_card(conn, card_id, new_name, new_elixir, new_rarity):
    try:
        cursor = conn.cursor()
        cursor.execute("UPDATE cards SET name = ?, elixir_cost = ?, rarity = ? WHERE id = ?",
                       (new_name, new_elixir, new_rarity, card_id))
        conn.commit()
        print("✅ Card updated successfully!")
    except sqlite3.Error as e:
        print(f"Error updating card: {e}")
```

# ❌ Deleting a Card (With Error Handling)

```python
def delete_card(conn, card_id):
    try:
        cursor = conn.cursor()
        cursor.execute("DELETE FROM cards WHERE id = ?", (card_id,))
        conn.commit()
        print("✅ Card deleted successfully!")
    except sqlite3.Error as e:
        print(f"Error deleting card: {e}")
```

# 🎮 Building the CLI Menu (With Error Handling)

```python
def menu():
    conn = create_connection()
    if not conn:
        return
    try:
        while True:
            print("\n🏆 Clash Royale Card Manager 🏆")
            print("1. Add Card")
            print("2. View Cards")
            print("3. Update Card")
            print("4. Delete Card")
            print("5. Exit")
            
            choice = input("Enter choice: ")
            
            if choice == "1":
                name = input("Card Name: ")
                elixir = int(input("Elixir Cost: "))
                rarity = input("Rarity (Common/Rare/Epic/Legendary): ")
                add_card(conn, name, elixir, rarity)
            elif choice == "2":
                list_cards(conn)
            elif choice == "3":
                card_id = int(input("Card ID: "))
                name = input("New Name: ")
                elixir = int(input("New Elixir Cost: "))
                rarity = input("New Rarity: ")
                update_card(conn, card_id, name, elixir, rarity)
            elif choice == "4":
                card_id = int(input("Card ID to delete: "))
                delete_card(conn, card_id)
            elif choice == "5":
                break
            else:
                print("Invalid choice!")
    finally:
        conn.close()
```

# 💡 SQLite3 Advanced Tips & Tricks

## 🔹 Using `sqlite3.Row` for Better Data Handling
- Allows accessing columns by name (e.g., `row['name']` instead of `row[1]`).
- Makes code **more readable** and less error-prone.

# 🔹 PRAGMA Statements
- Control SQLite settings.
- **Example: Enable Foreign Keys**
  ```python
  conn.execute("PRAGMA foreign_keys = ON;")
  ```
- **Example: Check Database Integrity**
  ```python
  conn.execute("PRAGMA integrity_check;")
  ```

# ⚡ Lesser-Known SQLite Features

## 🔸 Use an **In-Memory Database** for Fast Testing
```python
conn = sqlite3.connect(":memory:")
```
- Faster than file-based databases.
- Data is lost when the program exits.

# 🔸 Using SQLite Functions in Queries
```python
cursor.execute("SELECT upper(name) FROM cards")
```
- Runs `UPPER(name)` inside SQLite instead of Python.

# 🔸 Storing JSON Data
```python
cursor.execute("INSERT INTO cards (name, elixir_cost, rarity) VALUES (?, ?, ?)",
               (json.dumps({"name": "P.E.K.K.A", "elixir": 7}), 7, "Epic"))
```
- Useful for dynamic attributes.

# 🚀 Wrapping Up
- You’ve built a **Clash Royale-themed CLI app** with SQLite3.
- You learned **error handling, row_factory, and advanced SQLite features**.
- Try adding **custom queries** and new features!

# Assignment 4.7

- Build your own CLI application using Python and Sqlite3

# Ideas for your app

## To-Do List Manager

Users can add, delete, and mark tasks as complete.
Store tasks in an SQLite database with due dates and statuses.

# Book Tracker

Users can add books they’ve read, rate them, and track progress.
Store book titles, authors, completion status, and ratings in SQLite.

# Expense Tracker

Users can log daily expenses with categories.
Store transactions in SQLite with timestamps and amounts.

# Simple Contact Manager

Users can store names, phone numbers, and emails.
Provide search and update functionality using SQLite.

# Student Grades Manager

Allow users to enter student names, subjects, and grades.
Compute average grades and store records in SQLite.
