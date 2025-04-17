-- name: SaveStandard :one
INSERT INTO standards (
  parent_id, set_id, number, name, description
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetStandardByID :one
SELECT o.*, s.number as parent_num FROM standards o
JOIN standards s ON s.id = o.parent_id
WHERE o.id = ?;

-- name: GetCourseStandards :many
SELECT * FROM standards WHERE set_id = ? AND parent_id IS NULL;

-- name: GetAllObjectives :many
SELECT * FROM standards WHERE set_id = ? AND parent_id IS NOT NULL;

-- name: GetStandardObjectives :many
SELECT o.*, s.number as parent_num
FROM standards o
JOIN standards s ON s.id = o.parent_id 
WHERE o.parent_id = ?;

-- name: GetLessonStandards :many
SELECT o.*, s.number as parent_num FROM standards o
JOIN standards s ON s.id = o.parent_id 
JOIN lesson_standards ls 
ON ls.std_id = o.id AND ls.lesson_id = ?;
