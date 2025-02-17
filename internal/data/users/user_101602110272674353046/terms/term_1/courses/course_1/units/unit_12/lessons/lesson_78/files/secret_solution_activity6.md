# Solution Activity 6

```python
def write_summary_report(sensor_data, output_file):
    with open(output_file, "w") as file:
        file.write("Sensor Log Analysis Report\n")
        file.write("===========================\n")
        for sensor, temps in sensor_data.items():
            minimum = min(temps)
            maximum = max(temps)
            average = sum(temps) / len(temps)
            file.write(f"Sensor {sensor} -> Min: {minimum}, Max: {maximum}, Avg: {average:.2f}\n")

sensor_data = {
    "#001": [72.5, 90.2],
    "#002": [70.9]
}
write_summary_report(sensor_data, "summary.txt")
```
