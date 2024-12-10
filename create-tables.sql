-- DROP TABLE IF EXISTS students;
-- CREATE TABLE students (
--     id      INT AUTO_INCREMENT NOT NULL,
--     name    VARCHAR(255) NOT NULL,
--     PRIMARY KEY (`id`)
-- );

-- DROP TABLE IF EXISTS lecturers;
-- CREATE TABLE lecturers (
--     id      INT AUTO_INCREMENT NOT NULL,
--     name    VARCHAR(255) NOT NULL,
--     PRIMARY KEY (`id`)
-- );

-- DROP TABLE IF EXISTS courses;
-- CREATE TABLE courses (
--   id            INT AUTO_INCREMENT NOT NULL,
--   name          VARCHAR(255) NOT NULL,
--   lecturer_id    INT NOT NULL,
--   PRIMARY KEY (`id`),
--   FOREIGN KEY (`lecturer_id`) REFERENCES lecturers(id)
-- );

-- DROP TABLE IF EXISTS student_course;
-- CREATE TABLE student_course (
--   id            INT AUTO_INCREMENT NOT NULL,
--   student_id    INT NOT NULL,
--   course_id     INT NOT NULL,
--   PRIMARY KEY (`id`),
--   FOREIGN KEY (`student_id`) REFERENCES students(id),
--   FOREIGN KEY (`course_id`) REFERENCES courses(id)
-- );

-- DROP TABLE IF EXISTS post;
-- CREATE TABLE post (
--   id            INT AUTO_INCREMENT NOT NULL,
--   lecturer_id    INT NOT NULL,
--   title         VARCHAR(255) NOT NULL,
--   body          TEXT NOT NULL,
--   PRIMARY KEY (`id`),
--   FOREIGN KEY (`lecturer_id`) REFERENCES lecturers(id)
-- );

-- INSERT INTO students
--   (name)
-- VALUES
--   ('Enzo'),
--   ('Robin');

-- INSERT INTO lecturers
--   (name)
-- VALUES
--   ('Agus');

-- INSERT INTO courses
--   (name, lecturer_id)
-- VALUES
--   ('Framework Based Programming', 1);

-- INSERT INTO student_course
--   (student_id, course_id)
-- VALUES
--   (1, 1),
--   (2, 1);

-- INSERT INTO post
--   (lecturer_id, title, body)
-- VALUES
--   (1, "Session 16", "This session will be the presentation of the final projects.");


-- Drop tables if they exist
DROP TABLE IF EXISTS student_courses;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS courses;
DROP TABLE IF EXISTS lecturers;
DROP TABLE IF EXISTS students;

-- Create students table
CREATE TABLE students (
    id      INT AUTO_INCREMENT NOT NULL,
    name    VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);

-- Create lecturers table
CREATE TABLE lecturers (
    id      INT AUTO_INCREMENT NOT NULL,
    name    VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);

-- Create courses table
CREATE TABLE courses (
  id            INT AUTO_INCREMENT NOT NULL,
  name          VARCHAR(255) NOT NULL,
  lecturer_id    INT NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`lecturer_id`) REFERENCES lecturers(id)
);

-- Create student_course table
CREATE TABLE student_courses (
  id            INT AUTO_INCREMENT NOT NULL,
  student_id    INT NOT NULL,
  course_id     INT NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`student_id`) REFERENCES students(id),
  FOREIGN KEY (`course_id`) REFERENCES courses(id)
);

-- Create post table
CREATE TABLE posts (
  id            INT AUTO_INCREMENT NOT NULL,
  lecturer_id    INT NOT NULL,
  course_id     INT NOT NULL,
  title         VARCHAR(255) NOT NULL,
  body          TEXT NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`lecturer_id`) REFERENCES lecturers(id),
  FOREIGN KEY (`course_id`) REFERENCES courses(id)
);

-- Insert students
INSERT INTO students (name) VALUES
  ('Alice Johnson'), ('Bob Smith'), ('Charlie Brown'), ('David Wilson'), ('Eva Davis'),
  ('Frank Moore'), ('Grace Taylor'), ('Hannah White'), ('Ian Thompson'), ('Jack Garcia'),
  ('Kara Martinez'), ('Liam Anderson'), ('Mia Harris'), ('Noah Lewis'), ('Olivia Clark'),
  ('Paul Walker'), ('Quincy Lee'), ('Rachel Hall'), ('Sam Young'), ('Tina Allen'),
  ('Uma Scott'), ('Victor Hill'), ('Wendy Adams'), ('Xander Baker'), ('Yara Nelson'),
  ('Zachary Carter'), ('Abby Green'), ('Ben King'), ('Caitlyn Phillips'), ('Daniel Cooper'),
  ('Ethan Turner'), ('Fiona Wright'), ('Gabe Morris'), ('Helen Mitchell'), ('Isla Perez'),
  ('Jake Rogers'), ('Kimberly Stewart'), ('Logan Howard'), ('Madison Wood'), ('Nathan Kelly'),
  ('Owen Rivera'), ('Piper Cox'), ('Quinn Ward'), ('Ryan Flores'), ('Samantha Barnes'),
  ('Tyler Ross'), ('Victoria Powell'), ('Will Brooks'), ('Xavier Bell'), ('Yasmine Reed');

-- Insert lecturers
INSERT INTO lecturers (name) VALUES
  ('Agus S'), ('Brian T'), ('Cathy V'), ('Daniel M'), ('Elaine R'),
  ('Frank Z'), ('Georgia P'), ('Hank F'), ('Isabel G'), ('Jack L'),
  ('Kara N'), ('Liam H'), ('Mandy J'), ('Nora C'), ('Oliver W');

-- Insert courses
INSERT INTO courses (name, lecturer_id) VALUES
  ('Math 101', 1), ('English Literature', 2), ('Physics', 3), 
  ('Chemistry', 4), ('Biology', 5), ('History', 6), 
  ('Computer Science', 7), ('Art', 8), ('Music', 9), 
  ('Physical Education', 10);

-- Insert student_course
INSERT INTO student_courses (student_id, course_id) VALUES
  (1, 1), (1, 2), (1, 3), (2, 1), (3, 2), (4, 3), (5, 4),
  (6, 5), (7, 6), (8, 7), (9, 8), (10, 9),
  (11, 10), (12, 1), (13, 2), (14, 3), (15, 4),
  (16, 5), (17, 6), (18, 7), (19, 8), (20, 9),
  (21, 10), (22, 1), (23, 2), (24, 3), (25, 4),
  (26, 5), (27, 6), (28, 7), (29, 8), (30, 9),
  (31, 10), (32, 1), (33, 2), (34, 3), (35, 4),
  (36, 5), (37, 6), (38, 7), (39, 8), (40, 9),
  (41, 10), (42, 1), (43, 2), (44, 3), (45, 4),
  (46, 5), (47, 6), (48, 7), (49, 8), (50, 9);

-- Insert post
INSERT INTO posts (lecturer_id, course_id, title, body) VALUES
  (1, 1, "Welcome to Math 101", "This is the first session of Math 101."),
  (2, 2, "Introduction to English Literature", "We will start with Shakespeare."),
  (3, 3, "Physics Fundamentals", "Understanding motion and forces."),
  (4, 4, "Chemistry Basics", "Today we discuss atoms."),
  (5, 5, "Biology 101", "Introduction to cellular biology."),
  (6, 6, "Historical Perspectives", "The Renaissance era."),
  (7, 7, "CS Intro", "Getting started with programming."),
  (8, 8, "Art Appreciation", "Analyzing famous paintings."),
  (9, 9, "Music Theory", "Understanding scales and chords."),
  (10, 10, "PE Introduction", "Overview of physical fitness activities.");

