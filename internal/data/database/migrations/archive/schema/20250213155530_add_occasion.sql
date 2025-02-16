-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS occasions (
    id INTEGER PRIMARY KEY,
    term_id INTEGER NOT NULL REFERENCES terms(id) ON DELETE CASCADE,
    date TEXT NOT NULL,
    name TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE occasions;
-- +goose StatementEnd
