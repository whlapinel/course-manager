-- name: DeleteLesson :exec
DELETE FROM lessons
WHERE
  id = ?;

-- name: SaveLesson :one
INSERT INTO
  lessons (number, name, description, unit_id)
VALUES
  (?, ?, ?, ?) RETURNING *;

-- name: GetLesson :one
SELECT
  l.*,
  u.number AS unit_number,
  u.name AS unit_name
FROM
  lessons l
  JOIN units u ON l.unit_id = u.id
WHERE
  l.id = ?;

-- name: GetLessons :many
SELECT
  l.id,
  l.number,
  l.name,
  l.description,
  u.number AS unit_number,
  u.name AS unit_name
FROM
  lessons l
  JOIN units u ON l.unit_id = u.id
WHERE
  l.unit_id = ?;


-- name: GetLessonDates :many
SELECT
  d.date
FROM
  lesson_dates ld
  JOIN dates d ON d.id = ld.date_id
WHERE
  lesson_id = ?;

-- name: SaveLessonDate :one
INSERT INTO
  lesson_dates (lesson_id, date_id)
VALUES
  (?, ?) RETURNING *;

-- name: GetLessonsOnDateForCourse :many
SELECT
  ld.*,
  l.*,
  u.number as unit_number,
  u.name as unit_name
from
  lesson_dates ld
  JOIN dates d ON d.id = ld.date_id
  JOIN lessons l ON l.id = ld.lesson_id
  JOIN units u ON u.id = l.unit_id
  JOIN courses c ON c.id = u.course_id
  JOIN terms t ON t.id = c.term_id
WHERE
  d.date = ?
  and c.id = ?;

-- name: GetLessonsOnDate :many
SELECT
  ld.*,
  l.*
from
  lesson_dates ld
  JOIN dates d ON d.id = ld.date_id
  JOIN lessons l ON l.id = ld.lesson_id
WHERE
  d.date = ?
  and d.term_id = ?;

-- name: UpdateLesson :exec
UPDATE lessons
set
  name = ?,
  number = ?,
  description = ?
WHERE
  id = ?;

-- name: DeleteLessonDate :exec
DELETE FROM lesson_dates
WHERE
  lesson_id = ?
  AND date_id = ?;

-- name: DeleteAllLessonDates :exec
DELETE FROM lesson_dates
WHERE
  lesson_id = ?;

-- name: SaveLessonStandard :exec
INSERT INTO
  lesson_standards (std_id, lesson_id)
VALUES
  (?, ?);

-- name: DeleteLessonStandard :exec
DELETE FROM lesson_standards
WHERE
  std_id = ?
  AND lesson_id = ?;