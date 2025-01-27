-- +goose Up
-- +goose StatementBegin
DROP TABLE lesson_standards;
CREATE TABLE IF NOT EXISTS lesson_standards (
    id INTEGER PRIMARY KEY,
    std_id INTEGER NOT NULL,
    lesson_id INTEGER NOT NULL,
    FOREIGN KEY (std_id) REFERENCES standards(id) ON DELETE CASCADE,
    FOREIGN KEY (lesson_id) REFERENCES lessons(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
