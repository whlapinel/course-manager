# Solution Activity 3

```python
def extract_temperatures(log_file):
    temperatures = []
    with open(log_file, "r") as file:
        for line in file:
            if "SENSOR" in line:
                parts = line.split()
                temp = float(parts[-1])  # Temperature is the last item in the line
                temperatures.append(temp)
    print("Temperatures:", temperatures)

extract_temperatures("logs.log")
```
