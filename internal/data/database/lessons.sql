-- name: DeleteLesson :exec
DELETE FROM lessons WHERE id = ?;

-- name: SaveLesson :one
INSERT INTO lessons (
  number, name, description, unit_id
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: GetLesson :one
SELECT 
  l.*
FROM 
  lessons l
WHERE 
  l.id = ?;

-- name: GetLessons :many
SELECT
  l.id,
  l.number,
  l.name,
  l.description
FROM
  lessons l
WHERE
  l.unit_id = ?;

-- name: GetLessonDates :many
SELECT
  d.date
FROM
  lesson_dates ld
JOIN
  dates d ON d.id = ld.date_id
WHERE
  lesson_id = ?;

-- name: GetDateID :one
SELECT
    d.id
FROM
    dates d
WHERE
    d.date = ?;

-- name: SaveLessonDate :one
INSERT INTO lesson_dates (
  lesson_id, date_id
) VALUES (
  ?, ?
)
RETURNING *;

-- name: GetLessonsOnDate :many
SELECT ld.*, l.* from lesson_dates ld
JOIN dates d ON d.id = ld.date_id
JOIN lessons l ON l.id = ld.lesson_id
WHERE d.date = ? and d.term_id = ?;

-- name: UpdateLesson :exec
UPDATE lessons
set name = ?, number = ?, description = ?
WHERE id = ?;

-- name: DeleteLessonDate :exec
DELETE FROM lesson_dates
WHERE lesson_id = ? AND date_id = ?;

-- name: DeleteAllLessonDates :exec
DELETE FROM lesson_dates
WHERE lesson_id = ?;
