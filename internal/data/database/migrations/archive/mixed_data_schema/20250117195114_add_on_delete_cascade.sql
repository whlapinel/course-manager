-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS courses_new (
    id INTEGER PRIMARY KEY,
    term_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    FOREIGN KEY (term_id) REFERENCES terms(id) ON DELETE CASCADE 
);

CREATE TABLE IF NOT EXISTS units_new (
    id INTEGER PRIMARY KEY,
    course_id INTEGER NOT NULL,
    number INTEGER NOT NULL,
    sequence INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE,
    UNIQUE(course_id, number)
);

CREATE TABLE IF NOT EXISTS lessons_new (
    id INTEGER PRIMARY KEY,
    unit_id INTEGER NOT NULL,
    number INTEGER NOT NULL,
    name TEXT,
    description TEXT,
    FOREIGN KEY (unit_id) REFERENCES units(id) ON DELETE CASCADE,
    UNIQUE(unit_id, number)
);

CREATE TABLE IF NOT EXISTS images_new (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    base_path TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS course_images_new (
    course_id INTEGER NOT NULL,
    image_id INTEGER NOT NULL,
    FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE,
    FOREIGN KEY (image_id) REFERENCES images(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS unit_images_new (
    unit_id INTEGER NOT NULL,
    image_id INTEGER NOT NULL,
    FOREIGN KEY (unit_id) REFERENCES units(id) ON DELETE CASCADE,
    FOREIGN KEY (image_id) REFERENCES images(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS lesson_images_new (
    lesson_id INTEGER NOT NULL,
    image_id INTEGER NOT NULL,
    FOREIGN KEY (lesson_id) REFERENCES lessons(id) ON DELETE CASCADE,
    FOREIGN KEY (image_id) REFERENCES images(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS dates_new (
    id INTEGER PRIMARY KEY,
    term_id INTEGER NOT NULL,
    day_number INTEGER NOT NULL,
    date TEXT NOT NULL,
    FOREIGN KEY (term_id) REFERENCES terms(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS lesson_dates_new (
    lesson_id INTEGER NOT NULL,
    date_id INTEGER NOT NULL,
    FOREIGN KEY (lesson_id) REFERENCES lessons(id) ON DELETE CASCADE,
    FOREIGN KEY (date_id) REFERENCES dates(id) ON DELETE CASCADE,
    UNIQUE(lesson_id, date_id)
);


CREATE TABLE IF NOT EXISTS terms_new (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    start TEXT NOT NULL,
    end TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS non_instruct_days_new (
    id INTEGER PRIMARY KEY,
    term_id INTEGER NOT NULL,
    date TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS standards_new (
    id INTEGER PRIMARY KEY,
    course_id INTEGER NOT NULL,
    number INTEGER NOT NULL,
    name TEXT NOT NULL,
    FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS unit_standards_new (
    id INTEGER PRIMARY KEY,
    course_id INTEGER NOT NULL,
    standard_id INTEGER NOT NULL,
    FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE,
    FOREIGN KEY (standard_id) REFERENCES standards(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS objectives_new (
    id INTEGER PRIMARY KEY,
    std_id INTEGER NOT NULL,
    number INTEGER NOT NULL,
    name TEXT NOT NULL,
    FOREIGN KEY (std_id) REFERENCES standards(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS lesson_objectives_new (
    id INTEGER PRIMARY KEY,
    obj_id INTEGER NOT NULL,
    lesson_id INTEGER NOT NULL,
    FOREIGN KEY (obj_id) REFERENCES objectives(id) ON DELETE CASCADE,
    FOREIGN KEY (lesson_id) REFERENCES lessons(id) ON DELETE CASCADE
);

INSERT INTO terms_new (id, name, start, end)
SELECT id, name, start, end FROM terms;

INSERT INTO courses_new (id, term_id, name, description)
SELECT id, term_id, name, description FROM courses;

INSERT INTO units_new (id, course_id, number, sequence, name, description)
SELECT id, course_id, number, sequence, name, description FROM units;

INSERT INTO lessons_new (id, unit_id, number, name, description)
SELECT id, unit_id, number, name, description FROM lessons;

INSERT INTO images_new (id, name, description, base_path)
SELECT id, name, description, base_path FROM images;

INSERT INTO course_images_new (course_id, image_id)
SELECT course_id, image_id FROM course_images;

INSERT INTO unit_images_new (unit_id, image_id)
SELECT unit_id, image_id FROM unit_images;

INSERT INTO lesson_images_new (lesson_id, image_id)
SELECT lesson_id, image_id FROM lesson_images;

INSERT INTO dates_new (id, term_id, day_number, date)
SELECT id, term_id, day_number, date FROM dates;

INSERT INTO lesson_dates_new (lesson_id, date_id)
SELECT lesson_id, date_id FROM lesson_dates;

INSERT INTO non_instruct_days_new (id, term_id, date)
SELECT id, term_id, date FROM non_instruct_days;

INSERT INTO standards_new (id, course_id, number, name)
SELECT id, course_id, number, name FROM standards;

INSERT INTO unit_standards_new (id, course_id, standard_id)
SELECT id, course_id, standard_id FROM unit_standards;

INSERT INTO objectives_new (id, std_id, number, name)
SELECT id, std_id, number, name FROM objectives;

INSERT INTO lesson_objectives_new (id, obj_id, lesson_id)
SELECT id, obj_id, lesson_id FROM lesson_objectives;

DROP TABLE terms;
DROP TABLE courses;
DROP TABLE units;
DROP TABLE lessons;
DROP TABLE images;
DROP TABLE course_images;
DROP TABLE unit_images;
DROP TABLE lesson_images;
DROP TABLE dates;
DROP TABLE lesson_dates;
DROP TABLE non_instruct_days;
DROP TABLE standards;
DROP TABLE unit_standards;
DROP TABLE objectives;
DROP TABLE lesson_objectives;

ALTER TABLE terms_new RENAME TO terms;
ALTER TABLE courses_new RENAME TO courses;
ALTER TABLE units_new RENAME TO units;
ALTER TABLE lessons_new RENAME TO lessons;
ALTER TABLE images_new RENAME TO images;
ALTER TABLE course_images_new RENAME TO course_images;
ALTER TABLE unit_images_new RENAME TO unit_images;
ALTER TABLE lesson_images_new RENAME TO lesson_images;
ALTER TABLE dates_new RENAME TO dates;
ALTER TABLE lesson_dates_new RENAME TO lesson_dates;
ALTER TABLE non_instruct_days_new RENAME TO non_instruct_days;
ALTER TABLE standards_new RENAME TO standards;
ALTER TABLE unit_standards_new RENAME TO unit_standards;
ALTER TABLE objectives_new RENAME TO objectives;
ALTER TABLE lesson_objectives_new RENAME TO lesson_objectives;



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
