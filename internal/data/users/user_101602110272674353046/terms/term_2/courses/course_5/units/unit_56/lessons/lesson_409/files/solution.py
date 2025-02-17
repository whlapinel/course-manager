# Import necessary libraries
import pandas as pd
import matplotlib.pyplot as plt
import numpy as np


def main():
    # Task 1: Load a large dataset (you can download any dataset from online data repositories like Kaggle)
    # For example, use a dataset on 'Video Game Sales' or any other dataset of interest
    data = pd.read_csv("your_dataset.csv")
    print("Data loaded successfully.")

    # Task 2: Filter the dataset to include only the entries with sales greater than 1 million units
    high_sales = data[
        data["Global_Sales"] > 1.0
    ]  # Adjust column name as per your dataset
    print("Filtered high sales data:\n", high_sales.head())

    # Task 3: Create a new column that categorizes the sales into 'Low', 'Medium', and 'High'
    bins = [0, 0.5, 1.5, np.inf]
    names = ["Low", "Medium", "High"]
    data["Sales_Category"] = pd.cut(data["Global_Sales"], bins, labels=names)
    print(
        "Data with sales categories:\n", data[["Global_Sales", "Sales_Category"]].head()
    )

    # Task 4: Merge this dataset with another dataset that includes additional information
    # For example, merge with a dataset that includes developer information
    developer_data = pd.read_csv(
        "developer_dataset.csv"
    )  # This should be pre-downloaded
    merged_data = pd.merge(
        data, developer_data, on="Game_ID"
    )  # Adjust key column as necessary
    print("Merged dataset:\n", merged_data.head())

    # Task 5: Create a scatter plot of two numerical variables from the dataset
    plt.scatter(
        data["User_Score"], data["Critic_Score"]
    )  # Adjust column names as per your dataset
    plt.title("User Score vs Critic Score")
    plt.xlabel("User Score")
    plt.ylabel("Critic Score")
    plt.show()

    # Task 6: Create a histogram of a numerical column, e.g., global sales
    data["Global_Sales"].hist(bins=20)
    plt.title("Distribution of Global Sales")
    plt.xlabel("Sales (Millions)")
    plt.ylabel("Number of Games")
    plt.show()


if __name__ == "__main__":
    main()
