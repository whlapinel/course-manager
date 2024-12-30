# Course Manager

## Progress Log

- 12/30/24 Today was a day of stepping back and realizing I have created a whole new mess of problems.
- 12/30/24 I am increasingly thinking that I should try to eliminate the central importance of the unit in the system.
  - A lesson should belong to a course rather than a unit
  - A course should have a slice of lessons rather than a slice of units
  - Parts of the course manager and course site that lists all units will also need to list lessons that don't have a unit
  - New queries to fetch lessons by unit id.
- 12/30/24 I realized today that allowing the user to change the name of a lesson would completely mess up the system, since the file system for the site generator depends on the names of lessons remaining constant across terms.
  - New hidden file system where directories match lessonIDs rather than lesson names.
  - User will no longer touch slide text directly but will rather edit this text indirectly either:
    1. Write to temporary file with file watcher and allow user to edit through another program
    2. Edit through the app, while (at some point) providing marp preview
  - Directory will be copied over for each course i.e. when a course is fit to a new term
  - Slide text will be stored in Database as text and added as a field of Lesson
  - Site generator will write the text to a markdown file before running marp to render to HTML
  - To save size requirements, images will be kept in a single directory so they don't need to be copied over for each course
  - If a lesson has a new ID
- 12/30/24 Worked on the edit lesson form a bit.
- 12/29/24 Major milestone today. Finally got the calendar interface working with the ability to move lessons one date left or right. Got rid of CourseSchedule across all layers.

## Features

### Pending (keep in priority order)

- 12/30/24 Go command to generate site should include other build commands e.g. tailwind, templ, etc (all the things that the task command does). Should probably also pass in the data instead of Generator using its own connection and fetching data?
- 12/30/24 Create Copy Course to Term interface (this will use the Course.FitToTerm method)
- 12/30/24 Create `New Lesson` UI
- 12/30/24 Create `New Unit` UI
- 12/30/24 Create New Course interface

### Complete

## Issues

### Pending

- 12/30/24 (minor) Lesson Edit form shows description as single-line with horizontal scroll rather than wrapping text.

### Complete

- 12/30/24 (major) ShowCalendar shows wrong course! Resolved 12/30/24

- 12/26/24 There's a serious problem with how the courses are saved, at least when imported from CSV. I noticed that the calendar for Spring 2025 was showing dates for January but none of the other term months. After some digging I realized that the problem in the "lesson dates" table and how lesson dates are saved.  Basically all I know is that the date_id column in the lesson_dates table consists only of Fall 2024 dates. Resolved 12/26 (new domain.Course method `FitToTerm(term Term) Course` which just sets the dates sequentially for each lesson to dates from the new term, with one date per lesson. After instructional days are exhausted, lesson dates are left empty. Lessons will be visible from course tree but not in course calendar.)

- 12/18/24 Resolved 12/27/24 (see proposal 2): figure out how to relate templates to instances. should the template
id only be present for a row if they are marked by user as synchronized? i.e. until the user marks sync?
  - Proposed solution 1 (rejected): whenever a user edits a course, unit or lesson that is an instance, i.e. where template is NOT NULL, the change should be also be made to the template row (and vice versa). If the user has elected to de-link the course, unit or lesson instance from the template, the change will not be reflected in the template object.
  - Proposed solution 2 (implemented): completely eliminate the entire concept and all implementations of course template as opposed to course instance.
    - There would only be a course, which will always be associated with a term. The fact that the dates table has a day number column will facilitate copying a course from one term to another.  This would be a very big change but the impact would be to dramatically simplify the code.
    - Would probably add the fields currently exclusively found in CourseInstance and move them to Course. Need to identify all parts of the code that depend on a distinction between the two, for example any use of templateID, and alter essential parts and eliminate non-essential ones.
    - Make sure the Database tables reflect this change completely.
- 12/19/24 calendar looks better now
- 12/17/24 lesson struct should have Date field be Dates []time.Time instead of Date time.Time so that a lesson can span multiple dates if necessary (not ideal but sometimes necessary) (completed 12/17/24)
