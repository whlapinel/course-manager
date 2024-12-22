-- name: GetUnits :many
SELECT
  u.id,
  u.course_id,
  u.template_id,
  u.number,
  u.sequence,
  u.name,
  u.description
FROM
  units u
WHERE
  u.course_id = ?
ORDER BY
  u.sequence;

-- name: SaveUnit :one
INSERT INTO units (
  number, sequence, name, description, course_id, template_id
) VALUES (
  ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateUnit :exec
UPDATE units
SET number = ?, sequence = ?, name = ?, description = ?
WHERE id = ?;

-- name: DeleteUnit :one
DELETE FROM units WHERE id = ?
RETURNING *;
