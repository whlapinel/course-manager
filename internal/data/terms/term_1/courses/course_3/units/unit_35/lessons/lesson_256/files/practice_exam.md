# PCAP Exam Practice

## **1. Data Types and Variables**

**Question**: What is the output of the following code?

```python
x = 3.5
y = int(x)
z = x + y
print(z)
```

- A) 6.5  
- B) 7  
- C) 6  
- D) 3.5

## **2. Control Flow**

**Question**: What will the following code output?

```python
for i in range(3, -1, -1):
    print(i, end=" ")
```

- A) 0 1 2 3  
- B) 3 2 1 0  
- C) 3 2 1  
- D) SyntaxError

## **3. Functions**

**Question**: What is the output of this code?

```python
def func(a, b=5):
    return a * b

print(func(2))
```

- A) 2  
- B) 10  
- C) TypeError  
- D) None

## **4. Lists**

**Question**: What is the result of the following code?

```python
list1 = [1, 2, 3, 4]
list2 = list1
list1[0] = 99
print(list2)
```

- A) [1, 2, 3, 4]  
- B) [99, 2, 3, 4]  
- C) None  
- D) [1, 2, 3]

## **5. Dictionaries**

**Question**: What will the following code output?

```python
d = {'a': 1, 'b': 2}
d['c'] = 3
print(len(d))
```

- A) 2  
- B) 3  
- C) 4  
- D) KeyError

## **6. Exception Handling**

**Question**: What is the output of the code below?

```python
try:
    x = 5 / 0
except ZeroDivisionError:
    print("Division error")
```

- A) Division error  
- B) ZeroDivisionError  
- C) None  
- D) Program crash

## **7. Classes and Objects**

**Question**: What is true about the following code?

```python
class Example:
    def __init__(self, value):
        self.value = value

obj = Example(10)
print(obj.value)
```

- A) Prints `None`  
- B) Prints `10`  
- C) AttributeError  
- D) SyntaxError

## **8. File Handling**

**Question**: What mode opens a file for writing and overwrites the file if it already exists?

- A) `r`  
- B) `w`  
- C) `a`  
- D) `r+`

## **9. Modules**

**Question**: Which of the following is the correct way to import the `math` module and use its `sqrt` function?

- A) `import math.sqrt`  
- B) `import math` and `math.sqrt(4)`  
- C) `from math import sqrt` and `sqrt(4)`  
- D) Both B and C

## **10. Logical Operators**

**Question**: What will the following code print?

```python
x = True
y = False
print(x and not y)
```

- A) True  
- B) False  
- C) None  
- D) SyntaxError
