# Course Manager

## PENDING

- 12/18/24 figure out how to relate templates to instances. should the template
id only be present for a row if they are marked by user as synchronized? i.e. until the user marks sync?
  - Proposed solution 1: whenever a user edits a course, unit or lesson that is an instance, i.e. where template is NOT NULL, the change should be also be made to the template row (and vice versa). If the user has elected to de-link the course, unit or lesson instance from the template, the change will not be reflected in the template object.
  - Proposed solution 2: completely eliminate the entire concept and all implementations of course template as opposed to course instance.
    - There would only be a course, which will always be associated with a term. The fact that the dates table has a day number column will facilitate copying a course from one term to another.  This would be a very big change but the impact would be to dramatically simplify the code.
    - Would probably add the fields currently exclusively found in CourseInstance and move them to Course. Need to identify all parts of the code that depend on a distinction between the two, for example any use of templateID, and alter essential parts and eliminate non-essential ones.
    - Make sure the Database tables reflect this change completely.

## COMPLETE

- 12/19/24 calendar looks better now
- 12/17/24 lesson struct should have Date field be Dates []time.Time instead of
Date time.Time so that a lesson can span multiple dates if necessary (not ideal
but sometimes necessary) (completed 12/17/24)