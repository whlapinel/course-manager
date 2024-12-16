CREATE TABLE IF NOT EXISTS courses (
    id INTEGER PRIMARY KEY,
    template_id INTEGER,
    term_id INTEGER,
    name TEXT NOT NULL,
    description TEXT,
    FOREIGN KEY (term_id) REFERENCES terms(id),
    FOREIGN KEY (template_id) REFERENCES courses(id)
);


CREATE TABLE IF NOT EXISTS units (
    id INTEGER PRIMARY KEY,
    course_id INTEGER NOT NULL,
    template_id INTEGER,
    number INTEGER NOT NULL,
    sequence INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    FOREIGN KEY (course_id) REFERENCES courses(id),
    UNIQUE(course_id, number)
);

CREATE TABLE IF NOT EXISTS lessons (
    id INTEGER PRIMARY KEY,
    unit_id INTEGER NOT NULL,
    template_id INTEGER,
    number INTEGER NOT NULL,
    name TEXT,
    description TEXT,
    FOREIGN KEY (unit_id) REFERENCES units(id),
    FOREIGN KEY (template_id) REFERENCES lessons(id),
    UNIQUE(unit_id, number)
);

CREATE TABLE IF NOT EXISTS dates (
    id INTEGER PRIMARY KEY,
    term_id INTEGER NOT NULL,
    day_number INTEGER NOT NULL,
    date TEXT NOT NULL,
    FOREIGN KEY (term_id) REFERENCES terms(id)
);

CREATE TABLE IF NOT EXISTS lesson_dates (
    lesson_id INTEGER NOT NULL,
    date_id INTEGER NOT NULL,
    FOREIGN KEY (lesson_id) REFERENCES lessons(id),
    FOREIGN KEY (date_id) REFERENCES dates(id),
    UNIQUE(lesson_id, date_id)
);


CREATE TABLE IF NOT EXISTS terms (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    start TEXT NOT NULL,
    end TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS non_instruct_days (
    id INTEGER PRIMARY KEY,
    term_id INTEGER NOT NULL,
    date TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS standards (
    id INTEGER PRIMARY KEY,
    course_id INTEGER NOT NULL,
    number INTEGER NOT NULL,
    name TEXT NOT NULL,
    FOREIGN KEY (course_id) REFERENCES courses(id)
);

CREATE TABLE IF NOT EXISTS unit_standards (
    id INTEGER PRIMARY KEY,
    course_id INTEGER NOT NULL,
    standard_id INTEGER NOT NULL,
    FOREIGN KEY (course_id) REFERENCES courses(id),
    FOREIGN KEY (standard_id) REFERENCES standards(id)
);

CREATE TABLE IF NOT EXISTS objectives (
    id INTEGER PRIMARY KEY,
    std_id INTEGER NOT NULL,
    number INTEGER NOT NULL,
    name TEXT NOT NULL,
    FOREIGN KEY (std_id) REFERENCES standards(id)
);

CREATE TABLE IF NOT EXISTS lesson_objectives (
    id INTEGER PRIMARY KEY,
    obj_id INTEGER NOT NULL,
    lesson_id INTEGER NOT NULL,
    FOREIGN KEY (obj_id) REFERENCES objectives(id),
    FOREIGN KEY (lesson_id) REFERENCES lessons(id)
);

