# =========================================================
#  Simple Grade Tracker
#  (Working version)
# =========================================================

FILENAME = "grades.txt"
VALID_GRADES = ("A", "B", "C", "D", "F")  # A tuple of valid letter grades


def load_grades():
    """
    Reads students and grades from a file and returns a list of dictionaries.
    Each dictionary has keys: 'name' and 'grade'.
    """
    grades_list = []
    try:
        with open(FILENAME, "r") as file:
            for line in file:
                line = line.strip()
                if line:
                    # Format: "student_name,grade"
                    parts = line.split(",")
                    if len(parts) == 2:
                        student_name = parts[0]
                        grade = parts[1]
                        # Validate the grade
                        if grade not in VALID_GRADES:
                            grade = "F"  # default if invalid
                        grades_list.append({"name": student_name, "grade": grade})
    except FileNotFoundError:
        # If the file doesn't exist, just start with an empty list
        pass
    return grades_list


def save_grades(grades_list):
    """
    Saves the student/grade data to the FILENAME.
    """
    with open(FILENAME, "w") as file:
        for record in grades_list:
            line = f"{record['name']},{record['grade']}\n"
            file.write(line)


def display_grades(grades_list):
    """
    Prints all the student records in a formatted way.
    """
    if not grades_list:
        print("No records to display.")
        return
    print("\n--- Grade Records ---")
    for idx, record in enumerate(grades_list, start=1):
        print(f"{idx}. {record['name']} - Grade: {record['grade']}")
    print("--------------------\n")


def add_grade(grades_list):
    """
    Prompts the user to enter a new student's name and grade,
    then adds it to the list.
    """
    name = input("Enter student name: ").strip()
    if not name:
        print("Student name cannot be empty.\n")
        return

    grade = (
        input(f"Enter letter grade ({'/'.join(VALID_GRADES)}) [default='F']: ")
        .strip()
        .upper()
    )
    if grade not in VALID_GRADES:
        grade = "F"

    record = {"name": name, "grade": grade}
    grades_list.append(record)
    print(f"Added record: {name} - Grade: {grade}\n")


def remove_grade(grades_list):
    """
    Displays the current records and asks for a record number to remove.
    """
    if not grades_list:
        print("No records to remove.\n")
        return

    display_grades(grades_list)
    try:
        record_num = int(input("Enter the record number to remove: "))
        index = record_num - 1
        if index < 0 or index >= len(grades_list):
            print("Invalid record number.\n")
            return
        removed_record = grades_list.pop(index)
        print(f"Removed: {removed_record['name']} - Grade: {removed_record['grade']}\n")
    except ValueError:
        print("Invalid input. Please enter a number.\n")


def main():
    """
    The main menu loop for the application.
    """
    grades_list = load_grades()
    print("Welcome to the Simple Grade Tracker!\n")

    while True:
        print("Menu:")
        print("1. Display Grades")
        print("2. Add Grade")
        print("3. Remove Grade")
        print("4. Save Grades")
        print("5. Exit\n")

        choice = input("Enter your choice (1-5): ").strip()

        if choice == "1":
            display_grades(grades_list)
        elif choice == "2":
            add_grade(grades_list)
        elif choice == "3":
            remove_grade(grades_list)
        elif choice == "4":
            save_grades(grades_list)
            print("Grades saved successfully.\n")
        elif choice == "5":
            print("Goodbye!")
            break
        else:
            print("Invalid choice. Please try again.\n")


if __name__ == "__main__":
    main()
