

-- name: SaveAssessment :one
INSERT INTO assessments (
    lesson_id, name, instructions, category, date_assigned, date_due, dropped
) VALUES (
    ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetAssessmentByID :one
SELECT * FROM assessments WHERE id = ?;

-- name: GetAssessmentsByLessonID :many
SELECT * FROM assessments WHERE lesson_id = ?;

-- name: GetAssessmentsByDueDate :many
SELECT * FROM assessments WHERE date_due = ?;

-- name: GetAssessmentsByCategory :many
SELECT a.* FROM assessments a
JOIN lessons l ON a.lesson_id = l.id
JOIN units u ON l.unit_id = u.id
JOIN courses c ON u.course_id = c.id
WHERE c.id = ? AND a.category = ?;