-- +goose Up
-- +goose StatementBegin
ALTER TABLE assessments ADD COLUMN file TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
