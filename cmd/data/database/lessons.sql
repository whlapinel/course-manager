

-- name: DeleteLesson :exec
DELETE FROM lessons WHERE id = ?;

-- name: SaveLesson :one
INSERT INTO lessons (
  number, name, description, unit_id, template_id
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;


-- name: GetLessons :many
SELECT
  l.id,
  l.template_id,
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

-- name: UpdateLesson :exec
UPDATE lessons
set name = ?, number = ?, description = ?
WHERE id = ?;

-- name: DeleteLessonDates :exec
DELETE FROM lesson_dates
WHERE lesson_id = ? and date_id = ?;
