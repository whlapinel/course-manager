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
  date = ?;

-- name: DeleteDate :exec
DELETE FROM dates
WHERE
  date = ?
  AND term_id = ?;

-- name: GetDateID :one
SELECT
  d.id
FROM
  dates d
WHERE
  d.date = ?;