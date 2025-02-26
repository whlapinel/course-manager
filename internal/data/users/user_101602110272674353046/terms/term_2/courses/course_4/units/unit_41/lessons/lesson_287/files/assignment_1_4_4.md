# Assignment 4.4: More on Functions in Python - Practice Problems

## Instructions
- Complete each problem by writing a function.
- Ensure your function includes **exception handling** where appropriate.
- Use different types of parameters (`*args`, `**kwargs`, default values).
- Run your function with example inputs to test it.
- Use the hints if you get stuck.

## Problem 1: Safe Division
Write a function `safe_divide` that takes two numbers and returns the result of dividing the first by the second. If division by zero is attempted, raise a `ZeroDivisionError` with a custom message.

**Hint:** Use `try-except` to catch the error.

```python
def safe_divide(a, b):
    pass

print(safe_divide(10, 2))
print(safe_divide(5, 0))  # Should raise an error
```

## Problem 2: Validate Age Input
Write a function `check_age` that takes an age as input. If the age is not a positive integer, raise a `ValueError`.

**Hint:** Use `isinstance()` and an `if` condition.

```python
def check_age(age):
    pass

print(check_age(25))
print(check_age(-5))  # Should raise an error
```

## Problem 3: Process Multiple Scores
Write a function `calculate_average` that takes any number of scores (`*args`) and returns their average. If no scores are provided, raise a `ValueError`.

**Hint:** Use `len(args)` to check if arguments exist.

```python
def calculate_average(*scores):
    pass

print(calculate_average(90, 80, 70))
print(calculate_average())  # Should raise an error
```

## Problem 4: Flexible Greeting
Write a function `greet_person` that accepts a name and an optional greeting message (defaulting to "Hello"). If the name is not a string, raise a `TypeError`.

**Hint:** Use `kwargs` to allow additional custom messages.

```python
def greet_person(name, greeting="Hello", **kwargs):
    pass

print(greet_person("Alice"))
print(greet_person(42))  # Should raise an error
```

## Problem 5: Get Dictionary Value Safely
Write a function `get_value` that takes a dictionary and a key. If the key is not found, return "Key not found" instead of raising an error.

**Hint:** Use `.get()` method of dictionaries.

```python
def get_value(data, key):
    pass

sample_dict = {"name": "Ash", "pokemon": "Pikachu"}
print(get_value(sample_dict, "name"))
print(get_value(sample_dict, "age"))  # Should return "Key not found"
```

## Problem 6: Validate a List of Numbers
Write a function `validate_numbers` that takes a list of numbers. If any element is not a number, raise a `TypeError`.

**Hint:** Use `all()` to check types.

```python
def validate_numbers(numbers):
    pass

print(validate_numbers([1, 2, 3.5]))
print(validate_numbers([1, "two", 3]))  # Should raise an error
```

## Problem 7: Reverse a String Safely
Write a function `safe_reverse` that takes a string and returns it reversed. If the input is not a string, raise a `TypeError`.

**Hint:** Use slicing `[::-1]`.

```python
def safe_reverse(text):
    pass

print(safe_reverse("hello"))
print(safe_reverse(123))  # Should raise an error
```

## Submission
- Write your solutions in a `.py` file.
- Ensure each function runs correctly with test inputs.
- Submit your `.py` file as per instructions.
