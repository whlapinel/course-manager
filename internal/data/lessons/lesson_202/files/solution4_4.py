### **Assignment: SQLAlchemy ORM for Student Records**

#### Objective:
# Create a Python program that manages student records using SQLAlchemy ORM (Object Relational Mapping). The program should be able to perform the following operations:
# - Create a database and table.
# - Insert, query, update, and delete records.
# - Use the SQLAlchemy session to manage database transactions properly.

#### Tasks:

# 1. **Create a Database and Define the Model:**
#    - Create a new SQLite database named `school.db` using SQLAlchemy.
#    - Define a `Student` model with the following attributes:
#      - `id` (integer, primary key)
#      - `name` (string, not null)
#      - `age` (integer, not null)
#      - `grade` (string)

from sqlalchemy import create_engine, Column, Integer, String
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()

class Student(Base):
    __tablename__ = 'students'
    id = Column(Integer, primary_key=True)
    name = Column(String, nullable=False)
    age = Column(Integer, nullable=False)
    grade = Column(String)

engine = create_engine('sqlite:///school.db')
Base.metadata.create_all(engine)

# 2. **Create a Session and Insert Records:**
#    - Create a session using SQLAlchemy.
#    - Insert at least 3 student records into the `students` table using the session.
#    - Commit the changes.

#    **Example:**
#    ```python

from sqlalchemy.orm import sessionmaker

Session = sessionmaker(bind=engine)
session = Session()

student1 = Student(id=1, name='Alice', age=21, grade='A')
student2 = Student(id=2, name='Bob', age=22, grade='B')
session.add_all([student1, student2])
session.commit()

# 3. **Query the Data:**
#    - Write a function that queries all student records and prints them.
#    - Use SQLAlchemy’s query interface.

def fetch_all_students():
    students = session.query(Student).all()
    for student in students:
        print(student.id, student.name, student.age, student.grade)

# 4. **Update a Record:**
#    - Write a function that updates the grade of a specific student based on their `id`.
#    - Allow the user to input the student's `id` and the new `grade`.

def update_student_grade(student_id, new_grade):
    student = session.query(Student).filter(Student.id == student_id).first()
    if student:
        student.grade = new_grade
        session.commit()
    else:
        print(f"No student found with id {student_id}")

# 5. **Delete a Record:**
#    - Write a function that deletes a student record based on the `id`.
#    - Allow the user to input the `id` of the student they want to delete.

def delete_student(student_id):
    student = session.query(Student).filter(Student.id == student_id).first()
    if student:
        session.delete(student)
        session.commit()
    else:
        print(f"No student found with id {student_id}")

# 6. **Close the Session:**
#    - Ensure your program closes the session properly when all operations are completed.

session.close()

# ---

# ### **Submission Requirements:**
# - Submit the Python program (`.py` file) that completes all tasks using SQLAlchemy.
# - Ensure that your code is well-commented, explaining each operation.

# ### **Bonus:**
# - Add error handling for cases where a user tries to update or delete a student that doesn’t exist, as shown above.

# ---

# This assignment will assess the student’s ability to:
# - Work with SQLAlchemy to interact with an SQLite database.
# - Use the ORM approach to perform CRUD operations.
# - Understand and manage sessions effectively using SQLAlchemy.