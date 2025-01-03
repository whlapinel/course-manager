-- name: SaveSlides :one
INSERT INTO slides (
 name, description
) VALUES (
    ?, ?
)
RETURNING *;

-- name: GetSlides :one
SELECT s.* FROM slides s
JOIN lesson_slides ls ON ls.slides_id = s.id
WHERE ls.lesson_id = ?;

-- name: SaveLessonSlides :one
INSERT INTO lesson_slides (
  slides_id, lesson_id
) VALUES (
  ?, ?
)
RETURNING *;