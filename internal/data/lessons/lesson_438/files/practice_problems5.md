# Practice Problems 9.5

## Section 1: Modules

## **1. Importing Modules**

### Which of the following statements correctly imports the `sqrt` function from the `math` module?

A. `import sqrt`  
B. `from math import sqrt`  
C. `import math.sqrt`  
D. `math sqrt import`  

## **2. Nested Modules**  

### How can you access a function `func` inside a nested module `mod.submod`?  

A. `import mod.submod` → `mod.func()`  
B. `from mod.submod import func` → `func()`  
C. `import func from mod.submod`  
D. `import mod` → `mod.submod.func()`  

## **3. The `dir()` Function**  

### What does the `dir()` function return when applied to a module?  

A. Documentation of the module  
B. List of all functions and attributes in the module  
C. The path of the module  
D. The size of the module  

## **4. The `sys.path` Variable**  

### Which of the following statements about `sys.path` is correct?  

A. It lists the directories where Python searches for modules.  
B. It lists the imported modules only.  
C. It deletes unused modules automatically.  
D. It stores the system’s environment variables.  

## **5. Using the `math` Module**  

### What does the following code return?  

```python
import math  
print(math.ceil(4.3))  
```  

A. `4`  
B. `5`  
C. `4.3`  
D. `Error`  

## **6. The `random` Module**  

### Which function from the `random` module is used to select a random item from a list?

A. `random()`  
B. `choice()`  
C. `seed()`  
D. `sample()`  

## **7. Discovering Platform Properties**  

### Which function retrieves the processor name of the host platform?  

A. `platform.system()`  
B. `platform.processor()`  
C. `platform.version()`  
D. `platform.machine()`  

## **8. User-Defined Modules**  

### What is the role of the `__name__` variable in a module?  

A. It specifies the module’s version.  
B. It checks if the module is imported or run directly.  
C. It defines the module’s search path.  
D. It lists all variables in the module.  

## **9. The `__pycache__` Directory**  

### What does the `__pycache__` directory store?

A. Plain text copies of the module  
B. Encrypted module files  
C. Compiled bytecode files of the module  
D. Backup copies of the module  

## **10. Public vs. Private Variables**  

### How is a **private variable** defined in a Python module?  

A. It starts with a single underscore `_`.  
B. It starts with two underscores `__`.  
C. It is declared with the `private` keyword.  
D. It is defined using the `@private` decorator.  

## Section 2: Exceptions

## **11. Python-Defined Exceptions**  

### What is the correct way to catch multiple exceptions?  

A. `except (e1, e2):`  
B. `catch e1, e2:`  
C. `try: e1, e2:`  
D. `handle e1 or e2:`  

## **12. Raising Exceptions**  

### What does the `raise` keyword do in Python?  

A. Terminates the program unconditionally  
B. Allows defining a custom exception  
C. Re-raises the current exception  
D. Ignores all exceptions  

## **13. The `assert` Statement**  

### What happens when an `assert` statement fails?  

A. The program crashes silently  
B. An `AssertionError` is raised  
C. A warning is logged  
D. The program continues execution  

## **14. Exception Hierarchy**  

### Which of the following is the parent class of all exceptions in Python?  

A. `BaseException`  
B. `Exception`  
C. `Error`  
D. `RuntimeError`  

## **15. Using `except` Blocks**  

### Which syntax allows you to catch an exception and access its details?  

A. `except Exception as e:`  
B. `catch Exception -> e:`  
C. `try Exception as e:`  
D. `handle Exception {e}:`  

## **16. Else Clause in `try` Blocks**  

### When does the `else` clause in a `try` block execute?  

A. If an exception occurs  
B. If no exception occurs  
C. Before the `finally` block  
D. Always  

## **17. Custom Exceptions**  

### How do you define a custom exception in Python?  

A. Inherit from `Exception` class  
B. Use `raise new Exception()`  
C. Define it using `try`  
D. Use `catch Exception`  

## **18. Multiple `except` Blocks**  

### What is the correct order of catching exceptions?  

A. From most specific to most general  
B. From most general to most specific  
C. Random order  
D. Specific exceptions cannot be caught  

## **19. The `finally` Block**  

### What is the purpose of the `finally` block?  

A. It executes only if no exception occurs  
B. It executes regardless of whether an exception occurs or not  
C. It catches unhandled exceptions  
D. It suppresses exceptions  

## **20. Self-Defined Exceptions**  

### What is required when creating a self-defined exception?  

A. Override the `__str__` method  
B. Inherit from `BaseException` or `Exception`  
C. Define a `try` block  
D. Use the `raise` keyword only  

## **21. Encoding Standards**  

### What is the default character encoding used in Python 3?  

A. ASCII  
B. UTF-8  
C. Unicode  
D. ISO-8859-1  

## **22. `ord()` and `chr()`**  

### What does the function `chr(65)` return?  

A. `'A'`  
B. `65`  
C. `'65'`  
D. `'\x41'`  

## **23. String Indexing**  

### What is the output of the following code?  

```python
s = "hello"
print(s[1])
```  

A. `'e'`  
B. `'h'`  
C. `'l'`  
D. `IndexError`  

## **24. String Slicing**  

### What does `s[2:5]` return if `s = "python"`?  

A. `'pyt'`  
B. `'yth'`  
C. `'tho'`  
D. `'hon'`  

## **25. String Immutability**  

### What happens if you attempt to modify a string in Python, like this:

```python
s = "immutable"
s[0] = "M"
```  

A. The string becomes `"Mutable"`.  
B. A `TypeError` is raised.  
C. The string remains unchanged.  
D. It creates a new string.  

## **26. Concatenating Strings**  

### What is the result of the following code?  

```python
s1 = "Hello"
s2 = "World"
print(s1 + " " + s2)
```  

A. `'HelloWorld'`  
B. `'Hello World'`  
C. `HelloWorld`  
D. `Error`  

## **27. Membership Operators**  

### What does the following code return?  

```python
s = "Python"
print("y" in s)
```  

A. `True`  
B. `False`  
C. `'y'`  
D. `None`  

## **28. String Methods**  

### Which string method checks if all characters in a string are alphabetic?  

A. `.isdigit()`  
B. `.isalpha()`  
C. `.isalnum()`  
D. `.isupper()`  

## **29. String Splitting**  

### What does the following code return?

```python
s = "one,two,three"
print(s.split(","))
```  

A. `['one,two,three']`  
B. `['one', 'two', 'three']`  
C. `['one two three']`  
D. `Error`  

## **30. String Comparison**  

### What will the following code output?  

```python
s1 = "apple"
s2 = "banana"
print(s1 < s2)
```  

A. `True`  
B. `False`  
C. `None`  
D. `Error`  

## Section 4: Classes and Objects

## **31. Class and Object Basics**  

### What is the correct definition of a **class** in Python?  

A. A blueprint for creating objects  
B. An instance of an object  
C. A method that modifies data  
D. A data structure that holds functions  

## **32. Instance vs. Class Variables**  

### Which statement correctly distinguishes between **instance** and **class** variables?

A. Instance variables are shared across all objects.  
B. Class variables are unique to each object.  
C. Instance variables are specific to each object.  
D. Class variables can only be modified using `self`.  

## **33. The `self` Parameter**  

### What is the purpose of the `self` parameter in a class method?  

A. It initializes the class.  
B. It refers to the instance of the class.  
C. It refers to a global variable.  
D. It defines the class properties.  

## **34. Class Properties**  

### Which of the following properties retrieves the name of a class?  

A. `__module__`  
B. `__name__`  
C. `__bases__`  
D. `__dict__`  

## **35. Constructor Method**  

### What is the purpose of the `__init__()` method in a class?  

A. It deletes objects.  
B. It constructs objects when a class is defined.  
C. It initializes an object when it is created.  
D. It retrieves class properties.  

## **36. Private Components**  

### How do you define a **private attribute** in a Python class?  

A. Prefix it with a single underscore (`_`).  
B. Prefix it with two underscores (`__`).  
C. Use the `@private` decorator.  
D. Use the `private` keyword.  

## **37. Inheritance**  

### What does single inheritance allow you to do in Python?  

A. Use multiple base classes.  
B. Override the base class’s methods.  
C. Prevent method overriding.  
D. Combine two classes into one.  

## **38. The `isinstance()` Function**  

### What does the `isinstance()` function check?  

A. If an object is a subclass of another class  
B. If an object is an instance of a class or subclass  
C. If a class has private components  
D. If a method has the `self` parameter  

## **39. Polymorphism**  

### What does **polymorphism** mean in object-oriented programming?  

A. Methods with the same name but different functionality  
B. Combining two classes into one  
C. Accessing private attributes directly  
D. Creating multiple objects of the same class  

## **40. Overriding Methods**  

### What happens when you **override** a method in a subclass?  

A. The original method in the base class is deleted.  
B. The subclass method replaces the base class method.  
C. Both methods execute in sequence.  
D. The base class method is made private.  

## **41. List Comprehensions**  

### What does the following list comprehension produce?  

```python
nums = [1, 2, 3, 4]
squares = [x**2 for x in nums if x % 2 == 0]
print(squares)
```  

A. `[1, 4, 9, 16]`  
B. `[4, 16]`  
C. `[2, 4]`  
D. `[1, 2, 3, 4]`  

## **42. Nested List Comprehensions**  

### Which statement correctly creates a list of pairs `(x, y)` where `x` is from `[1, 2]` and `y` is from `[3, 4]`?

A. `[(x, y) for x in [1, 2], y in [3, 4]]`  
B. `[(x, y) for x in [1, 2] for y in [3, 4]]`  
C. `[(x, y) if x in [1, 2] and y in [3, 4]]`  
D. `[x, y for x in [1, 2] for y in [3, 4]]`  

## **43. Lambda Functions**  

### What does the following lambda function do?  

```python
add = lambda x, y: x + y
print(add(2, 3))
```  

A. Prints `2`  
B. Prints `3`  
C. Prints `5`  
D. Raises an error  

## **44. Using `map()`**  

### What does the `map()` function do?  

A. Iterates through two lists simultaneously  
B. Applies a function to each item in an iterable  
C. Filters items from a list based on a condition  
D. Combines multiple iterables  

## **45. Using `filter()`**  

### What does the following code return?  

```python
nums = [1, 2, 3, 4]
even = list(filter(lambda x: x % 2 == 0, nums))
print(even)
```  

A. `[1, 2, 3, 4]`  
B. `[2, 4]`  
C. `[1, 3]`  
D. `[2]`  

## **46. Closures**  

### What is a closure in Python?  

A. A nested function that has access to variables from its enclosing scope  
B. A function that closes open files  
C. A private class method  
D. A function that takes another function as input  

## **47. Input/Output Modes**  

### What does the mode `'rb'` mean when opening a file?  

A. Read binary  
B. Read-only  
C. Write binary  
D. Append binary  

## **48. Writing to a File**  

### What does the following code do?  

```python
with open("example.txt", "w") as f:
    f.write("Hello World")
```  

A. Appends `"Hello World"` to the file.  
B. Overwrites the file with `"Hello World"`.  
C. Reads the file and prints `"Hello World"`.  
D. Raises an error.  

## **49. Predefined Streams**  

### Which predefined stream is used to output data to the console?  

A. `sys.stdin`  
B. `sys.stdout`  
C. `sys.stderr`  
D. `sys.console`  

## **50. Bytearray Buffer**  

### What is the purpose of using `bytearray` in file operations?  

A. To create a binary buffer for I/O operations  
B. To compress file data  
C. To filter text files  
D. To convert binary files to text  
