-- name: SaveFile :one
INSERT INTO files (
 name, description, file_name, modified
) VALUES (
    ?, ?, ?, ?
)
RETURNING *;

