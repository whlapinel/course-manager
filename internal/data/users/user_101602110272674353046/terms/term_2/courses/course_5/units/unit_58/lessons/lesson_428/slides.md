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

# Lesson 8.5

HTTP Methods and Dynamic Data Retrieval in Flask  

# Warmup: 1/7/25

# Recap  

**What We Learned So Far**:  

- HTML and CSS basics for structuring and styling web pages.  
- Flask templates (`render_template`) for dynamic content.  
- Using reusable components for maintainable code.  

# Objectives for Today  

**By the end of this lesson, you will:**

1. Use forms to request data through URLs.  
2. Dynamically retrieve and display data using query parameters.  

# What Are HTTP Methods?  

- HTTP methods define the action to be performed on a resource.  
- They are the verbs of web communication.  
- Examples: `GET`, `POST`, `PUT`, `DELETE`, `PATCH`.

# Common HTTP Methods  

| **Method** | **Purpose**               | **Example Use Case**                         |
| ---------- | ------------------------- | -------------------------------------------- |
| `GET`      | Retrieve data             | Fetching a webpage or search results.        |
| `POST`     | Submit data to the server | Creating a new user or submitting a form.    |
| `PUT`      | Update/replace data       | Replacing a profile picture.                 |
| `DELETE`   | Remove data               | Deleting a comment or account.               |
| `PATCH`    | Partially update data     | Updating only an email address in a profile. |

# **GET** Requests  

- **Purpose**: Retrieve data from the server.  
- **Key Features**:  
  - Data is sent in the URL as query parameters.  
  - Safe (does not modify server state).  
  - Bookmarkable and cacheable.  

**Example**:  

```plaintext
GET /search?term=python
```

# **POST** Requests  

- **Purpose**: Submit data to the server.  
- **Key Features**:  
  - Data is sent in the request body.  
  - Used for creating or processing new resources.  
  - Not cacheable or bookmarkable.  

**Example**:  

```plaintext
POST /register
Body: { "username": "JohnDoe", "password": "1234" }
```

# **PUT** Requests  

- **Purpose**: Update or replace a resource.  
- **Key Features**:  
  - Replaces the entire resource.  
  - Idempotent (repeated requests produce the same result).  

**Example**:  

```plaintext
PUT /profile/123
Body: { "name": "Jane Doe", "email": "jane@example.com" }
```

# **PATCH** Requests  

- **Purpose**: Partially update a resource.  
- **Key Features**:  
  - Updates only specific fields.  
  - Ideal for small changes.  

**Example**:  

```plaintext
PATCH /profile/123
Body: { "email": "jane.new@example.com" }
```

# **DELETE** Requests  

- **Purpose**: Remove a resource.  
- **Key Features**:  
  - Idempotent.  
  - Cannot be undone!  

**Example**:  

```plaintext
DELETE /profile/123
```

# Mapping HTTP Methods to CRUD  

| **Operation** | **HTTP Method** | **Example**          |
| ------------- | --------------- | -------------------- |
| Create        | `POST`          | Add a new blog post. |
| Read          | `GET`           | View blog posts.     |
| Update        | `PUT`/`PATCH`   | Edit a blog post.    |
| Delete        | `DELETE`        | Remove a blog post.  |

# Choosing the Right Method  

- Use **`GET`** for retrieving data.  
- Use **`POST`** for submitting or creating new data.  
- Use **`PUT`** for replacing entire resources.  
- Use **`PATCH`** for minor updates.  
- Use **`DELETE`** for removing resources.  


# Introduction to GET Requests  

**GET Requests** are used to:  

- Retrieve data without modifying the server's state.  
- Pass data as query parameters visible in the URL.  

**Example URL**:  

```plaintext
http://127.0.0.1:5000/result?term=Do
```

# Updating Our App for Forms  

1. **Add a Form Component**: Create `form.html` in the `components/` directory:  

   ```html
   <form action="{{ url_for('result') }}" method="GET">
       <label for="term">Enter a term:</label>
       <input type="text" id="term" name="term" required>
       <button type="submit">Search</button>
   </form>
   ```

2. **Include the Form in a Page**: Add this to a new page `search.html`:  

   ```html
   {% extends 'base.html' %}
   {% block content %}
       <h2>Search the Do Re Mi Dictionary</h2>
       {% include 'components/form.html' %}
   {% endblock %}
   ```

# Setting Up Routes  

**1. Add a Route for the Form Page**  

```python
@app.route('/search')
def search():
    return render_template('search.html', title="Search")
```

**2. Add a Route to Handle the Result**  

```python
@app.route('/result')
def result():
    term = request.args.get('term', '').capitalize()
    result = do_re_mi.get(term, "Not found in the Do Re Mi dictionary.")
    return render_template('result.html', term=term, result=result, title="Result")
```

# Dictionary for the Do Re Mi App  

Define a dictionary in `app.py`:  

```python
do_re_mi = {
    "Do": "a deer, a female deer",
    "Re": "a drop of golden sun",
    "Mi": "a name I call myself",
    "Fa": "a long, long way to run",
    "So": "a needle pulling thread",
    "La": "a note to follow So",
    "Ti": "a drink with jam and bread",
}
```

# Creating the Result Page  

Create `result.html`:  

```html
{% extends 'base.html' %}
{% block content %}
   <h2>Search Result</h2>
   <p><strong>{{ term }}:</strong> {{ result }}</p>
   <a href="{{ url_for('search') }}">Search Again</a>
{% endblock %}
```

# Running the Application  

**Steps**:  
1. At the top of `app.py` include:

```python
from flask import Flask, render_template, request
```

1. Add the following to the bottom of `app.py`:  

   ```python
   if __name__ == '__main__':
       app.run(debug=True)
   ```  

2. Run the app:  

   ```bash
   python app.py
   ```  

3. Visit the `/search` page, submit a term, and see the result!

# Why Use GET Requests?

**Advantages of GET Requests**:  

- Easy to bookmark or share specific searches.  
- Makes the app transparent to users.  
- Perfect for retrieving and displaying data.

# Assignment  

**Task**:  

1. Add a **Help Page** (`/help`) that explains the available search terms.  
   - Extend `base.html` and include a list of keys from `do_re_mi`.  
   - Example content: "Available terms: Do, Re, Mi, Fa, So, La, Ti."

**************

