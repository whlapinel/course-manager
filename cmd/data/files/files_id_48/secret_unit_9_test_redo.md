---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# **Challenge: Sensor Log Analyzer**

## **Scenario**

Your company has a network of temperature sensors running on different machines. These sensors periodically record measurements and sometimes generate alerts if something goes wrong. All events are written to a single log file named `sensor.log`. Each log entry appears in one of the following formats:

```bash
[2025-01-15 07:12:00] SENSOR #001 Temperature 72.5
[2025-01-15 07:13:00] SENSOR #002 Temperature 70.9
[2025-01-15 07:14:00] ALERT Overheating in sensor #001
[2025-01-15 07:15:00] SENSOR #001 Temperature 90.2
[2025-01-15 07:16:00] ERROR Sensor #002 offline
```

You need to **parse** these entries, gather various statistics, and produce a **summary report** in a separate file, `summary.txt`.

# **Your Tasks**

1. **Read** the `sensor.log` file.
2. **Analyze** the log entries:
   - **Count** each type of log entry:
     - **SENSOR** (i.e., normal sensor readings)
     - **ALERT**
     - **ERROR**
   - **Extract** all **temperature values** from any `SENSOR` entries.
   - **Identify** the **highest and lowest** temperature readings overall.
   - **Calculate** the **average** temperature reading per sensor. 
     - For instance, if sensor `#001` has recorded `[72.5, 90.2]`, its average is `(72.5 + 90.2) / 2`.
3. **Write** these results (and any other observations you find helpful) to a new file named `summary.txt`.

# **Details and Requirements**

## **File Handling**

- **Open** the file named `sensor.log`.
- If it doesn’t exist, display an error message (e.g. “File `sensor.log` not found!”) and stop.
- **Create** a new file named `summary.txt` to store your findings.

## **Log Entry Analysis**

1. **Count Log Entry Types**  
   - How many `SENSOR` entries are there?  
   - How many `ALERT` entries are there?  
   - How many `ERROR` entries are there?

2. **Temperature Readings**  
   - Extract every temperature from lines that begin with `SENSOR`.  
   - Track each sensor’s readings (e.g., sensor `#001`, sensor `#002`, etc.).  
   - Compute:
     - The **minimum** temperature you encounter (across all sensors).
     - The **maximum** temperature you encounter (across all sensors).
     - The **average** temperature **per sensor**.

3. **Edge Cases**  
   - What if there are no `SENSOR` entries? (Then you can’t compute temperatures—handle carefully.)  
   - What if a sensor appears only once? (Then its average temperature is just that single reading.)

# **Sample `summary.txt` (Hypothetical Example)**

```
Sensor Log Analysis Report
===========================
Log Entry Counts:
 - SENSOR: 4
 - ALERT: 1
 - ERROR: 1

Temperature Stats (All Sensors):
 - Minimum Temperature: 70.9
 - Maximum Temperature: 90.2

Average Temperature per Sensor:
 - #001 -> 81.35
 - #002 -> 70.9
```

> **Note**: The exact numbers above are just an example. Your computed values might differ based on the actual contents of the log.

# **Bonus Challenges**

1. **Sort the sensor list**  
   - In the “Average Temperature per Sensor” section, sort the sensors by their ID (e.g., `#001`, `#002`, etc.) before printing.
2. **Detect Overheating**  
   - If any temperature reading is above a certain threshold (e.g., 85.0), add a note in the report (e.g., “Sensor #001 exceeded safe temperature!”).

# **Assessment Criteria**

1. **File Management**  
   - Opens and reads from `sensor.log`.  
   - Creates and writes to `summary.txt`.  
   - Gracefully handles the case where `sensor.log` does not exist.

2. **Data Processing**  
   - Correctly identifies `SENSOR`, `ALERT`, and `ERROR` lines.  
   - Extracts numerical temperature values and keeps track by sensor ID.  
   - Correctly calculates min, max, average temperatures, and total counts.

3. **Output Quality**  
   - Report in `summary.txt` is well-structured, easy to read, and accurately reflects the data.  
   - Handles unusual or missing data (no sensor lines, no alerts, etc.) without breaking.

# **Starter Code (Optional)**

```python
import os

def read_log_file(file_path):
    if not os.path.exists(file_path):
        print(f"Error: {file_path} was not found.")
        return []
    with open(file_path, "r") as file:
        return file.readlines()

def write_report(file_path, content):
    with open(file_path, "w") as file:
        file.write(content)

# Example usage:
log_entries = read_log_file("sensor.log")
if log_entries:
    # TODO: Parse these log entries and build your final string for summary.txt
    report_content = "Your final report goes here..."
    write_report("summary.txt", report_content)
```

Feel free to expand on this code or write your own. Good luck analyzing those sensor logs!