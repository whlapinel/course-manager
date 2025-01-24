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
    t.goto(0, t.ycor() - size)
    t.pendown()

turtle.done()
