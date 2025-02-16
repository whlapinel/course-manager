-- +goose Up
-- +goose StatementBegin
INSERT INTO
    users (id, first_name, last_name, email)
VALUES
    (1, 'William', 'Lapinel', 'whlapinel@gmail.com');

CREATE TABLE
    new_terms (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL,
        start TEXT NOT NULL,
        end TEXT NOT NULL,
        description TEXT,
        user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE
    );

INSERT INTO
    new_terms (id, name, start, end, description, user_id)
SELECT
    id,
    name,
    start,
    end,
    description,
    1
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