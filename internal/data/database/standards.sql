-- name: SaveStandard :one
INSERT INTO standards (
  parent_id, set_id, number, name, description
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetStandardByID :one
SELECT * FROM standards where id = ?;

-- name: GetCourseStandards :many
SELECT * FROM standards WHERE set_id = ? AND parent_id IS NULL;

-- name: GetAllObjectives :many
SELECT * FROM standards WHERE set_id = ? AND parent_id IS NOT NULL;

-- name: GetStandardObjectives :many
SELECT * FROM standards WHERE parent_id = ?;

-- name: GetLessonStandards :many
SELECT s.* FROM standards s
JOIN lesson_standards ls 
ON ls.std_id = s.id AND ls.lesson_id = ?;
