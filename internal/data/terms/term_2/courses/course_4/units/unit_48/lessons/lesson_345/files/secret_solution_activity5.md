# Solution Activity 5

```python
def group_by_sensor(log_file):
    sensor_data = {}
    with open(log_file, "r") as file:
        for line in file:
            if "SENSOR" in line:
                parts = line.split()
                sensor_id = parts[2]  # Extract sensor ID
                temp = float(parts[-1])  # Extract temperature
                if sensor_id not in sensor_data:
                    sensor_data[sensor_id] = []
                sensor_data[sensor_id].append(temp)
    print(sensor_data)

group_by_sensor("logs.log")
```
