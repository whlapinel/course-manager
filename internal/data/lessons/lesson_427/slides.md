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

# Warmup Lesson 8.4

Glance over lessons 8.1 - 8.3 to refresh your memory on web development

- What is HTML?
- What is CSS?

# Agenda

- Templating using Flask
- Assignment 8.4

# Looking ahead

- PCAP Exam? 🤷🏻‍♂️🤷🏻‍♂️🤷🏻‍♂️
- Unit 8 Project due Friday: details soon
- Final Project due the following Friday

# Web Development Refresher

What is this an example of?

```css

body {
  background-color: lightblue;
}

h1 {
  color: white;
  text-align: center;
}

p {
  font-family: verdana;
  font-size: 20px;
}
```

Is this code sent to the client/browser?

# Web Dev Refresher

What is this an example of?

```html
<!DOCTYPE html>
<html>
    <head>
        <title>Page Title</title>
    </head>
    <body>

        <h1>This is a Heading</h1>
        <p>This is a paragraph.</p>

    </body>
</html>
```
Is this code sent to the client/browser?

# Web Dev Refresher

What is this an example of?

```python
import os
from flask import Flask, render_template

app = Flask(__name__)

@app.route("/")
def index():
    return render_template("index.html", name="Harold")

@app.route("/about")
def about():
    return render_template("about.html", title="About")

if __name__ == "__main__":
    app.run(debug=True)
```

Is this code sent to the client/browser?

# What does `render_template` allow us to do?

Variables in HTML is essential for making web applications.

We can make our web pages more **dynamic** by inserting variables into HTML. This could be data calculated, entered by the user, and/or pulled from a database!

# What are templates and why use them?

Remember how we had to put the same header in each of our files?

```html
<nav>
 <a href="index.html">Index</a>
 <a href="contact.html">Contact</a>
</nav>
```

Wouldn't it be nice if we could just write such things once and re-use it, like we do with functions, instead of copy-pasting everywhere?

Enter templates!

# Introduction to Flask Templates

- Flask uses **Jinja2** templating engine.
- Allows you to embed Python code in HTML.
- Helps create dynamic web pages with ease.

Example: Make a button once and use it everywhere!

Need to change the button?  Just change it in one place!

# Variables: Key syntax

{{ var_name }} to inject variables into your html

Example:
```html
<h1>Hi, {{ name }} how are you today?</h1>
```

Then pass it in using your route handler:

```python
@app.route('/')
def home():
    return render_template('index.html', name="Mr. Lapinel")
```

# Template key syntax

**Templates**
{% extends 'base.html' %}

**Components**
{% include 'header.html' %}

{% block content %}
{% endblock %}

# Setting Up Templates

You will use the below as starter code in your assignment.

1. Create a `project/`, `templates/` and `components/` directories as indicated below.
2. You will put your app.py in the `project/` directory, and your HTML files in the `templates/` and `components/` directory.

Example folder structure:

```plaintext
project/
   app.py
   templates/
      base.html
      index.html
      about.html
      components/
         header.html
```

# **Create a Base Template**

- Create a file `base.html` in the `templates/` folder.
- Add a basic HTML structure that includes a `{% block content %}` placeholder.

   Example content for `base.html`:

   ```html
   <!DOCTYPE html>
   <html>
   <head>
       <title>{{ title }}</title>
   </head>
   <body>
       <div>
           {% block content %}{% endblock %}
       </div>
   </body>
   </html>
   ```

# **Create Content Pages**

- Create `index.html` and `about.html` files, extending the `base.html` template.
- Use the `{% block content %}` to insert page-specific content.

   Example `index.html`:

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
       <h2>About {{name}}</h2>
       <p>Learn more about our Flask project here.</p>
   {% endblock %}
   ```

# Rendering the Homepage

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

# Rendering About Page

**app.py:**

```python
@app.route('/about')
def about():
    return render_template('about.html', title="About", name="John")
```

# Try it out!

Run app.py and go to the URL indicated in your terminal. Did it work?

# Template Components (Reusable HTML)

- Use `{% include %}` to insert reusable components.

# Step 1: Create the Component

**templates/components/header.html:**

```html
<header>
    <h1>My Website</h1>
    <nav>
        <a href="{{ url_for('home') }}">Home</a> |
        <a href="{{ url_for('about') }}">About</a>
    </nav>
</header>
```

**`url_for()`**

The url_for() is Flask's "url-building" function. You pass in your route handler's function name and it returns the URL.

That way you can just type the url once!

[Flask Docs: url_for](https://flask.palletsprojects.com/en/stable/quickstart/#url-building)

# Step 2: Include the Component in Pages

## **templates/index.html:**

```html
{% extends 'base.html' %}
  {% block content %}
  {% include 'components/header.html %}
      <h2>Welcome to the Homepage!</h2>
      <p>This is the main page of the site.</p>
  {% endblock %}
```

# Step 2: Include the Component in Pages (continued)

# **templates/about.html:**

(Do the same as above!)

# Benefits of Components

- **Reusability**: Avoid duplicating HTML code.
- **Consistency**: Ensure uniform styling across pages.
- **Maintainability**: Update the component once, reflect changes everywhere.

# Assignment 8.4 (Due today)

- Practice creating and using templates and reusable components in Flask.

# Assignment

- Add a component `form.html` using the code below (feel free to add to it)

```html
<form>
    <label for="msg">Please leave me a message here</label>
    <input id="msg" type="text" name="message">
    <button type="submit">Submit</button>
</form>
```

# **Requirements**

- Add a contact page that uses a variable somewhere on the page, extends base.html, and includes the form component
- Add some styling to your page by putting a `styles.css` in a `static/` directory. Be sure to add a link to your head element!
- Flask automatically serves static files from the static directory.

```text
project/
  app.py
    static/
      styles.css
    templates/
      base.html
      index.html
      about.html
      components/
        header.html
```

# **Requirements**

- Verify that your code runs without error, all your components appear on pages as expected and all links work correctly.
- Submit your completed Flask project as a .zip file. (right-click your project folder and then click 'send to compressed' or something like that)
- Bonus: modify your form and server code to handle form submission!

# Next Steps

1. Practice creating templates and components.
2. Explore other Jinja2 features:
   - Conditionals (`{% if %}`)
   - Loops (`{% for %}`)
3. Build a simple multi-page Flask application.

# Resources and tools

[Flask Snippets](https://marketplace.visualstudio.com/items?itemName=cstrap.flask-snippets) Provides nice autocompletion on common Flask things. Example: `finclude` will give you this: {% include "template" %}

[Flask Docs](https://flask.palletsprojects.com/en/stable/)

[MDN Beginner's Tutorial on HTML](https://developer.mozilla.org/en-US/docs/Learn_web_development/Getting_started/Your_first_website/Creating_the_content)

[MDN Beginner's Tutorial on CSS](https://developer.mozilla.org/en-US/docs/Learn_web_development/Core/Styling_basics)
