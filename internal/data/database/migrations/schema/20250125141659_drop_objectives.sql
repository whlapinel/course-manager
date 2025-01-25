-- +goose Up
-- +goose StatementBegin
DROP TABLE objectives;
DROP TABLE unit_standards;
ALTER TABLE lesson_objectives RENAME TO lesson_standards;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
