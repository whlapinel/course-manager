sensor_temps: dict[str, list[float]] = {}

with open("logs.log", "r") as file:

    for line in file:
        if "SENSOR" in line:
            split_line = line.split()
            id = split_line[3]
            temp = float(split_line[5])
            if id in sensor_temps:
                sensor_temps[id].append(temp)
            else:
                sensor_temps[id] = [temp]

with open("summary.txt", "w") as report_file:
    report_file.write("Sensor Temps:\n")
    for key, val in sensor_temps.items():
        line = key + str(val)
        report_file.write(line)
        report_file.write("\n")
