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


entries = read_log_file("log.txt")
type_count = {}
user_logins = {}
user_logouts = {}
error_timestamps = []

for entry in entries:
    entry_type = entry.split()[2]
    if entry_type in type_count:
        type_count[entry_type] += 1
    else:
        type_count[entry_type] = 1
    if entry_type == "ERROR":
        error_timestamps.append(entry.split()[:2])
    if entry_type == "INFO":
        user = entry.split()[4]
        in_out = entry.split()[6]
        if in_out == "in.":
            if user in user_logins:
                user_logins[user] += 1
            else:
                user_logins[user] = 1
        elif in_out == "out.":
            if user in user_logouts:
                user_logouts[user] += 1
            else:
                user_logouts[user] = 1

report_header = f"""
 Server Log Analysis Report
 ===========================
 Log Entry Counts:
 """

for e_type, count in type_count.items():
    print(f" - {e_type}: {count}")

print("User Login Activity:")
for name, count in user_logins.items():
    print(f" - {name}: {count} login(s)")

print("User Logout Activity:")
for name, count in user_logouts.items():
    print(f" - {name}: {count} logouts(s)")

print("Error Timestamps:")
for stamp in error_timestamps:
    print(f" - {stamp}")
