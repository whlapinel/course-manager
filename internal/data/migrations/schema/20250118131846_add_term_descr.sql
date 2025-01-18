-- +goose Up
-- +goose StatementBegin
ALTER TABLE terms ADD COLUMN description TEXT;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
