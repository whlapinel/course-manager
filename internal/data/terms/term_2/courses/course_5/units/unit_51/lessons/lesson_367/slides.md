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

# **Warmup**

- Go to [VS Code Keyboard Shortcuts](https://code.visualstudio.com/shortcuts/keyboard-shortcuts-windows.pdf)
  - Identify 3 cool keyboard shortcuts and try them out
  - Be prepared to share your favorite!

# Lesson 2.3: Recursion: Going Deeper

## Review of Basic Recursion

- A function that calls itself
- Needs a **base case** to stop
- Uses a **recursive case** to reduce the problem
- Examples: Counting down, factorial, reversing a string

# Why is Recursion Hard?

- Hard to visualize function calls
- Stack memory can get deep
- Need to trust the function to work correctly
- Base cases must be well-defined to prevent infinite recursion

# Call Stack: How Recursion Works Internally

When a function calls itself, Python:
1. Saves the current function call in a stack (memory)
2. Moves to the next recursive call
3. Returns and removes functions from the stack when base case is reached

# Example: Tracing the Call Stack

```python
def countdown(n):
    if n <= 0:
        print("Blast off!")
    else:
        print(n)
        countdown(n - 1)
```

# What Happens?

```text
countdown(3) → prints 3 → calls countdown(2)
countdown(2) → prints 2 → calls countdown(1)
countdown(1) → prints 1 → calls countdown(0)
countdown(0) → prints "Blast off!"
```

Each call is stored until it finishes.

# Visualizing Recursion: The Stack

Example for `countdown(3)`:
```text
countdown(3)
  countdown(2)
    countdown(1)
      countdown(0)  # Base case, starts returning
```

Each function waits for the next one to complete.

# Example: Recursively Searching a List

```python
def search(lst, target):
    if not lst:
        return False  # Base case
    if lst[0] == target:
        return True
    return search(lst[1:], target)  # Recursive case
```

**How It Works:**
1. Check if the list is empty (base case)
2. If the first element is the target, return `True`
3. Otherwise, call `search()` on the rest of the list

# Tail Recursion vs. Regular Recursion

- **Regular Recursion**: The function does work after the recursive call (e.g., factorial)
- **Tail Recursion**: The function makes the recursive call as the last step
- Some languages optimize tail recursion to prevent deep stacks (Python does not)

# Example: Tail Recursive Version of Countdown

```python
def countdown(n):
    if n <= 0:
        print("Blast off!")
        return
    print(n)
    countdown(n - 1)  # Last action (tail recursion)
```

Tail recursion is more memory-efficient in some languages, but not in Python.

# Common Mistakes in Recursion

1. **Forgetting the base case** → Infinite recursion
2. **Modifying input incorrectly** → Skipping elements
3. **Returning incorrect values** → Breaking recursion logic

# Practice: Recursive Sum

```python
def recursive_sum(n):
    if n == 0:
        return 0  # Base case
    return n + recursive_sum(n - 1)  # Recursive case
```

**Try It:** What happens when `recursive_sum(5)` is called?

# Recap

- Recursion builds on itself through repeated function calls
- The call stack holds function states until returning
- Base cases stop infinite recursion
- Recursive search and sum are simple applications
