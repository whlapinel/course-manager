-- name: SaveImage :one
INSERT INTO images (
 name, description, base_path
) VALUES (
    ?, ?, ?
)
RETURNING *;

-- name: SaveCourseImage :one
INSERT INTO course_images (
    course_id, image_id
) VALUES (
    ?, ?
)
RETURNING *;

-- name: SaveUnitImage :one
INSERT INTO unit_images (
    unit_id, image_id
) VALUES (
    ?, ?
)
RETURNING *;

-- name: SaveLessonImage :one
INSERT INTO lesson_images (
    lesson_id, image_id
) VALUES (
    ?, ?
)
RETURNING *;

-- name: GetCourseImage :one
SELECT i.* FROM images i 
JOIN course_images ci ON ci.image_id = i.id
WHERE ci.course_id = ?;

-- name: GetUnitImage :one
SELECT i.* FROM images i 
JOIN unit_images ui ON ui.image_id = i.id
WHERE ui.unit_id = ?;

-- name: GetLessonImage :one
SELECT i.* FROM images i 
JOIN lesson_images li ON li.image_id = i.id
WHERE li.lesson_id = ?;