-- name: GetCourseOccasions :many
SELECT
    *
FROM
    course_occasions
WHERE
    course_id = ?;

-- name: GetCourseOccasionByID :one
SELECT
    *
FROM
    course_occasions
WHERE
    id = ?;

-- name: SaveCourseOccasion :one
INSERT INTO
    course_occasions (course_id, date, name)
VALUES
    (?, ?, ?) RETURNING *;

-- name: UpdateCourseOccasion :exec
UPDATE course_occasions
SET
    name = ?
WHERE
    id = ?;

-- name: DeleteCourseOccasion :exec
DELETE FROM course_occasions
WHERE
    id = ?;