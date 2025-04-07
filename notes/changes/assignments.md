# Proposed Change: Change Assignment Parent ID from lesson to course

## Summary

This will make the lesson_id an optional field while making course_id mandatory, allowing more flexibility in the scope of an assignment. This will not impact functionality but will help prepare for additional functionality.

## Justification

While most assessments aka "assignments" are specific to a topic introduced in a lesson, there are some, such as unit exams, or midterm/final exams where it does not make sense to associate it with a particular lesson. In fact, on exam days or project presentations, the assessment usually takes the entire class period and there is no "lesson" per se where new content is introduced. It is currently quite inconvenient to move an assessment from one lesson to another; actually, it's impossible. The assessment must be deleted and then recreated. The assessment is tied strictly to its lesson_id and the file itself that is referenced in `Lesson.File` is stored within the lesson's files. Note: the lesson file issue may become moot due to [this](./assessment_files.md) change which will facilitate moving assessments up and down the tree as needed. I am not sure how frequently this would need to happen but what I am trying to avoid is a user having to delete and recreate an assessment rather than simply disassociating with the current node and associating it with another one, the case that they have a lesson assessment and want to make it a unit assessment, or a course assessment and want to make it a unit assessment, etc.

## Impact

- a few hours of time

## Required actions

- [X] Add `CourseID int` field to struct
- [X] Add `UnitID int` field to struct
- [X] Write migration script (probably needs to be in go vs. sql):
  - [X] Remove `NOT NULL` constraint from `assessments.lesson_id` column
  - [X] Add column `course_id INT NOT NULL` to `assessments` table
  - [X] Add column `unit_id INT` to `assessments` table
- [ ] *Eventually* need to add functionality for moving an assessment from one node to another
