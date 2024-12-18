-- name: GetCourses :many
SELECT * FROM courses;

-- name: GetInstances :many
SELECT
  c.id as course_id,
  c.name as course_name,
  c.description as course_descr,
  c.term_id
FROM 
  courses c
WHERE 
  c.template_id IS NOT NULL;

-- name: GetInstance :one
SELECT
  c.id as course_id,
  c.template_id as template_id,
  c.name as course_name,
  c.description as course_descr,
  c.term_id
FROM
  courses c
WHERE
  c.id = ?;

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

-- name: GetTemplate :one
SELECT
  c.id as course_id,
  c.name as course_name,
  c.description as course_descr
FROM 
  courses c
WHERE 
  c.id = ?;


-- name: GetTemplates :many
SELECT
  c.id as course_id,
  c.name as course_name,
  c.description as course_descr
FROM 
  courses c
WHERE 
  c.template_id IS NULL;
  
-- name: GetUnits :many
SELECT
  u.id,
  u.course_id,
  u.template_id,
  u.number,
  u.sequence,
  u.name,
  u.description
FROM
  units u
WHERE
  u.course_id = ?
ORDER BY
  u.sequence;

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


-- name: GetDate :one
SELECT * FROM dates WHERE date = ?;

-- name: SaveLessonDate :one
INSERT INTO lesson_dates (
  lesson_id, date_id
) VALUES (
  ?, ?
)
RETURNING *;

-- name: SaveCourse :one
INSERT INTO courses (
  template_id, term_id, name, description
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;


-- name: SaveUnit :one
INSERT INTO units (
  number, sequence, name, description, course_id, template_id
) VALUES (
  ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: SaveLesson :one
INSERT INTO lessons (
  number, name, description, unit_id, template_id
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;

-- name: SaveTerm :one
INSERT INTO terms (
  name, start, end
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: SaveDate :one
INSERT INTO dates (
  term_id, day_number, date
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: DeleteLesson :one
DELETE FROM lessons WHERE id = ?
RETURNING *;

-- name: DeleteUnit :one
DELETE FROM units WHERE id = ?
RETURNING *;

-- name: DeleteTerm :one
DELETE FROM terms WHERE id = ?
RETURNING *;

-- name: DeleteNonInstructDays :one
DELETE FROM non_instruct_days WHERE id = ?
RETURNING *;

-- name: DeleteCourse :one
DELETE FROM courses WHERE id = ?
RETURNING *;






