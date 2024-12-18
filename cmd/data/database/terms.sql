
-- name: SaveTerm :one
INSERT INTO terms (
  name, start, end
) VALUES (
  ?, ?, ?
)
RETURNING *;



-- name: DeleteTerm :one
DELETE FROM terms WHERE id = ?
RETURNING *;



-- name: GetTerm :one
SELECT
  t.id,
  t.name
FROM
  terms t
JOIN
  dates d ON d.term_id = t.id
WHERE
  d.date = ?;


-- name: DeleteNonInstructDays :one
DELETE FROM non_instruct_days WHERE id = ?
RETURNING *;
