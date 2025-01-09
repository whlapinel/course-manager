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

# Introduction to HTML and CSS

Learn how to create beautiful and functional websites!

# What is HTML?

- **HTML (HyperText Markup Language)** is the standard language for creating web pages.
- It defines the structure of a webpage.
- Uses a system of **tags** to organize content.

# Basic HTML Structure

```html
<!DOCTYPE html>
<html>
<head>
    <title>My First Page</title>
</head>
<body>
    <h1>Welcome to My Website</h1>
    <p>This is my first HTML page.</p>
</body>
</html>
```

# Common HTML Tags

| Tag            | Purpose                        |
| -------------- | ------------------------------ |
| `<h1> to <h6>` | Headings (largest to smallest) |
| `<p>`          | Paragraphs                     |
| `<a>`          | Links                          |
| `<img>`        | Images                         |
| `<ul>`         | Unordered lists                |
| `<ol>`         | Ordered lists                  |
| `<div>`        | Division (grouping content)    |

# What is CSS?

- **CSS (Cascading Style Sheets)** controls the style and layout of a webpage.
- Used to make your webpage **beautiful** and **responsive**.
- Separates content (HTML) from presentation (CSS).

# Basic CSS Syntax

```css
selector {
    property: value;
}
```

Example:

```css
body {
    background-color: lightblue;
}
```

# Adding CSS to HTML

## Inline CSS

```html
<p style="color: red;">This is red text.</p>
```

## Internal CSS

```html
<style>
    h1 {
        font-family: Arial, sans-serif;
    }
</style>
```

## External CSS

```html
<link rel="stylesheet" href="styles.css">
```

# Common CSS Properties

| Property           | Description              |
| ------------------ | ------------------------ |
| `color`            | Text color               |
| `background-color` | Background color         |
| `font-size`        | Size of text             |
| `margin`           | Space outside an element |
| `padding`          | Space inside an element  |
| `border`           | Border around an element |

# The Box Model

Every HTML element is a **box** made up of:

1. **Content**
2. **Padding**
3. **Border**
4. **Margin**

![bg right:50% contain](./box-model.png)

# CSS Selectors

- **Type Selector:** Targets elements by tag.

  ```css
  p {
      color: blue;
  }
  ```

- **Class Selector:** Targets elements with a class.

  ```css
  .my-class {
      font-size: 20px;
  }
  ```

- **ID Selector:** Targets elements with an ID.

  ```css
  #my-id {
      background-color: yellow;
  }
  ```

# Responsive Design

- Makes your website look good on all devices.
- Use **media queries** for responsiveness:

  ```css
  @media (max-width: 600px) {
      body {
          background-color: lightgray;
      }
  }
  ```

# Practice Exercise: HTML

Create a simple HTML page with:

1. A title
2. A heading
3. A paragraph
4. An image
5. A link

# Practice Exercise: CSS

Style your page with:

1. A background color
2. Styled text (font, color, size)
3. Padding and margin adjustments
4. A border around an element

# Summary

- **HTML** defines the structure of a webpage.
- **CSS** defines the style of a webpage.
- Together, they create visually appealing, functional websites.