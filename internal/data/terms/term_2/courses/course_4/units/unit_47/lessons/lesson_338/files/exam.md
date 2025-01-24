---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# **Challenge: Log File Analyzer**

## **Scenario:**

You are working as a software developer for a company that needs to analyze its server logs. The log file contains entries in the following format:

```bash
[2024-12-01 10:15:32] INFO User 'Alice' logged in.
[2024-12-01 10:20:45] ERROR Failed to connect to database.
[2024-12-01 10:45:10] INFO User 'Bob' logged out.
[2024-12-01 11:05:21] WARNING Disk space running low.
[2024-12-01 11:25:34] INFO User 'Alice' logged out.
```

# Challenge

Your task is to write a program that:

1. Reads the log file.
2. Processes the log entries to extract and summarize key information.
3. Outputs a report.

# **Requirements:**

# File Operations

- Read the log file.
- Create a summary report and save it as a new text file (`report.txt`).

# Analysis Tasks

- Count the number of each type of log entry (`INFO`, `ERROR`, `WARNING`).
- Identify the users who logged in and out and the number of times each user logged in.
- Extract all unique timestamps when `ERROR` messages occurred.

# Format (Saved to `report.txt`)

 ```bash
 Server Log Analysis Report
 ===========================
 Log Entry Counts:
 - INFO: 3
 - ERROR: 1
 - WARNING: 1

 User Login Activity:
 - Alice: 1 login(s)
 - Bob: 1 login(s)

 Error Timestamps:
 - 2024-12-01 10:20:45
 ```

# **Bonus Challenges:**

- Sort the users alphabetically in the report.
- Calculate the total time each user was logged in (assume logs are in order).

# **Assessment Criteria**

# File Handling

- Successfully read from and write to files.
- Properly handle missing files (e.g., display an error if the log file doesn’t exist).

# Data Processing

- Correctly parse log entries and extract required information.
- Use Python data structures like lists and dictionaries effectively.

# Output Quality

- The generated report is clear, well-organized, and matches the format.
- Handles edge cases (e.g., no `ERROR` messages in the log).

# **Starter Code**

```python
import os

def read_log_file(file_path):
    if not os.path.exists(file_path):
        print(f"File {file_path} does not exist.")
        return []
    with open(file_path, "r") as file:
        return file.readlines()

def write_report(file_path, content):
    with open(file_path, "w") as file:
        file.write(content)

# Example usage:
log_entries = read_log_file("server.log")
# Process log_entries to generate the report...
```
