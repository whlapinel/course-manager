-- name: SaveDate :one
INSERT INTO dates (
  term_id, day_number, date
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: GetDate :one
SELECT * FROM dates WHERE date = ?;

