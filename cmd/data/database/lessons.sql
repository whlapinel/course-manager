-- name: DeleteLesson :exec
DELETE FROM lessons WHERE id = ?;

-- name: SaveLesson :one
INSERT INTO lessons (
  number, name, description, unit_id
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;


-- name: GetLessons :many
SELECT
  l.id,
  l.number,
  l.name,
  l.description,
  f.name as file_name,
  f.description as file_descr,
  f.file_name,
  f.modified as file_modified
FROM
  lessons l
LEFT JOIN
  lesson_files lf ON lf.lesson_id = l.id
JOIN
  files f ON f.id = lf.file_id 
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

-- name: DeleteLessonDate :exec
DELETE FROM lesson_dates
WHERE lesson_id = ? AND date_id = ?;

-- name: DeleteAllLessonDates :exec
DELETE FROM lesson_dates
WHERE lesson_id = ?;
