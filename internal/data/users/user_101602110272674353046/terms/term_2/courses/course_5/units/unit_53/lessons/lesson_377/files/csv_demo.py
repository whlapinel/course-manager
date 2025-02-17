import csv


def main():
    data = [
        ["Name", "Age", "City"],
        ["Alice", 30, "New York"],
        ["Bob", 25, "San Francisco"],
    ]
    write_file(data)
    print("Data written to csv file!")
    input("Press any key to continue")
    data = read_file()
    print(data)


def write_file(data):
    with open("example.csv", mode="w", newline="") as file:
        csv_writer = csv.writer(file)
        csv_writer.writerows(data)  # Write all rows at once


def read_file():
    with open("example.csv", mode="r") as file:
        csv_reader = csv.reader(file)
        rows = []
        for row in csv_reader:
            rows.append(row)  # Each row is a list of strings
        return rows


if __name__ == "__main__":
    main()
