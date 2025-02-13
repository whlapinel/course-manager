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

# Warmup

- Go to [VS Code Keyboard Shortcuts](https://code.visualstudio.com/shortcuts/keyboard-shortcuts-windows.pdf)
  - Identify 3 cool keyboard shortcuts and try them out
  - Be prepared to share your favorite!

# 🔄 While Loops - Lesson 3.2

## Building on What We Learned Yesterday

# 🔁 Recap: What is a While Loop?
💡 A `while` loop runs as long as a **condition** is `True`.

```python
x = 1
while x <= 5:
    print(x)
    x += 1
```

✅ Remember: Be careful of **infinite loops!**

# 🚀 Challenge #1: Debug This Code
What's wrong with this loop? How can we fix it?

```python
x = 10
while x > 0:
    print("Counting down:", x)
```

# 🛠 Adding a Break Statement
💡 The `break` statement stops the loop immediately.

```python
while True:
    name = input("Enter 'quit' to stop: ")
    if name == "quit":
        break
    print("Hello,", name)
```

# 🧠 Challenge #2: Fix the Loop
The following loop is **meant to stop when the user enters a negative number**. Fix it!

```python
num = int(input("Enter a number: "))
while num >= 0:
    print("You entered:", num)
    num = int(input("Enter another number: "))
print("Done!")
```

# 🔄 Using `continue`
💡 The `continue` statement **skips the rest of the loop body** and jumps to the next iteration.

```python
x = 0
while x < 5:
    x += 1
    if x == 3:
        continue
    print(x)
```

❓ What will be the output?

# 🏆 Challenge #3: Print Only Even Numbers
Modify this code so that it **only prints even numbers** up to 10.

```python
x = 0
while x < 10:
    x += 1
    print(x)
```

# 🔍 Nested While Loops
💡 A `while` loop can be placed **inside another while loop**.

```python
i = 1
while i <= 3:
    j = 1
    while j <= 2:
        print(f"i = {i}, j = {j}")
        j += 1
    i += 1
```

# ❓ What will be the output?

## 🎯 Challenge #4: Create a Number Pyramid
Write a while loop that prints this pattern:

```text
1
12
123
1234
12345
```

# 🏁 Summary & Next Steps
## ✅ Today we covered:
- **Break** and **Continue**
- Debugging `while` loops
- Nested `while` loops
- More practice with challenges

# Assignment 3.2

- Python file [Hints](files/hints_assignment_1_3_2.html)
- CodeHS 3.1 and 3.2
