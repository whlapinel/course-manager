-- name: GetTermOccasions :many
SELECT
    *
FROM
    term_occasions
WHERE
    term_id = ?;

-- name: GetTermOccasionByID :one
SELECT
    *
FROM
    term_occasions
WHERE
    id = ?;

-- name: SaveTermOccasion :one
INSERT INTO
    term_occasions (term_id, date, name)
VALUES
    (?, ?, ?) RETURNING *;

-- name: UpdateTermOccasion :exec
UPDATE term_occasions
SET
    name = ?
WHERE
    id = ?;

-- name: DeleteTermOccasion :exec
DELETE FROM term_occasions
WHERE
    id = ?;