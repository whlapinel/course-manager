-- name: SaveTerm :one
INSERT INTO terms (
  name, description, start, end
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateTerm :exec
UPDATE terms SET name = ?, description = ?, start = ?, end = ?
WHERE id = ?;

-- name: DeleteTerm :one
DELETE FROM terms WHERE id = ?
RETURNING *;

-- name: GetTerm :one
SELECT
  t.id,
  t.name,
  t.description,
  t.start,
  t.end,
  d.date
FROM
  terms t
LEFT JOIN
  dates d ON d.term_id = t.id
WHERE
  d.date = ?;

-- name: GetTermByID :one
SELECT * FROM terms WHERE id = ?;

-- name: GetTerms :many
SELECT   
  t.id,
  t.name,
  t.description,
  t.start,
  t.end,
  d.date 
FROM terms t
LEFT JOIN dates d 
ON d.term_id = t.id
ORDER BY t.id, d.day_number;

-- name: GetTermDates :many
SELECT * FROM dates d
WHERE d.term_id = ?;

-- name: DeleteNonInstructDays :one
DELETE FROM non_instruct_days WHERE id = ?
RETURNING *;