-- name: SaveFile :one
INSERT INTO files (
 name, description, base_path
) VALUES (
    ?, ?, ?
)
RETURNING *;

-- name: GetLessonFiles :many
SELECT * FROM files f
JOIN lesson_files lf ON lf.file_id = f.id
WHERE lf.lesson_id = ?;