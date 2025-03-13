---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Lesson 5.2 OOP in Python (Part 2)

# Warmup

Write a `Rectangle` class with `width` and `height` attributes.  
Add a method `area()` that returns the area of the rectangle.  

# Agenda

1. Recap of OOP concepts
2. Encapsulation
3. Getters and Setters
4. Class methods and Static methods
5. Polymorphism

# Recap: What We Learned

- Classes and Objects
- `__init__` constructor
- Instance and class attributes
- Methods and `self`
- Inheritance

# Encapsulation

- **Encapsulation** protects data by restricting access to certain components.
- Use underscores `_` for internal attributes and methods.

**Example:**
```python
class BankAccount:
    def __init__(self, balance):
        self._balance = balance

    def get_balance(self):
        return self._balance
```

# Getters and Setters

- Use **getters** and **setters** to access and modify private attributes.

**Example:**
```python
class Person:
    def __init__(self, name):
        self._name = name
    
    def get_name(self):
        return self._name
    
    def set_name(self, name):
        self._name = name
```

# `@property` Decorator

- A Pythonic way to create getters and setters.

**Example:**
```python
class Car:
    def __init__(self, color):
        self._color = color
    
    @property
    def color(self):
        return self._color
    
    @color.setter
    def color(self, value):
        self._color = value
```

# Class Methods and Static Methods

## **Class Method**  
- Uses `@classmethod`
- Works with the class itself, not an instance.
  
### **Static Method**  
- Uses `@staticmethod`
- Doesn’t need `self` or `cls`.

**Example:**
```python
class Math:
    @staticmethod
    def add(x, y):
        return x + y
```

# Polymorphism

- **Polymorphism** allows objects of different classes to be treated as objects of a common superclass.
- Achieved through method overriding.

**Example:**
```python
class Animal:
    def speak(self):
        print("Animal speaks")

class Dog(Animal):
    def speak(self):
        print("Dog barks")

animals = [Animal(), Dog()]
for animal in animals:
    animal.speak()
```

# Exercise: Create a Class

1. Create a `Rectangle` class with `width` and `height`.
2. Add methods to calculate area and perimeter.
3. Use `@property` to create getter and setter for `width`.

# Final Thoughts

- Encapsulation helps protect data.
- Use `@property`, `@classmethod`, and `@staticmethod` appropriately.
- Polymorphism makes your code more flexible.
