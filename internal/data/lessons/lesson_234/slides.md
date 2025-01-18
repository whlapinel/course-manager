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

- Plot a graph that has 2 lines on the same graph, y=2x and y=3x
- Label the axes and give it a title

# Agenda

- Nothing due today, but lots to do:
  - PE2 Module 3 Quiz due Thursday
  - PE2 Module 3 Test due Friday
  - Unit 6 Project due Friday (perform grade)

# Session 9: Time Series Analysis

## Handling and Visualizing Temporal Data

# Introduction to Time Series Analysis

Time series analysis involves techniques for analyzing time series data in order to extract meaningful statistics and other characteristics.

# Loading and Handling Time Series Data

The first step in working with time series data is to correctly load and handle it, ensuring that datetime elements are recognized and manipulated effectively.

## Example: Loading Time Series Data

```python
import pandas as pd

# Load data
data = pd.read_csv('time_series_data.csv', parse_dates=True, index_col='Date')
print(data.head())
```

# Time-Based Indexing

Pandas allows for indexing data frames with datetime values, which simplifies retrieving data for a specific time frame.

## Example: Indexing by Date

```python
# Access data for a specific year
year_data = data['2020']
print(year_data)
```

# Resampling Time Series Data

Resampling involves changing the frequency of your time series observations.

## Example: Monthly Averages

```python
monthly_data = data.resample('M').mean()
print(monthly_data)
```

# Time Series Visualization

Visualizing time series data effectively helps in understanding trends, cycles, and seasonal variations.

## Example: Time Series Plot

```python
import matplotlib.pyplot as plt

data.plot()
plt.title('Time Series Data Over Time')
plt.xlabel('Date')
plt.ylabel('Value')
plt.show()
```

# Rolling and Expanding Metrics

Understanding rolling and expanding metrics can help in analyzing trends over sliding windows of time.

## Example: Rolling Average

```python
rolling_data = data.rolling(window=12).mean()  # 12-period rolling average
rolling_data.plot()
plt.title('Rolling Average')
plt.show()
```

# Decomposing Time Series

Decomposing time series data into its components is essential for understanding underlying patterns.

## Example: Seasonal Decomposition

```python
from statsmodels.tsa.seasonal import seasonal_decompose

result = seasonal_decompose(data, model='additive')
result.plot()
plt.show()
```

# Session Summary

Today, we:

- Introduced time series analysis concepts.
- Learned how to load, handle, and visualize time series data.
- Explored advanced techniques like resampling and decomposing time series.

**Next Session**: We will delve into predictive models for time series data, focusing on forecasting and model evaluation.
