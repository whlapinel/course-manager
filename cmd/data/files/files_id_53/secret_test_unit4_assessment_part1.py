import os
import pytest
from unit4.unit4_assessment_part1_solution import (
    write_name_to_file,
    write_scores_to_file,
    append_grade_to_file,
    read_and_calculate_total,
    read_and_calculate_average,
)

# Test file names
name_file = "test_students.txt"
scores_file = "test_scores.txt"


# Clean up the files before running any tests
@pytest.fixture(autouse=True)
def cleanup():
    # Remove the test files if they exist
    yield
    if os.path.exists(name_file):
        os.remove(name_file)
    if os.path.exists(scores_file):
        os.remove(scores_file)


def test_write_name_to_file():
    # Test writing a name to the file
    write_name_to_file(name_file, "TestStudent")
    assert os.path.exists(name_file), "File should exist after writing"

    with open(name_file, "r") as file:
        content = file.read().strip()
        assert content == "TestStudent", f"Expected 'TestStudent', got '{content}'"


def test_write_scores_to_file():
    # Test writing a list of scores to a file
    scores = [85, 90, 78]
    write_scores_to_file(scores_file, scores)

    with open(scores_file, "r") as file:
        content = file.readlines()
        content = [int(line.strip()) for line in content]
        assert content == scores, f"Expected {scores}, got {content}"


def test_append_grade_to_file():
    # Test appending a grade to the scores file
    scores = [85, 90, 78]
    write_scores_to_file(scores_file, scores)
    append_grade_to_file(scores_file, 92)

    with open(scores_file, "r") as file:
        content = file.readlines()
        content = [int(line.strip()) for line in content]
        assert content == scores + [92], f"Expected {scores + [92]}, got {content}"


def test_read_and_calculate_total():
    # Test calculating the total of the scores
    scores = [85, 90, 78]
    write_scores_to_file(scores_file, scores)
    append_grade_to_file(scores_file, 92)

    total = read_and_calculate_total(scores_file)
    assert total == sum(scores) + 92, f"Expected total {sum(scores) + 92}, got {total}"


def test_read_and_calculate_average():
    # Test calculating the average of the scores
    scores = [85, 90, 78]
    write_scores_to_file(scores_file, scores)
    append_grade_to_file(scores_file, 92)

    average = read_and_calculate_average(scores_file)
    expected_average = (sum(scores) + 92) / len(scores + [92])
    assert (
        average == expected_average
    ), f"Expected average {expected_average}, got {average}"
