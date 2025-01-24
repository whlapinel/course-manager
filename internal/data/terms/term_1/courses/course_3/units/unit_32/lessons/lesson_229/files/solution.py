# Import pandas library
import pandas as pd


def main():
    # Task 1: Create a DataFrame from the following dictionary
    data = {
        "Name": ["Alice", "Bob", "Cathy", "David"],
        "Age": [25, 30, 35, 40],
        "City": ["New York", "Los Angeles", "Chicago", "Houston"],
    }
    df = pd.DataFrame(data)
    print("DataFrame from dictionary:\n", df, "\n")

    # Task 2: Load data from a CSV file named 'employees.csv' and display the first 5 rows
    # Assume 'employees.csv' is in the same directory as this script
    # You can create this CSV file or use any existing CSV file you have
    try:
        employees_df = pd.read_csv("employees.csv")
        print("First 5 rows of employees DataFrame:\n", employees_df.head(), "\n")
    except FileNotFoundError:
        print("CSV file not found. Proceeding to the next task.\n")

    # Task 3: Select and print rows from the original DataFrame where Age is greater than 30
    older_than_30 = df[df["Age"] > 30]
    print("People older than 30:\n", older_than_30, "\n")

    # Task 4: Add a new column 'Salary' to the DataFrame with the following values
    df["Salary"] = [50000, 60000, 75000, 80000]
    print("DataFrame with Salary column added:\n", df, "\n")

    # Task 5: Group the original DataFrame by 'City' and calculate the average age
    average_age_by_city = df.groupby("City")["Age"].mean()
    print("Average age by city:\n", average_age_by_city)


if __name__ == "__main__":
    main()
