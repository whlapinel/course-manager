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

In a jupyter notebook: Create a dataframe of your own and use df.describe() on it.

# Session 7: Advanced Visualization Techniques

## Enhancing Data Insights with Matplotlib and Introduction to Seaborn

# Recap of Previous Session

Last session, we introduced basic plotting with matplotlib:
- Scatter plots
- Histograms
- Customizing plots with titles, labels, and styling

Today, we'll build on those skills and introduce new types of visualizations.

# Advanced Matplotlib Visualizations
Exploring more complex and detailed plots to better understand data relationships and distributions.

## Example: Box Plots
```python
import matplotlib.pyplot as plt
import numpy as np

data = np.random.normal(size=100)
plt.boxplot(data)
plt.title('Box Plot of Normally Distributed Data')
plt.ylabel('Values')
plt.show()
```

# Introduction to Seaborn
Seaborn is built on matplotlib and provides a high-level interface for drawing attractive statistical graphics.

## Example: Violin Plot
```python
import seaborn as sns
data = sns.load_dataset('tips')
sns.violinplot(x='day', y='total_bill', data=data)
plt.title('Violin plot of Total Bills by Day')
plt.show()
```

# Pair Plots
Ideal for exploring correlations between multidimensional data.

## Example: Pair Plot
```python
sns.pairplot(data, hue='sex')
plt.show()
```

# Heatmaps
Useful for visualizing matrix-like data, correlation matrices, and for displaying covariance.

## Example: Heatmap
```python
import seaborn as sns
import numpy as np

correlation_matrix = np.corrcoef(np.random.rand(10, 10))
sns.heatmap(correlation_matrix, annot=True)
plt.title('Heatmap of Correlation Matrix')
plt.show()
```

# Session Summary
Today, we advanced our visualization capabilities:
- Learned more complex matplotlib plots
- Introduced seaborn for statistical graphics
- Explored pair plots, violin plots, and heatmaps to enhance our data analysis

**Next Session**: We will delve deeper into data analysis using seaborn and start integrating these visualizations into our data science projects.
