# Solution Activity 1

```python
def count_lines(input_file, output_file):
    with open(input_file, "r") as file:
        lines = file.readlines()
    
    total_lines = len(lines)
    
    with open(output_file, "w") as file:
        file.write(f"Total lines: {total_lines}")

count_lines("example.log", "summary.txt")
```
