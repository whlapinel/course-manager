---
layout: none
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# **Category: Binary Numbers & Bitwise Operations**

## **Question 1:**
What is the result of `0b1101 & 0b1011` in binary?

# **Answer:**
`0b1001` (which is 9 in decimal)

# **Question 2:**
What will `~5` return in Python, and why?

# **Answer:**
It returns `-6` because `~x` is equivalent to `-(x+1)`.

# **Category: Creating and Importing Modules**

## **Question 3:**
You wrote a Python script `math_tools.py` with a function `square(n)`. How would you import and use it in another script?

# **Answer:**
```python
import math_tools
print(math_tools.square(4))
```
Or:
```python
from math_tools import square
print(square(4))
```

# **Question 4:**
What is the difference between `import mymodule` and `from mymodule import some_function`?

# **Answer:**
`import mymodule` requires `mymodule.some_function()` to call the function. `from mymodule import some_function` allows direct usage of `some_function()` without the module prefix.

# **Category: Using the `datetime` Module**

## **Question 5:**
Write a Python snippet that prints the current year without importing the entire `datetime` module.

# **Answer:**
```python
from datetime import datetime
print(datetime.now().year)
```

# **Category: Recursion**

## **Question 6:**
A recursive function must have what key feature to prevent infinite recursion?

# **Answer:**
A **base case**, which ensures the recursion stops at a certain condition.

# **Question 7:**
What will happen if a recursive function has no base case?

# **Answer:**
It will lead to infinite recursion and eventually cause a `RecursionError`.

# **Category: Nested Loops**

## **Question 8:**
What will be printed by the following code?

```python
for i in range(2):
    for j in range(2, 4):
        print(i, j)
```

# **Answer:**
```text
0 2
0 3
1 2
1 3
```

# **Category: Complex Conditionals**

## **Question 9:**
Rewrite the following condition using **De Morgan’s Laws**:

```python
if not (x > 5 or y < 3):
    print("Condition met")
```

# **Answer:**
```python
if x <= 5 and y >= 3:
    print("Condition met")
```

# **Category: If-Else, While, and For Loops**

## **Question 10:**
What will be printed by the following code?

```python
x = 10
while x > 0:
    if x % 3 == 0:
        print("Divisible by 3:", x)
    x -= 2
```

# **Answer:**
```text
Divisible by 3: 9
Divisible by 3: 3
```

# **Question 11:**
You need to loop through a dictionary and print both the **keys** and **values**. What built-in method allows you to do this efficiently?

# **Answer:**
`dict.items()`
```python
for key, value in my_dict.items():
    print(key, value)
```

# **Category: Functions**

## **Question 12:**
What is the output of the following function call?

```python
def mystery(a=[]):
    a.append(1)
    return a

print(mystery())
print(mystery())
```

# **Answer:**
```text
[1]
[1, 1]
```
Because the default list persists between function calls.

# **Category: Basic Data Structures**

## **Question 13:**
How would you swap the values of two variables `a` and `b` **without** using a third variable?

# **Answer:**
```python
a, b = b, a
```

# **Question 14:**
What is the difference between `list.append(x)` and `list.extend([x])`?

# **Answer:**
`append(x)` adds a single element, while `extend([x])` iterates over the argument and adds each item.
```python
lst = [1, 2]
lst.append([3])  # [1, 2, [3]]
lst.extend([3])  # [1, 2, 3]
```

# **Category: Assertions & Debugging**

## **Question 15:**
What will happen if the following code runs and `x` is `-5`?

```python
assert x >= 0, "x must be non-negative"
```

# **Answer:**
An `AssertionError` will be raised with the message: `x must be non-negative`.

# **Question 16:**
Write an assertion statement that ensures a given list `nums` has at least one element before accessing its first element.

# **Answer:**
```python
assert len(nums) > 0, "List must not be empty"
print(nums[0])
```
