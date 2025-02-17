---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# **Warmup**

## Plot a graph of `y = x²` using `matplotlib.pyplot`

# Agenda

Assignment 6.8

# Reminders

- Unit 6 Project due Friday
- Permission slips

# Session 8: Advanced Data Analysis with Seaborn

## Categorical Data Visualization and Regression Analysis

# Introduction to Advanced Seaborn Techniques

Building on our previous sessions, today we will explore more advanced visualization techniques that are pivotal for categorical and regression analysis.

# Categorical Data Visualization

Understanding how to visualize categorical data can reveal insights into patterns and relationships that are not evident in numerical data.

## Example: Bar Plot

```python
import seaborn as sns
import matplotlib.pyplot as plt
tips = sns.load_dataset('tips')
sns.barplot(x='sex', y='total_bill', data=tips)
plt.title('Average Total Bill for Each Sex')
plt.show()
```

# Regression Plots

Seaborn’s regression plots can be used to identify trends, relationships, and outliers in data.

## Example: Regression Plot

```python
import seaborn as sns
import matplotlib.pyplot as plt
tips = sns.load_dataset('tips')
sns.regplot(x='total_bill', y='tip', data=tips)
plt.title('Regression Plot of Total Bill vs. Tip')
plt.xlabel('Total Bill')
plt.ylabel('Tip')
plt.show()
```

# Swarm and Strip Plots

These plots provide a detailed look at the distribution of categorical data, showing each data point.

## Example: Swarm Plot

```python
sns.swarmplot(x='day', y='total_bill', data=tips)
plt.title('Swarm Plot of Total Bills by Day')
plt.show()
```

# Heatmaps for More Complex Data

Expanding on simple heatmaps, we will use them to explore complex datasets like flight data to find patterns over time.

## Example: Complex Heatmap

```python
flights = sns.load_dataset('flights')
flights_pivot = flights.pivot(index="month", columns="year", values="passengers")
sns.heatmap(flights_pivot, annot=True, fmt="d")
plt.title('Flights Passenger Traffic (1953 - 1960)')
plt.show()
```

# Session Summary

Today, we:

- Learned to visualize categorical data effectively.
- Explored advanced regression analysis techniques.
- Demonstrated the utility of swarm and strip plots.
- Used heatmaps to analyze complex datasets.

**Next Session**: We will focus on time series analysis and how to handle temporal data, which is essential for many real-world data science applications.
