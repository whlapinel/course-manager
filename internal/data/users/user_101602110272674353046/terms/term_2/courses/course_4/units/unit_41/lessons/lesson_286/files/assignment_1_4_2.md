# Python Exception Handling - Practice Problems

## Practice Problems

### 1. Validate Age
- Write a function `validate_age(age: int) -> int` that raises a `ValueError` if the age is negative.

```python
def validate_age(age: int) -> int:
    pass
```

---
### 2. Safe Division
- Write a function `safe_divide(a: float, b: float) -> float` that catches `ZeroDivisionError` and returns `None` instead.

```python
def safe_divide(a: float, b: float) -> float:
    pass
```

---
### 3. Read File Safely
- Write a function `read_file(filename: str) -> str` that catches `FileNotFoundError` and returns an error message.

```python
def read_file(filename: str) -> str:
    pass
```

---
### 4. Retrieve Dictionary Key
- Write a function `get_value(data: dict, key: str) -> any` that raises a `KeyError` if the key is missing.

```python
def get_value(data: dict, key: str) -> any:
    pass
```

---
### 5. Custom Exception for Temperature
- Define a custom exception `TooColdError`.
- Write a function `check_temperature(temp: float)` that raises `TooColdError` if the temperature is below -30 degrees.

```python
class TooColdError(Exception):
    pass

def check_temperature(temp: float):
    pass
```
