ALTER TABLE assessments
ADD COLUMN course_id INT REFERENCES courses (id)