# Assignment 5.2: Advanced OOP in Python

---

## 1. **Encapsulation**
Create a `BankAccount` class with an attribute `_balance` and a method `get_balance()` to return the balance.  
- Try to access `_balance` directly.  
- Modify the balance using a `set_balance()` method.

---

## 2. **Getters and Setters**
Create a `Student` class with a private attribute `_name`.  
- Use `@property` to create a getter for `name`.  
- Use `@name.setter` to create a setter that sets the name.  

---

## 3. **Class and Static Methods**
Create a `Calculator` class with:  
- A `@staticmethod` called `add(x, y)` that returns the sum.  
- A `@classmethod` called `info()` that prints "This is a calculator."  

---

## 4. **Polymorphism**
Create a `Vehicle` class with a method `move()`.  
Create two subclasses `Car` and `Bike` that override `move()`.  
- Create objects of both and call `move()`.

---

## 5. **Inheritance + Overriding**
Create a `Shape` class with a method `area()` that returns `0`.  
Create a `Square` subclass that overrides `area()` to return the square's area.  

---

## 6. **Bonus**
Create a `Pizza` class with an attribute `toppings` (a list).  
- Use a class method to return a list of all pizzas created.  
- Use a static method to check if a topping is vegetarian.  
- Example: `"pepperoni" -> False`, `"mushroom" -> True`.
