---
marp: true
---
<!-- headingDivider: 1 -->

# Warmup

Write a function `halve` that returns the first half of the string passed in.

```python
# function signature
def halve(s: str)->str:
    pass # replace 'pass' with your code

# test statements
assert halve('oooo') == 'oo'
assert halve("I don't want to be a king") == "I don't want"
```

# Agenda

- Lecture on SQLAlchemy
- Assignment 4.4 SQLAlchemy

# Glance-ahead

- Tomorrow (Thurs): Classes and Objects
- Friday: More classes and objects
- Unit 4 Test will be Oct. 8th

# Announcements

- Return signed permission forms for Certification Exam

# Introduction to SQLAlchemy

# Overview

- **SQLAlchemy** is a Python SQL toolkit and Object Relational Mapper (ORM).
- Allows interaction with databases using Python objects.
- Supports multiple database backends (e.g., SQLite, PostgreSQL, MySQL).

# Why Use SQLAlchemy?

- Simplifies database interaction through ORM.
- Provides flexibility for both SQL and Pythonic database access.
- Supports complex queries while maintaining performance.

# Key Components

- **Engine**: Core interface to the database.
- **Session**: Manages database transactions and queries.
- **Declarative Base**: Defines the database schema using Python classes.

# Setting Up SQLAlchemy

- Install SQLAlchemy using pip:

  ```bash
  pip install sqlalchemy
  ```

# Creating a Database

1. Import the necessary modules:

   ```python
   from sqlalchemy import create_engine, Column, Integer, String
   from sqlalchemy.ext.declarative import declarative_base
   ```

2. Create the engine and base:

   ```python
   engine = create_engine('sqlite:///example.db')
   Base = declarative_base()
   ```

# Defining a Model

- Example of defining a 'Student' model:

  ```python
  class Student(Base):
      __tablename__ = 'students'
      id = Column(Integer, primary_key=True)
      name = Column(String, nullable=False)
      age = Column(Integer, nullable=False)
      grade = Column(String)
  ```

- `Base.metadata.create_all(engine)` creates the table in the database.

# Working with Sessions

- A **Session** allows you to interact with the database:

  ```python
  from sqlalchemy.orm import sessionmaker
  Session = sessionmaker(bind=engine)
  session = Session()
  ```

# CRUD Operations

- **Insert**:

  ```python
  student = Student(id=1, name='Alice', age=21, grade='A')
  session.add(student)
  session.commit()
  ```

# **Query**

  ```python
  students = session.query(Student).all()
  for student in students:
      print(student.name)
  ```

# **Update**

  ```python
  student = session.query(Student).filter(Student.id == 1).first()
  student.grade = 'A+'
  session.commit()
  ```

# **Delete**
  
  ```python
  session.delete(student)
  session.commit()
  ```

# Best Practices

- Always close the session using `session.close()`.
- Use exception handling for database operations.

# Summary

- SQLAlchemy is a powerful tool for managing databases with Python.
- ORM simplifies data interaction by treating tables as Python objects.
- Practice CRUD operations to gain confidence in using SQLAlchemy.
