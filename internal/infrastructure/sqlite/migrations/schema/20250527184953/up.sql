CREATE TABLE
    IF NOT EXISTS course_occasions (
        id INTEGER PRIMARY KEY,
        course_id INTEGER NOT NULL REFERENCES courses (id) ON DELETE CASCADE,
        date TEXT NOT NULL,
        name TEXT NOT NULL
    );