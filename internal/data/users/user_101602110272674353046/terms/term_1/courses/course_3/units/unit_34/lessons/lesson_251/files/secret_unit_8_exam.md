# Unit 8 Exam

## **Build a Flask-Based Dynamic Web Application**

### Overview

In this assignment, you will create a **read-only Flask application** that displays and interacts with backend data dynamically. The app will allow users to search and explore data from a predefined dataset using a simple, intuitive interface.

### Objectives

1. Set up a basic Flask web application.
2. Create routes that serve dynamic content.
3. Use templates to render HTML with data passed from Flask.
4. Handle user input using GET requests.

### The Dataset

You will use a dictionary of **planets in the Solar System** as your dataset. Here is the dataset to include in your application:

```python
planets = {
    "Mercury": {"Type": "Terrestrial", "Day Length (Earth days)": 58.6, "Moons": 0},
    "Venus": {"Type": "Terrestrial", "Day Length (Earth days)": 243, "Moons": 0},
    "Earth": {"Type": "Terrestrial", "Day Length (Earth days)": 1, "Moons": 1},
    "Mars": {"Type": "Terrestrial", "Day Length (Earth days)": 1.03, "Moons": 2},
    "Jupiter": {"Type": "Gas Giant", "Day Length (Earth days)": 0.41, "Moons": 79},
    "Saturn": {"Type": "Gas Giant", "Day Length (Earth days)": 0.45, "Moons": 83},
    "Uranus": {"Type": "Ice Giant", "Day Length (Earth days)": 0.72, "Moons": 27},
    "Neptune": {"Type": "Ice Giant", "Day Length (Earth days)": 0.67, "Moons": 14}
}
```

### Application Requirements

1. **Homepage**:
   - Display a list of all the planets with links to detailed pages for each.
   - Use Flask’s `url_for()` to dynamically generate links.

2. **Search Page**:
   - Include a search form where users can enter a planet name.
   - Display results dynamically on the same page using a GET request.

3. **Navigation**:
   - Add a simple navigation bar to link the homepage, search page, and planet details.

### Step-by-Step Instructions

#### 1. **Set Up Flask App**

- Create a Flask application with the necessary routes.
- Store the dataset in your `app.py` file.

#### 2. **Build Templates**

- Create a `base.html` with a navigation bar.
- Create `index.html` to display the list of planets.
- Create `search.html` for the search functionality.
- Create `result.html` for displaying search results.

#### 3. **Implement Routes**

- `/`: Homepage that lists all planets.
- `/search`: A page with a search bar.
- `/result`: A page to display search results.

#### 4. **Dynamic Functionality**

- Use Flask’s `render_template()` to inject data into templates.
- Use query parameters (`request.args`) to handle the search functionality.

### Example Features

1. **Homepage (`/`)**
   - Explain what the site is for

2. **Search Page (`/search`)**

   - Display search results dynamically:

### Example Submission Requirements

1. **Folder Structure**:

   ```text
   project/
       app.py
       templates/
           base.html
           index.html
           search.html
           result.html
       static/
           styles.css
   ```

2. **Styling**:
   - Use a `styles.css` file to style the pages.
   - Add a navigation bar for easy access to pages.

3. **Testing**:
   - Verify all routes work without errors.
   - Test the search functionality with valid and invalid input.

4. **Submission**:
   - Submit a `.zip` file of your project folder.

### Grading Rubric

| **Criteria**                                                   | **Points** |
| -------------------------------------------------------------- | ---------- |
| Proper folder structure                                        | 10         |
| Flask app with functional routes                               | 30         |
| Dynamic templates (`index.html`, `result.html`, `search.html`) | 30         |
| Navigation and usability                                       | 10         |
| Styling with CSS                                               | 10         |
| Error handling (e.g., 404 for invalid planet)                  | 10         |
