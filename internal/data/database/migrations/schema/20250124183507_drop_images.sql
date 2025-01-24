-- +goose Up
-- +goose StatementBegin
DROP TABLE course_images;
DROP TABLE unit_images;
DROP TABLE lesson_images;
DROP TABLE images;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
