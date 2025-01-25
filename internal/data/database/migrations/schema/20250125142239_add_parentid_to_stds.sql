-- +goose Up
-- +goose StatementBegin
ALTER TABLE standards ADD COLUMN parent_id INTEGER REFERENCES standards(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
