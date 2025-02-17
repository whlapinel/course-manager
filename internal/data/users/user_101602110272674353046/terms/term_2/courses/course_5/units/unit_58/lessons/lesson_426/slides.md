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

- Using Python, read and print to the terminal the contents of the file `warmup_data.txt`
  - [Download](https://whlapinel.github.io/courses/python-ii-programming-honors/unit-8/lesson-8.3/files/warmup_data.txt)

- Create a dictionary consisting of words found in the file, where keys are the first letter of the word and values are a list of words starting with that letter.

- Write this dictionary to a new file called `sorted.txt`

# Lesson 3: More on HTML and CSS

## Objective

Strengthen understanding of HTML/CSS for creating aesthetically pleasing static sites.

# Semantic HTML Elements

- "Semantic" elements convey meaning about their content.

## Examples

- `<header>`: Represents introductory content.
- `<footer>`: Represents footer content.
- `<section>`: Groups related content.
- `<article>`: Represents self-contained content.

## Why Use Semantic Elements?

- Improves accessibility.
- Helps with SEO.
- Makes code more readable.

# CSS Layout Techniques

- Flexbox
- Grid

# Flexbox

- Align and distribute items within a container.
- **Properties to Learn:**
  - `display: flex;`
  - `justify-content`
  - `align-items`

# Grid

- Create two-dimensional layouts.
- **Properties to Learn:**
  - `display: grid;`
  - `grid-template-rows`
  - `grid-template-columns`

# Examples of flexbox and grid

[flexbox](https://whlapinel.github.io/courses/python-ii-programming-honors/unit-8/lesson-8.3/files/flex.html)
[grid](https://whlapinel.github.io/courses/python-ii-programming-honors/unit-8/lesson-8.3/files/grid.html)

# Responsive Design

## Why Responsive Design?

- Adapts layouts to different screen sizes.

# Tools

- **Media Queries:**

  ```css
  @media (max-width: 600px) {
    body {
      font-size: 14px;
    }
  }
  ```

- **Units:**
  - Relative units (e.g., `%`, `em`, `rem`).
  - Flexible layout with `minmax()` in Grid.

# Assignment 8.3

1. **HTML Structure**:
   - Use semantic elements like `<header>`, `<section>`, and `<footer>`.
   - Include sections for "About Me," "Projects," and "Contact."

2. **CSS Styling**:
   - Use Flexbox *and* Grid to create a visually appealing layout.
   - Apply at least one media query to make it responsive.

3. **Enhancements**:
   - Add a navigation bar and links.
   - Use colors, fonts, and spacing to improve aesthetics.
