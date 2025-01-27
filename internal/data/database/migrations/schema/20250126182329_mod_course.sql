-- +goose Up
-- +goose StatementBegin
ALTER TABLE courses 
ADD COLUMN 
std_set_id INTEGER REFERENCES standard_sets(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
