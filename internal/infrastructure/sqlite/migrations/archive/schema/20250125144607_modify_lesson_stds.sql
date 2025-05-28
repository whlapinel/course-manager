-- +goose Up
-- +goose StatementBegin
ALTER TABLE lesson_standards RENAME COLUMN obj_id TO std_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
