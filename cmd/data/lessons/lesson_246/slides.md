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

```python
my_list = [
    [1,2,3,4,5],
    [1,2,3,4,5],
    [1,2,3,4,5],
]
```

Write a function `col_sum` that will take any 2D list of integers and provide a sum for each column.

Example:

```python
assert col_sum(my_list) == [3,6,9,12,15]
```

# **Introduction to Web Development**

# Welcome to Web Development 101 🌐

## What You Will Learn

- How the web works
- Basic structure of a web page (HTML)
- Serving web pages with Flask

# **How the Web Works**

## How the Web Works

1. **Browser**: Sends a request for a webpage.
2. **Server**: Responds with the requested page (HTML).
3. **HTML**: The browser renders the content for the user.

# Key Terms

- **Client**: Your browser.
- **Server**: The computer hosting the website.
- **Request**: Asking for a webpage.
- **Response**: Sending back the webpage.

# **What is Flask?**

## What is Flask?

- A lightweight framework for building web applications in Python.
- Helps you:
  - Create web servers.
  - Send HTML pages to users.
- Great for beginners and small projects!

Run in your terminal:

```bash
pip install Flask
```

# Example

```python
from flask import Flask

app = Flask(__name__)

@app.route("/")
def home():
    return "Hello, World!"

app.run()
```

# **Structure of a Web Page**

1. **HTML**: The skeleton of a webpage.
2. **CSS**: Adds style (colors, fonts).
3. **JavaScript**: Adds interactivity.

# Example HTML Page

```html
<!DOCTYPE html>
<html>
<head>
    <title>My First Page</title>
</head>
<body>
    <h1>Hello, Web!</h1>
    <p>This is my first webpage.</p>
</body>
</html>
```

# **Flask + HTML**

## Example Code

```python
from flask import Flask, send_from_directory

app = Flask(__name__)

@app.route("/")
def home():
    return send_from_directory(directory="static", path="index.html")

app.run()
```

# What Happens

1. Flask serves the `index.html` file.
2. Your browser renders the file as a webpage.

# **Building Your Project**

## Directory Structure

project/
├── app.py          # Flask server
├── static/         # Static files (HTML, CSS)
    ├── index.html  # Home page
    ├── about.html  # About page

1. Write your `index.html` page.
2. Add a route in `app.py`.
3. Run your Flask server and visit the page in your browser.

# **Your Turn: First Web Page**

1. Create a file: `index.html`:

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>My Home Page</title>
</head>

<body>
    <h1>Welcome to My Home Page</h1>
    <p>This is a simple HTML page.</p>
    <a href="about">Learn more about me</a>
</body>
</html>
```

1. Run Flask and visit `http://127.0.0.1:5000/`.

🎉 Congratulations, you're on the web!

# Assignment 8.1

- Personalize your pages
- Add a header and footer to your home page (`index.html`)
- Create another page called `about.html` that includes a link back to your home page.
- Add a route in `app.py`.
- Submit your files as a zip file
  - server.py
  - index.html
  - about.html

🚀 Keep building and experimenting!
