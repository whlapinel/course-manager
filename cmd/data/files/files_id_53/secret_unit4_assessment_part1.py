# Main function to execute the quiz
def main():
    # Call each function to perform the respective tasks
    write_name_to_file("students.txt", "Alice")
    write_scores_to_file("scores.txt", [85, 90, 78])
    append_grade_to_file("scores.txt", 92)
    total = read_and_calculate_total("scores.txt")
    average = read_and_calculate_average("scores.txt")

    # Output the results
    print(f"Total score: {total}")
    print(f"Average score: {average}")


# 1. Function to write a student's name to a file
def write_name_to_file(filename: str, name: str) -> None:
    """
    Write the student's name to the specified file.
    Hint: Use 'w' mode to overwrite any existing content.
    """
    raise NotImplementedError


# 2. Function to write a list of scores to a file
def write_scores_to_file(filename: str, scores: list) -> None:
    """
    Write a list of scores to the specified file.
    Hint: Use a loop to iterate through the scores and write them one by one.
    """
    raise NotImplementedError


# 3. Function to append a new grade to the scores file
def append_grade_to_file(filename: str, grade: int) -> None:
    """
    Append a new grade to the existing scores file.
    Hint: Use 'a' mode to append without overwriting.
    """
    raise NotImplementedError


# 4. Function to read the scores from the file and calculate the total
def read_and_calculate_total(filename: str) -> int:
    """
    Read the scores from the file and calculate the total.
    Hint: Use a loop to sum the values after reading the file.
    """
    raise NotImplementedError


# 5. Function to read the scores and calculate the average
def read_and_calculate_average(filename: str) -> float:
    """
    Read the scores from the file and calculate the average.
    Hint: Keep track of both the sum and the number of scores.
    """
    raise NotImplementedError


# Call the main function to run the quiz
if __name__ == "__main__":
    main()

# CODE BANK
# with open("<file_name>", "<mode>") as file: # w for write (replaces all content), r for read, a for append (adds to end of file rather than replacing)
#   file.write("<stuff>") # for writing
#   content = file.read() # for reading
