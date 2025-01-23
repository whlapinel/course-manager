---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Warmup

Write a function `write_name` that writes a name to a text file `names.txt`.

```python
def write_name(name: str)->None:
```

Write a second function `read_name` that reads and returns the first name from the text file `names.txt`.

```python
def read_name()->str:
```

Add your test:

```python
write_name("Bill")
name = read_name()
assert name == "Bill"
```

# Agenda

- Review for Unit 4 test (tomorrow)

# Announcements

- Friday: Guest Speakers from MSC 8-9am in Cafe

# Looking ahead

## Unit 5: Classes and Objects, Generators

# Python File Operations & SQLite Review

1. File Operations
   - Writing to files
   - Appending to files
   - Reading from files (Total and Average)

2. SQLite Operations
   - Creating a connection
   - Creating a table
   - Inserting, Updating, Deleting
   - Querying (single and multiple)

# File Operations: Writing

## Task 1: Writing to a File

Write a function that takes in a **filename** and a **name** and writes the name to the file.  

*Hints*:

- Use the `open()` function with `'w'` mode to write.
- Don’t forget to close the file or use a `with` block!

```python
def write_name_to_file(filename, name):
    # Your code here
```

# File Operations: Appending

## Task 2: Appending to a File

Write a function that **appends** a new grade to a file of scores.

*Hints*:

- Use `'a'` mode in `open()` to append.
- Don’t forget to add a newline (`\n`) after the grade!

```python
def append_grade_to_file(filename, grade):
    # Your code here
```

# File Operations: Reading & Total

## Task 3: Reading and Calculating Total

Write a function that reads from the **scores file** and calculates the **total** of all scores.

*Hints*:

- Open the file in `'r'` mode.
- Use a loop to read each line, convert to an integer, and sum the values.

```python
def read_and_calculate_total(filename):
    # Your code here
```

# File Operations: Reading & Average

## Task 4: Calculating Average

Write a function to calculate the **average score** from a file of scores.

*Hints*:

- Use a loop to calculate both the total and the count of scores.
- Return the total divided by the number of scores.

```python
def read_and_calculate_average(filename):
    # Your code here
```

# SQLite Operations: Connection

## Task 5: Create a Database Connection

Write a function that creates a connection to an SQLite database.

*Hints*:

- Use `sqlite3.connect()` to create a connection.
- Don’t forget to return the connection object!

```python
def create_connection(db_file):
    # Your code here
```

# SQLite Operations: Create Table

## Task 6: Creating a Table for Pokémon

Write a function to create a **Pokemon** table with columns `id`, `name`, and `type`.

*Hints*:

- Use `CREATE TABLE` in your SQL statement.
- Execute the SQL using the connection's `execute()` method.

```python
def create_table(conn):
    # Your code here
```

# SQLite Operations: Insert Pokémon

## Task 7: Inserting a Pokémon

Write a function that inserts a Pokémon (name and type) into the table.

*Hints*:

- Use `INSERT INTO` SQL.
- Pass the values as parameters to avoid SQL injection.

```python
def insert_pokemon(conn, name, pokemon_type):
    # Your code here
```

# SQLite Operations: Update Pokémon

## Task 8: Updating a Pokémon's Type

Write a function that updates a Pokémon's **type** in the table.

*Hints*:

- Use `UPDATE` SQL to change the type where the name matches.
- Don't forget to commit the change!

```python
def update_pokemon_type(conn, name, new_type):
    # Your code here
```

# SQLite Operations: Delete Pokémon

## Task 9: Deleting a Pokémon

Write a function that deletes a Pokémon from the table by name.

*Hints*:

- Use `DELETE FROM` SQL with a `WHERE` condition.
- Make sure to commit the changes.

```python
def delete_pokemon(conn, name):
    # Your code here
```

# SQLite Operations: Query Pokémon

## Task 10: Query a Pokémon by Name

Write a function to query a specific Pokémon by its name and return the result.

*Hints*:

- Use `SELECT` SQL with a `WHERE` condition.
- Return the fetched result using `fetchone()`.

```python
def query_pokemon_by_name(conn, name):
    # Your code here
```

# SQLite Operations: Query All Pokémon

## Task 11: Query All Pokémon

Write a function to query all Pokémon from the table.

*Hints*:

- Use `SELECT *` SQL to fetch all rows.
- Return all rows using `fetchall()`.

```python
def query_all_pokemon(conn):
    # Your code here
```

# Conclusion

