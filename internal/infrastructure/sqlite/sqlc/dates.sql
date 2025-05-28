-- name: SaveDate :one
INSERT INTO
  dates (term_id, date)
VALUES
  (?, ?) RETURNING *;

-- name: GetDate :one
SELECT
  *
FROM
  dates
WHERE
  date = ?
  AND term_id = ?;

-- name: DeleteDate :exec
DELETE FROM dates
WHERE
  date = ?
  AND term_id = ?;