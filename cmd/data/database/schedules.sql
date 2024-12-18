-- name: GetDailySchedules :many
SELECT
  d.date,
  d.day_number,
  l.name as lesson_name,
  u.name as unit_name,
  l.description as lesson_description
FROM
  dates d
JOIN
  lesson_dates ld ON ld.date_id = d.id
JOIN
  lessons l ON l.id = ld.lesson_id
JOIN
  units u ON u.id = l.unit_id
JOIN
  courses c ON c.id = u.course_id
WHERE
  c.id = ?
ORDER BY
  d.day_number;