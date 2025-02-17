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

Create a pandas dataframe from a dictionary.

The dataframe should consist of 3 columns and 3 rows.

Write the dataframe to a .csv file.

# Session 5: Advanced Data Manipulation with pandas

## Handling Missing Data, Merging, and Aggregation

# Review of pandas Basics

In our previous session, we introduced:

- Creating DataFrames
- Importing data from CSV files
- Basic data filtering and selection

Today, we will build on that and learn advanced manipulation techniques!

# Handling Missing Data

In real-world data, missing values are common. pandas provides several ways to deal with them.

## Example: Handling Missing Values

```python
import pandas as pd
data = {'Name': ['John', 'Anna', 'Peter', 'Linda'],
        'Age': [28, None, 34, 32],
        'City': ['New York', 'Paris', None, 'Berlin']}
df = pd.DataFrame(data)

# Fill missing values with a placeholder
df_filled = df.fillna('Unknown')
print(df_filled)

# Drop rows with missing values
df_dropped = df.dropna()
print(df_dropped)
```

# Merging DataFrames

Merging (or joining) DataFrames is a powerful way to combine datasets based on a common column.

## Example: Merging Two DataFrames

```python
df1 = pd.DataFrame({'ID': [1, 2, 3], 'Name': ['John', 'Anna', 'Peter']})
df2 = pd.DataFrame({'ID': [1, 2, 4], 'City': ['New York', 'Paris', 'Berlin']})

merged_df = pd.merge(df1, df2, on='ID', how='inner')
print(merged_df)
```

# Sorting Data

pandas makes it easy to sort data either by column values or index.

## Example: Sorting a DataFrame by Column

```python
# Sorting by age
sorted_df = df.sort_values(by='Age', ascending=False)
print(sorted_df)
```

# Aggregation and Grouping

Aggregation allows you to calculate summary statistics (e.g., mean, sum) for each group of data.

## Example: Aggregating Data by Group

```python
grouped = df.groupby('City').agg({'Age': 'mean'})
print(grouped)
```

# Session Summary

Today, we explored more advanced data manipulation techniques in pandas:

- Handling missing data
- Merging DataFrames
- Sorting data and performing aggregation

**Next Session**: We will continue with more complex data analysis using pandas and start exploring visualizations with matplotlib!
