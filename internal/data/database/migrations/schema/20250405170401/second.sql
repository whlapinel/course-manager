UPDATE assessments
SET
    course_id = (
        SELECT
            u.course_id
        FROM
            lessons l
            JOIN units u ON l.unit_id = u.id
        WHERE
            l.id = assessments.lesson_id
    );