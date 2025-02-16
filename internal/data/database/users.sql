-- name: SaveUser :one
INSERT INTO
    users (id, first_name, last_name, email, picture)
VALUES
    (?, ?, ?, ?, ?) RETURNING *;

-- name: GetUser :one
SELECT
    *
FROM
    users
WHERE
    id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE
    id = ?;

-- name: UpdateUser :exec
UPDATE users
SET
    first_name = ?,
    last_name = ?,
    email = ?,
    picture = ?;