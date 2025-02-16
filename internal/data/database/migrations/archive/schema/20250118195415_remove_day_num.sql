-- +goose Up
-- +goose StatementBegin
ALTER TABLE dates DROP COLUMN day_number;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
