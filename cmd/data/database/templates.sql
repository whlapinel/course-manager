
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
  
-- name: SaveTemplate :one
INSERT INTO courses (
  name, description
) VALUES (
  ?, ?
)
RETURNING *;