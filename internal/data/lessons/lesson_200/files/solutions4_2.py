import sqlite3

# Function to create a database and table
def create_table(conn):
    """Function to create a table in the database."""
    conn.execute('''
        CREATE TABLE IF NOT EXISTS students (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL,
            age INTEGER NOT NULL,
            grade TEXT
        );
    ''')
    conn.commit()

# Function to insert a student record
def insert_student(conn, student_id, name, age, grade):
    """Function to insert a student record into the database."""
    conn.execute("INSERT INTO students (id, name, age, grade) VALUES (?, ?, ?, ?)",
                 (student_id, name, age, grade))
    conn.commit()

# Function to query all student records
def query_students(conn):
    """Function to query all student records from the database."""
    cursor = conn.execute("SELECT * FROM students")
    students = cursor.fetchall()  # Fetch all rows as a list of tuples
    for student in students:
        print(student)
    return students

# Function to update a student's grade
def update_student(conn, student_id, new_grade):
    """Function to update a student's grade based on the student's ID."""
    conn.execute("UPDATE students SET grade = ? WHERE id = ?", (new_grade, student_id))
    conn.commit()

# Function to delete a student record
def delete_student(conn, student_id):
    """Function to delete a student record based on the student's ID."""
    conn.execute("DELETE FROM students WHERE id = ?", (student_id,))
    conn.commit()

def delete_all(conn):
    """Function to delete all student records."""
    conn.execute("DELETE FROM students")
    conn.commit()

# Function to close the connection
def close_connection(conn):
    """Function to close the connection to the SQLite database."""
    conn.close()

# Main execution logic
if __name__ == "__main__":
    # Create a database connection
    conn = sqlite3.connect('school.db')

    # Create the table
    create_table(conn)

    # Insert some students
    insert_student(conn, 1, 'Alice', 21, 'A')
    insert_student(conn, 2, 'Bob', 22, 'B')
    insert_student(conn, 3, 'Charlie', 23, 'C')

    # Query and display all students
    print("All Students:")
    query_students(conn)

    # Update Bob's grade
    print("\nUpdating Bob's grade to A+")
    update_student(conn, 2, 'A+')

    # Query and display all students after the update
    print("\nAll Students after update:")
    query_students(conn)

    # Delete Charlie's record
    print("\nDeleting Charlie's record")
    delete_student(conn, 3)

    # Query and display all students after deletion
    print("\nAll Students after deletion:")
    query_students(conn)

    print("\nDeleting all students from the table")
    delete_all(conn)

    # Close the database connection
    close_connection(conn)

### **Breakdown of the Code:**

# 1. **`create_table(conn)`**:
#    - Creates the `students` table if it doesn't exist.
#    - Defines the table with columns: `id`, `name`, `age`, and `grade`.

# 2. **`insert_student(conn, student_id, name, age, grade)`**:
#    - Inserts a student record into the `students` table.
#    - Uses placeholders (`?`) to prevent SQL injection.

# 3. **`query_students(conn)`**:
#    - Fetches all rows from the `students` table.
#    - Prints each student's record.

# 4. **`update_student(conn, student_id, new_grade)`**:
#    - Updates the `grade` of a student with the specified `student_id`.

# 5. **`delete_student(conn, student_id)`**:
#    - Deletes the student with the specified `student_id` from the `students` table.

# 6. **`close_connection(conn)`**:
#    - Closes the connection to the SQLite database to ensure resources are properly released.

# 7. **Main Execution Block**:
#    - Connects to the SQLite database (`school.db`).
#    - Calls each function to demonstrate the functionality: create, insert, query, update, delete, and close.

### Example Output:
# All Students:
# (1, 'Alice', 21, 'A')
# (2, 'Bob', 22, 'B')
# (3, 'Charlie', 23, 'C')

# Updating Bob's grade to A+

# All Students after update:
# (1, 'Alice', 21, 'A')
# (2, 'Bob', 22, 'A+')
# (3, 'Charlie', 23, 'C')

# Deleting Charlie's record

# All Students after deletion:
# (1, 'Alice', 21, 'A')
# (2, 'Bob', 22, 'A+')

# ### **Explanation of the Flow**:
# 1. The `create_table()` function creates the database and table (if it doesn't already exist).
# 2. Three student records (Alice, Bob, Charlie) are inserted using the `insert_student()` function.
# 3. The students are queried and printed using the `query_students()` function.
# 4. Bob's grade is updated to "A+" using the `update_student()` function.
# 5. The updated student list is printed.
# 6. Charlie's record is deleted using the `delete_student()` function, and the final student list is displayed.
# 7. Finally, the connection is closed using `close_connection()`.