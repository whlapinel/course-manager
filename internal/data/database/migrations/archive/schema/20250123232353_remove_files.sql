-- +goose Up
-- +goose StatementBegin
DROP TABLE lesson_files;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
