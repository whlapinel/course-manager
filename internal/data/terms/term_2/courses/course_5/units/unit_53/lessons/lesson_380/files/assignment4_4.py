### **Assignment: SQLAlchemy ORM for Student Records**
# This assignment will assess the student’s ability to:
# - Work with SQLAlchemy to interact with an SQLite database.
# - Use the ORM approach to perform CRUD operations.
# - Understand and manage sessions effectively using SQLAlchemy.
# Name: ___________

#### Objective:
# Create a Python program that manages student records using SQLAlchemy ORM (Object Relational Mapping). The program should be able to perform the following operations:
# - Create a database and table.
# - Insert, query, update, and delete records.
# - Use the SQLAlchemy session to manage database transactions properly.

# 1. **Create a Database and Define the Model:**
#    - Create a new SQLite database named `school.db` using SQLAlchemy.
#    - Define a `Student` model with the following attributes:
#      - `id` (integer, primary key)
#      - `name` (string, not null)
#      - `age` (integer, not null)
#      - `grade` (string)

# 2. **Create a Session and Insert Records:**
#    - Create a session using SQLAlchemy.
#    - Insert at least 3 student records into the `students` table using the session.
#    - Commit the changes.

# 3. **Query the Data:**
#    - Write a function that queries all student records and prints them.
#    - Use SQLAlchemy’s query interface.

# 4. **Update a Record:**
#    - Write a function that updates the grade of a specific student based on their `id`.
#    - Allow the user to input the student's `id` and the new `grade`.

# 5. **Delete a Record:**
#    - Write a function that deletes a student record based on the `id`.
#    - Allow the user to input the `id` of the student they want to delete.

# 6. **Close the Session:**
#    - Ensure your program closes the session properly when all operations are completed.

# ### **Submission Requirements:**
# - Submit the Python program (`.py` file) that completes all tasks using SQLAlchemy.
# - Ensure that your code is well-commented, explaining each operation.

# ### **Bonus:**
# - Add error handling for cases where a user tries to update or delete a student that doesn’t exist, as shown above.


