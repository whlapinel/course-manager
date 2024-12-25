-- name: GetInstances :many
SELECT
  c.id as course_id,
  c.name as course_name,
  c.description as course_descr
FROM 
  courses c
WHERE 
  c.term_id = ?;

-- name: GetInstance :one
SELECT
  c.id as course_id,
  c.name as course_name,
  c.description as course_descr,
  c.term_id
FROM
  courses c
WHERE
  c.term_id = ? AND c.template_id = ?;

-- name: GetInstanceByID :one
SELECT * from courses WHERE id = ?;

-- name: SaveInstance :one
INSERT INTO courses (
  template_id, term_id, name, description
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateInstance :exec
UPDATE courses 
SET name = ?, description = ?
WHERE id = ?;

-- name: DeleteInstance :one
DELETE FROM courses WHERE id = ?
RETURNING *;

