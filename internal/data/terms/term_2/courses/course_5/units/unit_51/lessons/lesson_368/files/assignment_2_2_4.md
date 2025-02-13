# Practice Problems: Using Modules & datetime

## Problem 1: Importing a Module
Import the `math` module and use it to find the square root of 144.

**Hint:** Use `import math` and find the appropriate function in the `math` module.

---

## Problem 2: Importing Specific Functions
Import only the `choice` function from the `random` module and use it to pick a random item from a list: `['apple', 'banana', 'cherry']`.

**Hint:** Use `from random import choice` to access the function directly.

---

## Problem 3: Importing with an Alias
Import the `datetime` module with an alias `dt` and print the current date and time.

**Hint:** Use `import datetime as dt` and find a function that gives you the current date and time.

---

## Problem 4: Creating Your Own Module
Create a module called `greetings.py` that contains a function `say_hello(name)` which prints `Hello, <name>!`. Import and use this module in another script.

**Hint:** Create a `.py` file, define a function, then import it using `import greetings`.

---

## Problem 5: Formatting Dates
Use the `datetime` module to print today’s date in the format `YYYY-MM-DD`.

**Hint:** Use `strftime()` to format the date.

---

## Problem 6: Using `__name__ == "__main__"`
Write a script that defines a function `display_message()` that prints `This script is running directly!`. Ensure that the function only runs when the script is executed directly and not when imported.

**Hint:** Use `if __name__ == "__main__":` to control execution.

---

## Problem 7: Calculating Age
Write a script that asks the user for their birth year and calculates their age using the `datetime` module.

**Hint:** Get the current year from `datetime.date.today().year` and subtract the birth year.

---

## Problem 8: Importing Everything (Use with Caution)
Import all functions from the `math` module and use the `ceil()` function to round `3.4` up to the nearest integer.

**Hint:** Use `from math import *`, but be aware of potential naming conflicts.

---

## Problem 9: Finding the Difference Between Two Dates
Use the `datetime` module to find the number of days between January 1, 2023, and today.

**Hint:** Subtract two `datetime.date` objects to get the difference.

---

## Problem 10: Generating a Timestamp
Write a script that prints the current timestamp (seconds since the epoch) using the `datetime` module.

**Hint:** Use `datetime.datetime.timestamp()` on the current datetime object.
