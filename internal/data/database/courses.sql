-- name: GetCourseByCourseID :one
SELECT * FROM courses c WHERE c.id = ?;

-- name: GetCourses :many
SELECT
  c.id as course_id,
  c.name as course_name,
  c.description as course_descr
FROM 
  courses c
WHERE 
  c.term_id = ?;

-- name: GetCourse :one
SELECT
  c.id as course_id,
  c.name as course_name,
  c.description as course_descr,
  c.term_id
FROM
  courses c
WHERE
  c.term_id = ?;

-- name: GetCourseByID :one
SELECT * from courses WHERE id = ?;

-- name: SaveCourse :one
INSERT INTO courses (
  term_id, name, description
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateCourse :exec
UPDATE courses 
SET name = ?, description = ?
WHERE id = ?;

-- name: DeleteCourse :one
DELETE FROM courses WHERE id = ?
RETURNING *;

