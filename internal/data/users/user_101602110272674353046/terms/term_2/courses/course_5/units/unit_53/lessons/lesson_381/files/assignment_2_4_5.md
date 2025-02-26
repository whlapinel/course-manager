# Assignment 4.5: SQLite in Python - Practice Problems

## Instructions
- Use Python's built-in `sqlite3` module.
- Create a new database file called `dc_heroes.db`.
- Write a Python script to complete each problem.

## Problem 1: Creating a Table
Write a Python script that:
1. Connects to `dc_heroes.db`.
2. Creates a table `heroes` with columns:
   - `id` (INTEGER, PRIMARY KEY)
   - `name` (TEXT)
   - `alias` (TEXT)
   - `power_level` (INTEGER)

## Problem 2: Inserting Data
Insert the following DC heroes into the `heroes` table:

| id | name         | alias         | power_level |
|----|-------------|--------------|-------------|
| 1  | Clark Kent  | Superman      | 100         |
| 2  | Bruce Wayne | Batman        | 85          |
| 3  | Diana       | Wonder Woman  | 95          |

## Problem 3: Retrieving Data
Write a query to fetch all superheroes and display their real names, aliases, and power levels.

## Problem 4: Filtering Data
Retrieve superheroes with a power level above 90.

## Problem 5: Updating Data
Batman just got an upgraded suit! Update his power level to 90.

## Problem 6: Deleting Data
The multiverse collapsed, and Wonder Woman was erased! Remove her from the table.

## Problem 7: Creating Another Table
Create a table `teams` with columns:
- `id` (INTEGER, PRIMARY KEY)
- `team_name` (TEXT)

## Problem 8: Inserting into `teams`
Insert these teams:

| id | team_name          |
|----|-------------------|
| 1  | Justice League    |
| 2  | Teen Titans      |

## Problem 9: Joining Tables
Assume:
- Superman and Batman are in the Justice League (id=1)
- Wonder Woman was also in the Justice League (id=1)

Write a query to join `heroes` and `teams` to show each hero’s team.

## Problem 10: Aggregate Functions
Write a query to calculate the average power level of all superheroes.

## Submission
- Ensure each problem is implemented as a separate function.
- Submit your `.py` file as per instructions.
