# SQL Practice Problems

## Instructions
- Use [sqliteonline.com](https://sqliteonline.com/) to solve these problems.
- Write your SQL queries in a `.sql` file and submit your solutions.

## Problem 1: Create a Table
Create a table called `employees` with the following columns:
- `id` (INTEGER, Primary Key)
- `name` (TEXT)
- `position` (TEXT)
- `salary` (INTEGER)

## Problem 2: Insert Data
Insert the following records into the `employees` table:

| id | name   | position     | salary |
|----|--------|-------------|--------|
| 1  | Alice  | Manager     | 70000  |
| 2  | Bob    | Developer   | 60000  |
| 3  | Charlie| Designer    | 50000  |

## Problem 3: Retrieve Data
Write a query to retrieve all employee names and their positions.

## Problem 4: Filter Data
Write a query to get all employees earning more than 55,000.

## Problem 5: Update Data
Change Bob's salary to 65,000.

## Problem 6: Delete Data
Remove the employee named Charlie from the table.

## Problem 7: Create Another Table
Create a table called `departments` with the following columns:
- `id` (INTEGER, Primary Key)
- `department_name` (TEXT)

## Problem 8: Insert More Data
Insert these records into `departments`:

| id | department_name |
|----|---------------|
| 1  | Engineering  |
| 2  | Design       |

## (Bonus) Problem 9: Join Tables
Write a query to join `employees` with `departments`, assuming:
- Managers and Developers belong to Engineering (id=1)
- Designers belong to Design (id=2)

Display employee names along with their department names.

## (Bonus) Problem 10: Aggregate Function
Write a query to calculate the average salary of all employees.

## Submission
- Save all your queries in a single `.sql` file.
- Submit the file following the given instructions.
