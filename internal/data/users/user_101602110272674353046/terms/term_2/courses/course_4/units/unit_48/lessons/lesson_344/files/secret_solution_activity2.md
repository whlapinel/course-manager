# Solution Activity 2

```python
def count_log_types(log_file):
    counts = {"SENSOR": 0, "ALERT": 0, "ERROR": 0}
    with open(log_file, "r") as file:
        for line in file:
            if "SENSOR" in line:
                counts["SENSOR"] += 1
            elif "ALERT" in line:
                counts["ALERT"] += 1
            elif "ERROR" in line:
                counts["ERROR"] += 1
    print(counts)

count_log_types("logs.log")
```
