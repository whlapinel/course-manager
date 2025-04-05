# Proposed Change: Change Assignment Parent ID from lesson to course

## Summary

This will make the lesson_id an optional field while making course_id mandatory, allowing more flexibility in the scope of an assignment.

## Justification

While most assessments aka "assignments" are specific to a topic introduced in a lesson, there are some, such as unit exams, or midterm/final exams where it does not make sense to associate it with a particular lesson. In fact, on exam days or project presentations, the assessment usually takes the entire class period and there is no "lesson" per se where new content is introduced. It is currently quite inconvenient to move an assessment from one lesson to another, actually it's impossible. The assessment must be deleted and then recreated. The assessment is tied strictly to its lesson_id and the file itself that is referenced in `Lesson.File` is stored within the lesson's files. While the lesson file issue may become moot due to [this](./assessment_files.md) change

## Impact

## Required actions

- Add `CourseID int` field to struct
- Add `UnitID int` field to struct
- Write migration script (probably needs to be in go vs. sql):
  - Remove `NOT NULL` constraint from `assessments.lesson_id` column
  - Add column `course_id INT NOT NULL` to `assessments` table
  - Add column `unit_id INT` to `assessments` table
