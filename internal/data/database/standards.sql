-- name: SaveStandard :one
INSERT INTO standards (
  parent_id, set_id, number, name, description
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetCourseStandards :many
SELECT * FROM standards WHERE set_id = ?;

-- name: GetLessonStandards :many
SELECT * FROM standards s
JOIN lesson_standards ls 
ON ls.std_id = s.id AND ls.lesson_id = ?;
