-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS standard_set (
id INTEGER PRIMARY KEY,
course_name TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
