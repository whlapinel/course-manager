-- name: SaveFilesDir :one
INSERT INTO files_dir (
 name, description
) VALUES (
    ?, ?
)
RETURNING *;

-- name: GetLessonFilesDir :one
SELECT f.* FROM files_dir f
JOIN lesson_files lf ON lf.file_id = f.id
WHERE lf.lesson_id = ?;