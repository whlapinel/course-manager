---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Lesson 5.1 Intro to OOP in Python

# Warmup

Write a function that accepts two numbers a and b and returns True if a is evenly divisible by b

# Agenda

# Looking ahead

- PE2 Module 3 Quiz due Friday
- Unit 5 Exam on Thursday 3/20
- Midterm on 3/25 or 3/26

# Introduction to Object-Oriented Programming (OOP) in Python

- **OOP** is a programming paradigm that uses objects and classes.
- Key concepts:
  - **Class**: Blueprint for creating objects.
  - **Object**: Instance of a class.

# What is a Class?

- A **class** is a blueprint for creating objects (instances).
- It defines attributes (data) and methods (functions).
  
**Syntax:**

```python
class MyClass:
    # Constructor
    def __init__(self, attribute1):
        self.attribute1 = attribute1

    # Method
    def my_method(self):
        print(self.attribute1)
```

# What is an Object?

- An **object** is an instance of a class.
- Objects can have **attributes** (data) and **methods** (behavior).
  
**Example:**

```python
# Create an object
my_obj = MyClass('Hello')
my_obj.my_method()  # Output: Hello
```

# The `__init__()` Method

- The `__init__()` method is a **constructor** used to initialize an object’s attributes.
- It runs automatically when a new object is created.

**Example:**

```python
class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age

p1 = Person("Alice", 30)
```

# Attributes in Classes

- **Instance attributes**: Defined inside the constructor (`__init__`), unique to each object.
- **Class attributes**: Defined at the class level, shared by all instances.

**Example:**

```python
class Dog:
    species = 'Canine'  # Class attribute
    
    def __init__(self, name):
        self.name = name  # Instance attribute
```

# Methods in Classes

- **Instance methods**: Functions defined in a class, and take `self` as the first argument.
- They can access and modify object attributes.

**Example:**

```python
class Dog:
    def bark(self):
        print(f"{self.name} is barking!")

dog = Dog("Buddy")
dog.bark()  # Output: Buddy is barking!
```

# The `self` Keyword

- `self` refers to the current instance of the class.
- It allows access to the object’s attributes and methods.

**Example:**

```python
class Car:
    def __init__(self, model):
        self.model = model
    
    def show_model(self):
        print(f"The car model is {self.model}")
```

# Inheritance in Python

- Inheritance allows a class to inherit attributes and methods from another class.
- Helps reuse code and create relationships between classes.
- Caution! Use is typically discouraged nowadays due to potential for complexity and tight coupling

```python
class ParentClass:
    pass

class ChildClass(ParentClass):
    pass
```

**Example:**

```python
class Animal:
    def speak(self):
        print("Animal speaks")

class Dog(Animal):
    def speak(self):
        print("Dog barks")
```


# Exercise: Create a Class

1. Define a `Car` class.
2. Add attributes `make` and `year`.
3. Add a method `car_info()` that prints the car's make and year.
4. Create an instance of the class and call the `car_info()` method.

```python
class Car:
    def __init__(self, make, year):
        self.make = make
        self.year = year

    def car_info(self):
        print(f"Car: {self.make}, Year: {self.year}")

my_car = Car("Toyota", 2020)
my_car.car_info()
```

# Final Thoughts

- Object-Oriented Programming helps structure complex programs by organizing data and functions into **objects**.
- Practice by creating your own classes, methods, and utilizing the power of OOP concepts like inheritance and encapsulation.
