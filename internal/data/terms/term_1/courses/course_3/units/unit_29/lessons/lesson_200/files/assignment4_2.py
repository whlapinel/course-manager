# Name:
# ### **Assignment: SQLite Database for Student Records**
# This assignment will assess the student’s ability to:
# - Create and manage a SQLite database.
# - Use SQL to insert, update, delete, and query records.
# - Work with SQLite in Python using the `sqlite3` module.

# #### Objective:
# Create a Python program that manages student records using SQLite. The program should be able to perform the following operations:
# - Create a database and table.
# - Insert student records.
# - Query, update, and delete records.
# - Close the connection properly.

# ### **Submission Requirements:**
# - Put all the function definitions in the global scope (not indented inside any other block of code).
# - Put all your function calls in the main block of the program "if __name__ == "__main__":"
# - Include a docstring for each function, explaining each operation.
# - Ensure the program runs without errors and performs the specified operations.
# - Submit the Python program (`.py` file) that completes all the tasks outlined above.

# ### **Bonus:**
# - Add error handling for cases where the user tries to update or delete a student that doesn’t exist.

# #### Tasks:

# 1. **Create a Database:**
#    - Write a Python program that creates an SQLite database named `school.db`.
#    - Inside this database, create a table called `students` with the following columns:
#      - `id` (integer, primary key)
#      - `name` (text, not null)
#      - `age` (integer, not null)
#      - `grade` (text)
def create_table(conn):
    """Function to create the students table."""
    raise NotImplementedError

# 2. **Insert Records:**
#    - Add at least 3 student records into the `students` table. For each student, provide an `id`, `name`, `age`, and `grade`.
#    - Commit the changes after inserting the records.
def insert_student(conn, student_id, name, age, grade):
    """Function to insert a student record into the students table."""
    raise NotImplementedError

# 3. **Query the Data:**
#    - Write a function that queries the `students` table and prints all student records. 
def query_students(conn):
    """Function to query all student records from the database."""
    raise NotImplementedError

# 4. **Update a Record:**
#    - Write a function that updates the grade of a specific student based on their `id`.
#    - Allow the user to input the student's `id` and the new `grade`.
def update_student(conn, student_id, new_grade):
    """Function to update a student's grade based on the student's ID."""
    raise NotImplementedError

# 5. **Delete a Record:**
#    - Write a function that deletes a student record based on the `id`.
#    - Allow the user to input the `id` of the student they want to delete.
def delete_student(conn, student_id):
    """Function to delete a student record based on the student's ID."""
    raise NotImplementedError

# 6. **Delete All Records:**
def delete_all(conn):
    """Function to delete all student records."""
    raise NotImplementedError

# 7. **Close the Connection:**
#    - Ensure the program closes the connection to the database once all operations are done.
def close_connection(conn):
    """Function to close the connection to the SQLite database."""
    raise NotImplementedError

# ### **Main Program:**
if __name__ == "__main__":
    raise NotImplementedError # replace this line with your code, including each of the steps below

    # Create a database connection
    
    # Create the table (complete the function above, then call it here)

    # Insert some students (complete the function above then call the insert_student function for at least 3 students)

    # Query and display all students (complete the function above then call it here)

    # Update grade for at least one student (complete the function above then call it here)

    # Query and display all students after the update (complete the function above then call it here)

    # Delete one student's record (complete the function above then call it here)

    # Query and display all students after deletion (complete the function above then call it here)

    # Delete all students from the table so that you can run repeatedly without errors (complete the function above then call it here)

    # Close the database connection (complete the function above then call it here)
