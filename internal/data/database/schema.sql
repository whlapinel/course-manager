CREATE TABLE
    goose_db_version (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        version_id INTEGER NOT NULL,
        is_applied INTEGER NOT NULL,
        tstamp TIMESTAMP DEFAULT (datetime ('now'))
    );

CREATE TABLE
    sqlite_sequence (name, seq);

CREATE TABLE
    IF NOT EXISTS courses (
        id INTEGER PRIMARY KEY,
        term_id INTEGER NOT NULL,
        name TEXT NOT NULL,
        description TEXT,
        std_set_id INTEGER REFERENCES standard_sets (id) ON DELETE CASCADE,
        FOREIGN KEY (term_id) REFERENCES terms (id) ON DELETE CASCADE
    );

CREATE TABLE
    IF NOT EXISTS units (
        id INTEGER PRIMARY KEY,
        course_id INTEGER NOT NULL,
        number INTEGER NOT NULL,
        sequence INTEGER NOT NULL,
        name TEXT NOT NULL,
        description TEXT,
        FOREIGN KEY (course_id) REFERENCES courses (id) ON DELETE CASCADE,
        UNIQUE (course_id, number)
    );

CREATE TABLE
    IF NOT EXISTS lessons (
        id INTEGER PRIMARY KEY,
        unit_id INTEGER NOT NULL,
        number INTEGER NOT NULL,
        name TEXT,
        description TEXT,
        FOREIGN KEY (unit_id) REFERENCES units (id) ON DELETE CASCADE,
        UNIQUE (unit_id, number)
    );

CREATE TABLE
    IF NOT EXISTS dates (
        id INTEGER PRIMARY KEY,
        term_id INTEGER NOT NULL,
        date TEXT NOT NULL,
        FOREIGN KEY (term_id) REFERENCES terms (id) ON DELETE CASCADE
    );

CREATE TABLE
    IF NOT EXISTS lesson_dates (
        lesson_id INTEGER NOT NULL,
        date_id INTEGER NOT NULL,
        FOREIGN KEY (lesson_id) REFERENCES lessons (id) ON DELETE CASCADE,
        FOREIGN KEY (date_id) REFERENCES dates (id) ON DELETE CASCADE,
        UNIQUE (lesson_id, date_id)
    );

CREATE TABLE
    IF NOT EXISTS standard_sets (
        id INTEGER PRIMARY KEY,
        course_name TEXT NOT NULL UNIQUE
    );

CREATE TABLE
    standards (
        id INTEGER PRIMARY KEY,
        number INTEGER NOT NULL,
        name TEXT NOT NULL,
        description TEXT,
        set_id INTEGER NOT NULL REFERENCES standard_sets (id) ON DELETE CASCADE,
        parent_id INTEGER REFERENCES standards (id) ON DELETE CASCADE
    );

CREATE TABLE
    lesson_standards (
        id INTEGER PRIMARY KEY,
        std_id INTEGER NOT NULL,
        lesson_id INTEGER NOT NULL,
        FOREIGN KEY (std_id) REFERENCES standards (id) ON DELETE CASCADE,
        FOREIGN KEY (lesson_id) REFERENCES lessons (id) ON DELETE CASCADE
    );

CREATE TABLE
    assessments (
        id INTEGER PRIMARY KEY,
        lesson_id INTEGER NOT NULL,
        name TEXT NOT NULL,
        instructions TEXT NOT NULL,
        category INTEGER NOT NULL,
        date_assigned TEXT NOT NULL,
        date_due TEXT NOT NULL,
        dropped INTEGER NOT NULL,
        FOREIGN KEY (lesson_id) REFERENCES lessons (id)
    );

CREATE TABLE
    users (
        id TEXT PRIMARY KEY,
        first_name TEXT NOT NULL,
        last_name TEXT NOT NULL,
        email TEXT NOT NULL,
        picture TEXT
    );

CREATE TABLE
    occasions (
        id INTEGER PRIMARY KEY,
        term_id INTEGER NOT NULL REFERENCES terms (id) ON DELETE CASCADE,
        date TEXT NOT NULL,
        name TEXT NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS terms (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL,
        start TEXT NOT NULL,
        end TEXT NOT NULL,
        description TEXT,
        user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE
    );