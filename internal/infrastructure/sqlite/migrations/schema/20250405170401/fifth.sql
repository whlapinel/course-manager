INSERT INTO
    assessments (
        id,
        course_id,
        lesson_id,
        name,
        instructions,
        file,
        category,
        date_assigned,
        date_due,
        dropped
    )
SELECT
    id,
    course_id,
    lesson_id,
    name,
    instructions,
    file,
    category,
    date_assigned,
    date_due,
    dropped
FROM
    assessments_old