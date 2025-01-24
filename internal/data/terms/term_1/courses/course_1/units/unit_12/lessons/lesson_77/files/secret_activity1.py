file = open("example.log", "r")
file.close()

count = 0

with open("example.log") as file:
    lines = file.readlines()
    count = len(lines)
    print(count)

with open("summary.txt", "w") as file:
    prefix = "total number of lines in example.log: "
    file.write(prefix + str(count))
