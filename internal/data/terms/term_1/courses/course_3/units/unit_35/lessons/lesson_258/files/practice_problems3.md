# Practice Problems 9.3 (Object-Oriented Programming)

## **1. Class Definition**

**Question**: What is the output of the following code?

```python
class Animal:
    def __init__(self, name):
        self.name = name

cat = Animal("Whiskers")
print(cat.name)
```

- A) `"Whiskers"`
- B) `cat`
- C) `"Animal"`
- D) `None`

## **2. Encapsulation**

**Question**: What will this code output?

```python
class Animal:
    def __init__(self):
        self.__hidden = "Secret"

    def reveal_secret(self):
        return self.__hidden

animal = Animal()
print(animal.reveal_secret())
```

- A) `Secret`
- B) `None`
- C) `Error`
- D) `"__hidden"`

## **3. Inheritance**

**Question**: What is the output of this code?

```python
class Parent:
    def speak(self):
        return "Parent says hello"

class Child(Parent):
    pass

c = Child()
print(c.speak())
```

- A) `"Child says hello"`
- B) `"Parent says hello"`
- C) `Error`
- D) `None`

## **4. Method Overriding**

**Question**: What will this code output?

```python
class Parent:
    def greet(self):
        return "Hello from Parent"

class Child(Parent):
    def greet(self):
        return "Hello from Child"

c = Child()
print(c.greet())
```

- A) `"Hello from Parent"`
- B) `"Hello from Child"`
- C) `Error`
- D) `"Child"`

## **5. Class Attribute**

**Question**: What is the output of this code?

```python
class Animal:
    species = "Mammal"

print(Animal.species)
```

- A) `"Mammal"`
- B) `Error`
- C) `None`
- D) `"Animal"`

## **6. Static Method**

**Question**: What does the following code print?

```python
class Utility:
    @staticmethod
    def add(a, b):
        return a + b

print(Utility.add(2, 3))
```

- A) `5`
- B) `Error`
- C) `None`
- D) `"add"`

## **7. `__str__` Method**

**Question**: What will the following code output?

```python
class Car:
    def __str__(self):
        return "Car object"

car = Car()
print(car)
```

- A) `"Car object"`
- B) `Car`
- C) `None`
- D) `Error`

## **8. Superclass Call**

**Question**: What will the output of this code be?

```python
class Parent:
    def __init__(self):
        print("Parent constructor")

class Child(Parent):
    def __init__(self):
        super().__init__()
        print("Child constructor")

child = Child()
```

- A) `"Parent constructor\nChild constructor"`
- B) `"Child constructor"`
- C) `"Parent constructor"`
- D) `Error`

## **9. Private Attributes**

**Question**: What happens when this code is executed?

```python
class Dog:
    def __init__(self):
        self.__name = "Buddy"

dog = Dog()
print(dog.__name)
```

- A) `Buddy`
- B) `Error`
- C) `None`
- D) `"__name"`

## **10. Instance and Class Attributes**

**Question**: What is the output of this code?

```python
class Counter:
    value = 0

    def __init__(self):
        self.value = 1

counter = Counter()
print(counter.value, Counter.value)
```

- A) `0 1`
- B) `1 0`
- C) `Error`
- D) `None`
