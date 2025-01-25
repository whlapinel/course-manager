-- +goose Up
-- +goose StatementBegin

ALTER TABLE standards DROP COLUMN parent_id;
ALTER TABLE standards ADD COLUMN parent_id INTEGER REFERENCES standards(id) ON DELETE CASCADE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
