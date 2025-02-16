-- +goose Up
-- +goose StatementBegin
DROP TABLE assessments;
CREATE TABLE IF NOT EXISTS assessments (
    id INTEGER PRIMARY KEY,
    lesson_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    instructions TEXT NOT NULL,
    category TEXT NOT NULL,
    date_assigned TEXT NOT NULL,
    date_due TEXT NOT NULL,
    dropped INTEGER NOT NULL,
    FOREIGN KEY (lesson_id) REFERENCES lessons(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
