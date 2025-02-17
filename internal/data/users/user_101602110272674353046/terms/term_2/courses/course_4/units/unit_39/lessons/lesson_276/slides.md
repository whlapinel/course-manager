---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Warmup

- Make a new directory `lesson_3` in the `unit_2` directory and open it in VS Code
- Create a new file `lesson_3.py` and prepare to run example code from the slides.

```text
/python
  /unit_2
    /lesson_3
      /lesson_3.py
```

# Understanding `if` Statements: Finer Points, Common Mistakes, and Tips

# The Finer Points of `if` Statements

## 1. **Boolean Expressions Matter**
- `if` statements work with any expression that evaluates to `True` or `False`.
- Any non-zero number, non-empty string, list, or dictionary evaluates as `True`.
- `None`, `0`, `False`, and empty collections evaluate as `False`.

# 2. **Using `elif` Instead of Multiple `if` Statements**
- Instead of writing multiple `if` statements, use `elif` to avoid unnecessary checks.

```python
x = 10
if x < 0:
    print("Negative")
elif x == 0:
    print("Zero")
else:
    print("Positive")
```

# Common Mistakes

## 1. **Assigning Instead of Comparing**
- `=` is for assignment, `==` is for comparison.
- Mistake:
```python
if x = 5:  # SyntaxError!
```
- Correct:
```python
if x == 5:
```

# 2. **Not Using `else` When Necessary**
- Forgetting `else` can lead to logic errors where no action is taken when conditions aren't met.

# 3. **Chaining Comparisons Incorrectly**
- Incorrect:
```python
if x > 0 and x < 10:
```
- Better:
```python
if 0 < x < 10:
```

# 4. **Using `or` Incorrectly**
- Mistake:
```python
if x == 1 or 2:  # Always True!
```
- Correct:
```python
if x == 1 or x == 2:
```

# Tips and Tricks

## 1. **Use `in` for Multiple Comparisons**
Instead of:
```python
if x == 1 or x == 2 or x == 3:
```
Use:
```python
if x in [1, 2, 3]:
```

# 2. **Use Ternary `if` for Simple Conditions**
Instead of:
```python
if x > 0:
    sign = "positive"
else:
    sign = "negative"
```
Use:
```python
sign = "positive" if x > 0 else "negative"
```

# 3. **Leverage Boolean Short-Circuiting**
- `and` stops evaluating if the first condition is `False`.
- `or` stops evaluating if the first condition is `True`.

Example:
```python
x = None
y = x or "default"  # y will be "default" if x is None
```

# 4. **Use `not` Sparingly for Readability**
Instead of:
```python
if not (x > 10):
```
Use:
```python
if x <= 10:
```

# Summary
- Use `elif` instead of multiple `if` statements.
- Avoid common mistakes like `=` vs. `==`.
- Use `in` for multiple comparisons.
- Write clear and concise conditions.
- Leverage short-circuiting and ternary expressions wisely.

# Assignment 2.3

- [Enroll in Python Essentials 1 Course](https://pythoninstitute.org/python-essentials-1)
  - Use your personal email account because they have to send you an email to verify
- Solve and submit [Python file](./files/assignment_1_2_3.py)
- Complete CodeHS 2.15 and 2.16
