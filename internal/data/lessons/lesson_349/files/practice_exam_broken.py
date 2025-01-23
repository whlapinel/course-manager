# =========================================================
#  Simple Grade Tracker (Broken in 10 places)
# =========================================================

# 1) Filename mismatch?
FILENAM = "grades.txt"  # Bug #1: Perhaps we need to be consistent with usage

# 2) Are these valid grades or is the variable name correct?
VALID_GRADE = ("A", "B", "C", "D", "F")

def load_grades():
    grade_records = []
    try:
        # 3) Mismatch in variable name again? (FILENAM vs FILENAME?)
        with open(FILENAME, "r") as file:
            for line in file:
                line.strip()  # Bug #2: .strip() returns a new string, not in-place
                if line:
                    parts = line.split(",")
                    if len(parts) == 2:
                        student_name = parts[0]
                        letter_grade = parts[1]
                        # 4) If the grade isn't valid, do we fix it?
                        if letter_grade not in VALID_GRADES:  # Bug #3: variable mismatch: VALID_GRADE vs VALID_GRADES
                            letter_grade = "F"
                        grade_records.append({
                            "name": student_name,
                            "grade": letter_grade
                        })
    except FileNotFoundError:
        pass
    return grade_records

def save_grades(grades_list):
    with open(FILENAM, "w") as file:
        for rec in grades_list:
            # 5) Is the f-string correct?
            line = f"{rec['name']}, {rec'name']}\n"  # Bug #4: wrong dictionary key usage
            file.write(line)

def display_grades(grade_list):
    # 6) Checking for an empty list incorrectly?
    if grade_list is None:
        print("No records to display.")
        return
    print("\n--- Grade Records ---")
    for i, record in enumerate(grade_list, start=1):
        # 7) Mismatch in keys: is it record['student'] or record['name']?
        print(f"{i}. {record['student']} - Grade: {record['letter']}")  # Bug #5: might not match the actual keys used
    print("---------------------\n")

def add_grade(records):
    name = input("Enter student name: ")
    if name == "":
        print("Student name cannot be empty.")
        return

    # 8) Are we comparing to the correct tuple for valid grades?
    chosen_grade = input("Enter letter grade (A/B/C/D/F) [default='F']: ").strip().upper()
    if chosen_grade not in VALID_GRADES_TUPLE:  # Bug #6: variable mismatch
        chosen_grade = "F"
    
    # 9) Wrong dictionary keys?
    new_record = {
        "student_name": name,
        "grade_letter": chosen_grade
    }
    records.append(new_record)
    print("Grade added successfully!")

def remove_grade(rec_list):
    display_grades(rec_list)
    try:
        # 10) Should we parse an integer for the index or a float?
        idx = float(input("Enter the record number to remove: "))  # Bug #7: float or int?
        index = idx - 1
        removed = rec_list.pop(index)
        print(f"Removed: {removed['name']} - Grade: {removed['grade']}")
    except (ValueError, IndexError):
        print("Invalid record number.")

def main():
    # 11) Another mismatch or missing colon in while?
    loaded_records = load_grade()
    print("Welcome to the Grade Tracker!")

    while True  # Bug #8: missing colon
        print("1. Display Grades")
        print("2. Add Grade")
        print("3. Remove Grade")
        print("4. Save Grades")
        print("5. Exit")
        choice = input("Choose an option (1-5): ")

        if choice == "1":
            display_grades(loaded_records)
        elif choice == "2":
            add_grade(loaded_records)
        elif choice == "3":
            remove_grade(loaded_records)
        elif choice == "4":
            save_grades(loaded_records)
            print("Grades saved.")
        elif choice == "5":
            print("Goodbye!")
            break
        else:
            print("Invalid choice. Try again.")

if __name__ == "__main__":
    main()
   print("Finished!")  # Bug #9: indentation error?
