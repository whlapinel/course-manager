-- +goose Up
-- +goose StatementBegin
ALTER TABLE assignments RENAME TO assessments;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
