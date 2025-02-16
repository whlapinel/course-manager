-- +goose Up
-- +goose StatementBegin
ALTER TABLE standard_set RENAME TO standard_sets;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
