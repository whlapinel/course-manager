-- name: GetTermOccasions :many
SELECT * FROM occasions WHERE term_id = ?;

-- name: GetOccasionByID :one
SELECT * FROM occasions WHERE id = ?;

-- name: SaveOccasion :one
INSERT INTO occasions (
    term_id, date, name
) VALUES (
    ?, ?, ?
)
RETURNING *;

-- name: UpdateOccasion :exec
UPDATE occasions SET name=? WHERE id=?;

-- name: DeleteOccasion :exec
DELETE FROM occasions WHERE id = ?;