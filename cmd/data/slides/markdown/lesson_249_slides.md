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

# Agenda

- No exam this week
- Catch-up and/or work ahead
  - Please check your grades in Canvas for anything missing
  - Lesson 8.4 is posted for working ahead

# Looking ahead

- PCAP Exam?
- Unit 8 Project
- Final Project

# Flask Templates and Components

# Introduction to Flask Templates

- Flask uses **Jinja2** templating engine.
- Allows you to embed Python code in HTML.
- Helps create dynamic web pages with ease.

# Setting Up Templates

1. Create a `templates/` folder in your Flask project directory.
2. Add HTML files in the `templates/` folder.

Example folder structure:

```plaintext
project/
  app.py
  templates/
    index.html
    about.html
```

# Rendering a Template

- Use the `render_template` function to display templates.

**Example:**

```python
from flask import Flask, render_template

app = Flask(__name__)

@app.route('/')
def home():
    return render_template('index.html', title="Home")

if __name__ == '__main__':
    app.run(debug=True)
```

---

# Using Template Variables

- Pass variables to templates via `render_template`.

**app.py:**

```python
@app.route('/about')
def about():
    return render_template('about.html', title="About", name="John")
```

# **about.html:**

```html
<!DOCTYPE html>
<html>
<head>
    <title>{{ title }}</title>
</head>
<body>
    <h1>Welcome, {{ name }}!</h1>
</body>
</html>
```

# Template Components (Reusable HTML)

- Use `{% include %}` to insert reusable components.

# Step 1: Create the Component

**templates/header.html:**

```html
<header>
    <h1>My Website</h1>
    <nav>
        <a href="/">Home</a> |
        <a href="/about">About</a>
    </nav>
</header>
```

# Step 2: Include the Component in Pages

# **templates/index.html:**

```html
<!DOCTYPE html>
<html>
<head>
    <title>{{ title }}</title>
</head>
<body>
    {% include 'header.html' %}
    <h2>Welcome to the Homepage!</h2>
</body>
</html>
```

# **templates/about.html:**

```html
<!DOCTYPE html>
<html>
<head>
    <title>{{ title }}</title>
</head>
<body>
    {% include 'header.html' %}
    <h2>About Us</h2>
</body>
</html>
```

# Benefits of Components

- **Reusability**: Avoid duplicating HTML code.
- **Consistency**: Ensure uniform styling across pages.
- **Maintainability**: Update the component once, reflect changes everywhere.

# Assignment 8.4 (Due January 6th)

- Practice creating and using templates and reusable components in Flask.

# **Set Up the Project**:

- Create a Flask project with the following structure:

     ```bash
     project/
       app.py
       templates/
         base.html
         header.html
         home.html
         about.html
     ```

# **Create a Header Component**

- Create a file `header.html` in the `templates/` folder.
- Add a navigation bar with links to "Home" and "About".

   Example content for `header.html`:

   ```html
   <header>
       <h1>My Flask App</h1>
       <nav>
           <a href="/">Home</a> |
           <a href="/about">About</a>
       </nav>
   </header>
   ```

# **Create a Base Template**

- Create a file `base.html` in the `templates/` folder.
- Add a basic HTML structure that includes the `header.html` file and a `{% block content %}` placeholder.

   Example content for `base.html`:

   ```html
   <!DOCTYPE html>
   <html>
   <head>
       <title>{{ title }}</title>
   </head>
   <body>
       {% include 'header.html' %}
       <div>
           {% block content %}{% endblock %}
       </div>
   </body>
   </html>
   ```

# **Create Content Pages**

- Create `home.html` and `about.html` files, extending the `base.html` template.
- Use the `{% block content %}` to insert page-specific content.

   Example `home.html`:

   ```html
   {% extends 'base.html' %}
   {% block content %}
       <h2>Welcome to the Homepage!</h2>
       <p>This is the main page of the site.</p>
   {% endblock %}
   ```

# Create `about.html`

   Example `about.html`:

   ```html
   {% extends 'base.html' %}
   {% block content %}
       <h2>About Us</h2>
       <p>Learn more about our Flask project here.</p>
   {% endblock %}
   ```

# **Set Up Routes**

- Define routes in `app.py` for the homepage and about page.
- Pass a title to each template.

   Example `app.py`:

   ```python
   from flask import Flask, render_template

   app = Flask(__name__)

   @app.route('/')
   def home():
       return render_template('home.html', title="Home")

   @app.route('/about')
   def about():
       return render_template('about.html', title="About")

   if __name__ == '__main__':
       app.run(debug=True)
   ```

# **Submission**

- Submit your completed Flask project.
- Verify that the header component appears on both the homepage and the about page.
- Ensure the navigation links work correctly.

# Next Steps

1. Practice creating templates and components.
2. Explore other Jinja2 features:
   - Conditionals (`{% if %}`)
   - Loops (`{% for %}`)
3. Build a simple multi-page Flask application.