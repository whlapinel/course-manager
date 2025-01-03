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

Write a function that accepts two numbers a and b and returns True if a is evenly divisible by b

# Agenda

# Looking ahead

- Unit 5 Test Thursday 10/24
- Midterm: Monday 10/28

# Introduction to Object-Oriented Programming (OOP) in Python

- **OOP** is a programming paradigm that uses objects and classes.
- Key concepts:
  - **Class**: Blueprint for creating objects.
  - **Object**: Instance of a class.

<!-- ### Key Concepts Covered:
1. **What is a Class?**
2. **What is an Object?**
3. **The `__init__()` Method (Constructor)**
4. **Attributes: Instance vs. Class**
5. **Methods**
6. **The `self` Keyword**
7. **Inheritance**
8. **Encapsulation**
9. **Polymorphism**
10. **Abstraction** (briefly mentioned as optional) -->

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

# Encapsulation

- **Encapsulation** restricts direct access to some of an object’s components.
- Use **private** attributes by prefixing them with `__`.

**Example:**

```python
class Person:
    def __init__(self, name):
        self.__name = name  # Private attribute

    def get_name(self):
        return self.__name  # Getter method
```

# Polymorphism

- **Polymorphism** allows different classes to be treated the same way through a common interface.
- Achieved by overriding methods in different classes.

**Example:**

```python
class Cat:
    def speak(self):
        return "Meow"

class Dog:
    def speak(self):
        return "Woof"

def animal_sound(animal):
    print(animal.speak())

animal_sound(Cat())  # Output: Meow
animal_sound(Dog())  # Output: Woof
```

# Abstraction

- **Abstraction** hides complex details and shows only the essential features.
- In Python, we can use **abstract base classes** via the `abc` module (optional for this intro).

# Summary of OOP Concepts

- **Class**: Blueprint for objects.
- **Object**: Instance of a class.
- **Attributes**: Data stored in an object.
- **Methods**: Functions defined inside a class.
- **Inheritance**: One class inherits from another.
- **Encapsulation**: Restrict access to some attributes.
- **Polymorphism**: Different classes with the same interface.

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
