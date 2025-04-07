CREATE TABLE
    assessments (
        id INTEGER PRIMARY KEY,
        course_id INTEGER REFERENCES courses (id),
        unit_id INTEGER REFERENCES units (id),
        lesson_id INTEGER REFERENCES lessons (id),
        name TEXT NOT NULL,
        instructions TEXT NOT NULL,
        file TEXT,
        category TEXT NOT NULL,
        date_assigned TEXT NOT NULL,
        date_due TEXT NOT NULL,
        dropped INTEGER NOT NULL
    )