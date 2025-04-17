-- name: SaveStandardSet :one
INSERT INTO standard_sets (
    course_name
) VALUES (
    ?
) RETURNING *;

-- name: GetAllStandardSets :many
SELECT * FROM standard_sets;

-- name: GetStdSetByCourseName :one
SELECT * FROM standard_sets WHERE course_name = ?;

-- name: GetStdSetByID :one
SELECT * FROM standard_sets WHERE id = ?;

