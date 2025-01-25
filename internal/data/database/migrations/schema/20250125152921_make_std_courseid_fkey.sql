-- +goose Up
-- +goose StatementBegin
DROP TABLE standards;
CREATE TABLE IF NOT EXISTS standards (
    id INTEGER PRIMARY KEY,
    number INTEGER NOT NULL,
    name TEXT NOT NULL, description TEXT, 
    set_id INTEGER NOT NULL REFERENCES standard_set(id) ON DELETE CASCADE,
    parent_id INTEGER REFERENCES standards(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
