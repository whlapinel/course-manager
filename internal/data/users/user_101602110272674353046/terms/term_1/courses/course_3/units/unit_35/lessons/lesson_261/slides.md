---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# **Holiday-Themed Drawing Assignment**

## 🎄 Objective:

Create a **holiday-themed drawing** using Python.  
Examples:  

- A Christmas tree  
- A snowman  
- Festive lights  

# **Assignment Details**

## ✨ Guidelines:

- Use Python's **Turtle** or an alternative drawing tool.
- Be creative! Combine shapes and colors to bring your holiday scene to life.
- Submit your Python script and a screenshot of your drawing.

# **Getting Started with Python Turtle**

## What is Turtle?

- Turtle is a built-in module in Python.
- It allows you to create graphics using simple commands.

# How to Use:
1. Import the turtle module:

   ```python
   import turtle
   ```

2. Create a turtle:

   ```python
   t = turtle.Turtle()
   ```

3. Start drawing!

# **Basic Turtle Commands**

## Drawing a Line
```python
t.forward(100)  # Move forward by 100 pixels
t.left(90)      # Turn left by 90 degrees
```

## Changing Colors
```python
t.color("blue")  # Set the pen color
```

## Adding Shapes
```python
t.circle(50)     # Draw a circle with radius 50
```

# **Drawing Example: Christmas Tree**

```python
import turtle

t = turtle.Turtle()
t.color("green")

# Draw the tree
for size in [60, 80, 100]:
    t.begin_fill()
    t.forward(size)
    t.left(120)
    t.forward(size)
    t.left(120)
    t.forward(size)
    t.left(120)
    t.end_fill()
    t.penup()
    t.goto(0, t.ycor() - 20)
    t.pendown()

turtle.done()
```

# **Alternative Drawing Tools**

# 1. **Matplotlib**
Great for simple geometric designs.
```python
import matplotlib.pyplot as plt

plt.plot([0, 1], [0, 1])  # Draw a line
plt.show()
```

# 2. **Pygame**  
More advanced but good for dynamic visuals.
```python
import pygame

pygame.init()
screen = pygame.display.set_mode((400, 300))
pygame.draw.circle(screen, (255, 0, 0), (200, 150), 50)
pygame.display.flip()
pygame.time.wait(2000)
```

# **Submission**

## What to Submit:
- Python code.
- Screenshot of your drawing.

## Grading Criteria:
1. Creativity (30%)
2. Use of Python (40%)
3. Presentation (30%)
