-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    new_terms (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL,
        start TEXT NOT NULL,
        end TEXT NOT NULL,
        description TEXT,
        user_id TEXT NOT NULL REFERENCES users (id) ON DELETE CASCADE
    );

INSERT INTO
    new_terms (id, name, start, end, description, user_id)
SELECT
    id,
    name,
    start,
    end,
    description,
    '101602110272674353046'
FROM
    terms;

DROP TABLE terms;

ALTER TABLE new_terms
RENAME TO terms;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd