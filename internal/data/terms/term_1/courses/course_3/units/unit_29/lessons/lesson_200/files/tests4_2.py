### Sample Test File: `test_student_db.py`

import unittest
import sqlite3
import os

# Import the student's Python file containing their database code
# Example: from student_db import create_table, insert_student, query_students, update_student, delete_student
from solutions4_2 import create_table, insert_student, query_students, update_student, delete_student, close_connection

class TestStudentDatabase(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        # Create a test database file for the tests
        cls.db_name = 'test_school.db'
        cls.conn = sqlite3.connect(cls.db_name)
        create_table(cls.conn)  # Call the student's function to create the table

    @classmethod
    def tearDownClass(cls):
        # Close the database connection and remove the test database file after tests
        close_connection(cls.conn)
        os.remove(cls.db_name)

    def test_insert_student(self):
        # Insert a student record and check if it's correctly inserted
        insert_student(self.conn, 1, 'Alice', 21, 'A')
        insert_student(self.conn, 2, 'Bob', 22, 'B')

        cursor = self.conn.execute("SELECT * FROM students WHERE id = 1")
        student = cursor.fetchone()

        self.assertEqual(student[1], 'Alice')
        self.assertEqual(student[2], 21)
        self.assertEqual(student[3], 'A')

    def test_query_students(self):
        # Insert a few records and verify the query returns them
        insert_student(self.conn, 3, 'Charlie', 23, 'C')
        students = query_students(self.conn)

        self.assertEqual(len(students), 3)  # There should be 3 students now
        self.assertIn(('Charlie', 23, 'C'), [(s[1], s[2], s[3]) for s in students])

    def test_update_student(self):
        # Update a student's grade and check if it was updated correctly
        update_student(self.conn, 2, 'A+')

        cursor = self.conn.execute("SELECT grade FROM students WHERE id = 2")
        updated_grade = cursor.fetchone()[0]

        self.assertEqual(updated_grade, 'A+')

    def test_delete_student(self):
        # Delete a student and verify they are no longer in the database
        delete_student(self.conn, 3)

        cursor = self.conn.execute("SELECT * FROM students WHERE id = 3")
        student = cursor.fetchone()

        self.assertIsNone(student)  # Student with id = 3 should be deleted

if __name__ == '__main__':
    unittest.main()

### **Test Breakdown:**

# 1. **`setUpClass()` and `tearDownClass()`**:
#    - These methods are executed once at the start and end of the test class. 
#    - A temporary SQLite database (`test_school.db`) is created before any tests, and it's removed after all tests are done.
   
# 2. **Test `insert_student`**:
#    - Inserts students into the database.
#    - Verifies that the student is inserted correctly by querying the database and checking the values.

# 3. **Test `query_students`**:
#    - Inserts a new student and checks whether querying all students returns the correct number of students and expected details.

# 4. **Test `update_student`**:
#    - Updates a student's grade and verifies that the change was made by querying the updated record.

# 5. **Test `delete_student`**:
#    - Deletes a student record and verifies that the student is no longer in the database.

### How to Structure Your Assignment Code:

# The student’s code needs to be structured to have functions like:

# - `create_table(conn)`: Creates the `students` table.
# - `insert_student(conn, id, name, age, grade)`: Inserts a student record.
# - `query_students(conn)`: Queries and returns all students.
# - `update_student(conn, id, new_grade)`: Updates the grade of the student with the given `id`.
# - `delete_student(conn, id)`: Deletes the student with the given `id`.
# - `close_connection(conn)`: Closes the database connection.

# ### Running the Tests:

# - Make sure the student implements the required functions. 
# - Save the tests in a file (e.g., `test_student_db.py`).
# - Run the tests from the command line using:
#   ```bash
#   python -m unittest test_student_db.py