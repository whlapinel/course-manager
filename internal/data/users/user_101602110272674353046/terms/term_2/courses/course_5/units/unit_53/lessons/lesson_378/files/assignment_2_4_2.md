# Python 2: Assignment 4.2

## **Challenge 1: Read and Print a CSV File**
Write a Python script that reads a CSV file (`data.csv`) and prints its contents line by line.

Example CSV (`data.csv`):
```csv
Name,Age,City
Alice,25,New York
Bob,30,Los Angeles
Charlie,22,Chicago
```
**Hint:** Use `csv.reader()`.

---

## **Challenge 2: Count the Number of Rows**
Modify your script to count and print how many rows are in the CSV file (excluding the header).

**Hint:** Use a counter variable while iterating through the rows.

---

## **Challenge 3: Write Data to a CSV File**
Create a Python script that writes the following data to a new CSV file (`new_data.csv`):

```csv
Product,Price,Quantity
Laptop,1000,5
Phone,500,10
Tablet,300,7
```
**Hint:** Use `csv.writer()` and `writerows()`.

---

## **Challenge 4: Filter Data from a CSV File**
Write a script that reads `data.csv` and prints only the names of people who are younger than 30.

**Hint:** Use `csv.DictReader()` to access values by column name.

---

## **Challenge 5: Add a New Row to an Existing CSV**
Modify `new_data.csv` by adding a new product: `"Headphones", 150, 20`.

**Hint:** Open the file in `append` mode (`"a"`) and use `csv.writer()`.

---

## **Challenge 6: Convert CSV to Dictionary**
Read `data.csv` and store its contents in a dictionary where names are the keys and the rest of the row is stored as values.

Example output:
```python
{
    "Alice": {"Age": 25, "City": "New York"},
    "Bob": {"Age": 30, "City": "Los Angeles"},
    "Charlie": {"Age": 22, "City": "Chicago"}
}
```
**Hint:** Use `csv.DictReader()`.

---

## **Challenge 7: Calculate the Average Age**
Read `data.csv` and compute the average age of all people in the file.

**Hint:** Convert the age values to integers and sum them.
