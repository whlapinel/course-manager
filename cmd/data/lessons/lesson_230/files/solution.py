# Import pandas library
import pandas as pd


def main():
    # Task 1: Create a DataFrame with missing values and fill them
    data = {
        "Name": ["Alice", "Bob", None, "David"],
        "Age": [25, None, 35, 40],
        "City": ["New York", "Los Angeles", None, "Houston"],
    }
    df = pd.DataFrame(data)
    print("Original DataFrame with missing values:\n", df)
    df_filled = df.fillna({"Name": "Unknown", "Age": 30, "City": "Unknown"})
    print("DataFrame after filling missing values:\n", df_filled)

    # Task 2: Merge two DataFrames
    df1 = pd.DataFrame(
        {"ID": [1, 2, 3, 4], "Product": ["Apples", "Bananas", "Carrots", "Dates"]}
    )
    df2 = pd.DataFrame({"ID": [3, 4, 5, 6], "Price": [1.25, 0.99, 2.50, 5.00]})
    merged_df = pd.merge(df1, df2, on="ID", how="inner")
    print("Merged DataFrame on ID:\n", merged_df)

    # Task 3: Sort a DataFrame by multiple columns
    df_to_sort = pd.DataFrame(
        {
            "Name": ["Alice", "Bob", "David", "Cathy"],
            "Age": [24, 30, 29, 22],
            "Salary": [70000, 48000, 52000, 58000],
        }
    )
    sorted_df = df_to_sort.sort_values(by=["Age", "Salary"], ascending=[True, False])
    print("DataFrame sorted by Age and Salary:\n", sorted_df)

    # Task 4: Group and aggregate data
    sales_data = pd.DataFrame(
        {
            "City": ["New York", "Los Angeles", "New York", "Chicago"],
            "Revenue": [1000, 2300, 1200, 1600],
        }
    )
    aggregated_data = sales_data.groupby("City").agg({"Revenue": ["mean", "sum"]})
    print("Aggregated Revenue by City:\n", aggregated_data)


if __name__ == "__main__":
    main()
