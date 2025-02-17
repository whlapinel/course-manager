with open("logs.log", "r") as file:
    entry_counts: dict[str, int] = {}
    lines = file.readlines()
    entry_type = ""
    for line in lines:
        entry_type = line.split()[2]
        if entry_type in entry_counts:
            entry_counts[entry_type] += 1
        else:
            entry_counts[entry_type] = 1

print(entry_counts)
