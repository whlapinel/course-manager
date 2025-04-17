-- name: GetCourseByCourseID :one
SELECT * FROM courses c WHERE c.id = ?;

-- name: GetCoursesByTermID :many
SELECT * FROM courses c WHERE c.term_id = ?;

-- name: SaveCourse :one
INSERT INTO courses (
  term_id, name, description, std_set_id
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateCourse :exec
UPDATE courses 
SET name = ?, description = ?, std_set_id = ?
WHERE id = ?;

-- name: DeleteCourse :one
DELETE FROM courses WHERE id = ?
RETURNING *;

