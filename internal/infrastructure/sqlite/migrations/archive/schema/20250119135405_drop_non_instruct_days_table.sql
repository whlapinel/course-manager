-- +goose Up
-- +goose StatementBegin
DROP TABLE non_instruct_days;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
