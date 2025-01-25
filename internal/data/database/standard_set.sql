-- name: SaveStandardSet :one
INSERT INTO standard_set (
    course_name
) VALUES (
    ?
) RETURNING *;

-- name GetAllStandardSets :many
SELECT * FROM standard_set;

-- name GetStdSetByCourseName :one
SELECT * FROM standard_set WHERE course_name = ?;

-- name GetStdSetByID :one
SELECT * FROM standard_set WHERE id = ?;

