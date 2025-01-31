# Issues

## Pending

- (major) forgot to add ON CASCADE DELETE to assessments.lesson_id column

- (major) attempting to view lesson files returns a server error rather than a simple message stating that no files exist for the lesson
- (major) similarly, attempting to view lesson slides returns a server error

- 1/23/25 (minor) Deleting a course (unit, lesson may take a while as well) takes a long time. Maybe after a course row itself is deleted we should go ahead and return the response to the user and do the rest in the background. If child elements and files are not deleted we should log an error but the user doesn't need to wait on all of that, maybe?

## Complete

- 12/30/24 (minor) Lesson Edit form shows description as single-line with horizontal scroll rather than wrapping text. Resolved by switching from Fyne to web application (lol) 1/8/25
- 1/7/25 (major) Calendar doesn't fit to screen properly. Resolved by switching from Fyne to web application (lol) 1/8/25
- 1/4/25 (major) Shifting lessons across months doesn't show properly in the UI layer, probably because the map isn't updated or something. Going back out and then back into the calendar works though. Problem was I wasn't updating the in-memory Course upon shifting, I was only updating the UI and the DB. Resolved 1/4/24

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
