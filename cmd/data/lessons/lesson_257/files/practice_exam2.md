# Practice Problems 9.2

## **Day 2 Practice Problems (Intermediate Python Skills)**

## **1. String Methods**

**Question**: What is the output of the following code?

```python
text = "  Python Crash Course  "
print(text.strip().upper())
```

- A) `" PYTHON CRASH COURSE "`
- B) `"python crash course"`
- C) `"Python Crash Course"`
- D) `"PYTHON CRASH COURSE"`

## **2. String Replacement**

**Question**: What will this code print?

```python
text = "I love Python"
print(text.replace("love", "like"))
```

- A) `"I like Python"`
- B) `"I love Python"`
- C) `"I  Python"`
- D) `None`

## **3. String Formatting**

**Question**: What is the output of this code?

```python
name = "Alice"
age = 25
print(f"My name is {name} and I am {age} years old.")
```

- A) `"My name is {name} and I am {age} years old."`
- B) `"My name is Alice and I am years old."`
- C) `"My name is Alice and I am 25 years old."`
- D) `"My name is Alice and I am {age} years old."`

## **4. Exceptions**

**Question**: What does this code output?

```python
try:
    x = 5 / 0
except ZeroDivisionError:
    print("Cannot divide by zero.")
```

- A) `ZeroDivisionError`
- B) `"Cannot divide by zero."`
- C) `"Error"`
- D) `None`

## **5. Multiple Exception Types**

**Question**: Which exception will be caught in the following code?

```python
try:
    x = int("hello")
except ValueError:
    print("Value error caught.")
except TypeError:
    print("Type error caught.")
```

- A) ValueError
- B) TypeError
- C) Both
- D) None

## **6. Raising Exceptions**

**Question**: What is the output of this code?

```python
def check_positive(num):
    if num < 0:
        raise ValueError("Negative number not allowed.")
check_positive(-5)
```

- A) `Negative number not allowed.`
- B) `ValueError`
- C) Program crash
- D) Both A and B

## **7. Custom Exceptions**

**Question**: What will this code output?

```python
class MyException(Exception):
    pass

try:
    raise MyException("Custom error message.")
except MyException as e:
    print(e)
```

- A) `"Custom error message."`
- B) `MyException`
- C) `Custom Exception Raised`
- D) None

## **8. Nested Try-Except**

**Question**: What is the output of the following?

```python
try:
    try:
        x = int("NaN")
    except ValueError:
        print("Inner exception")
    x = 5 / 0
except ZeroDivisionError:
    print("Outer exception")
```

- A) `"Inner exception\nOuter exception"`
- B) `"Inner exception"`
- C) `"Outer exception"`
- D) None

## **9. String Formatting with `.format()`**

**Question**: What is the result of this code?

```python
print("The result is {0} and {1}".format(10, "success"))
```

- A) `The result is 0 and success`
- B) `The result is 10 and success`
- C) `The result is success and 10`
- D) Error

## **10. Exception Order**

**Question**: What happens when you run this code?

```python
try:
    x = 10 / 0
except Exception:
    print("Caught general exception.")
except ZeroDivisionError:
    print("Caught zero division error.")
```

- A) `"Caught general exception."`
- B) `"Caught zero division error."`
- C) SyntaxError
- D) RuntimeError
