# Solution Activity 4

```python
def calculate_stats(temperatures):
    if not temperatures:
        print("No temperatures provided.")
        return
    minimum = min(temperatures)
    maximum = max(temperatures)
    average = sum(temperatures) / len(temperatures)
    print(f"Min: {minimum}, Max: {maximum}, Avg: {average}")

temperatures = [72.5, 70.9, 90.2, 85.1]
calculate_stats(temperatures)
```
