-- +goose Up
-- +goose StatementBegin
DROP TABLE standard_set;

CREATE TABLE IF NOT EXISTS standard_set (
id INTEGER PRIMARY KEY,
course_name TEXT NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
