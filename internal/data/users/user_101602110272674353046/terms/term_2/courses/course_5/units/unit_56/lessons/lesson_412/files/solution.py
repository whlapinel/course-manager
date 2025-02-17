# Import necessary libraries
import pandas as pd
import matplotlib.pyplot as plt
import numpy as np


def main():
    # Task 1: Load time series data
    # You can use any time series dataset (e.g., stock prices, weather data)
    data = pd.read_csv("time_series_data.csv", parse_dates=True, index_col="Date")
    print("Data loaded successfully.")

    # Task 2: Perform time-based indexing and slicing
    # Get data for a specific year
    data_2020 = data["2020"]
    print("Data for 2020:\n", data_2020.head())

    # Task 3: Resample the data to find monthly averages
    monthly_avg = data.resample("M").mean()
    print("Monthly averages:\n", monthly_avg.head())

    # Task 4: Create a time series plot
    data["2020"].plot()
    plt.title("Time Series Plot for 2020")
    plt.xlabel("Date")
    plt.ylabel("Value")
    plt.show()

    # Task 5: Calculate and plot a rolling average
    rolling_avg = (
        data["Value"].rolling(window=30).mean()
    )  # Adjust 'Value' to your specific column
    rolling_avg.plot()
    plt.title("30-Day Rolling Average")
    plt.xlabel("Date")
    plt.ylabel("Average Value")
    plt.show()

    # Task 6: Decompose the time series data and plot the components
    from statsmodels.tsa.seasonal import seasonal_decompose

    result = seasonal_decompose(
        data["Value"], model="additive"
    )  # Adjust 'Value' to your specific column
    result.plot()
    plt.show()


if __name__ == "__main__":
    main()
