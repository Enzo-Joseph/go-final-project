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
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id      INT AUTO_INCREMENT NOT NULL,
    name    VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role    VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);

-- Create courses table
CREATE TABLE courses (
  id            INT AUTO_INCREMENT NOT NULL,
  name          VARCHAR(255) NOT NULL,
  description   TEXT,
  lecturer_id    INT NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`lecturer_id`) REFERENCES users(id)
);

-- Create student_course table
CREATE TABLE student_courses (
  id            INT AUTO_INCREMENT NOT NULL,
  student_id    INT NOT NULL,
  course_id     INT NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`student_id`) REFERENCES users(id),
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
  FOREIGN KEY (`lecturer_id`) REFERENCES users(id),
  FOREIGN KEY (`course_id`) REFERENCES courses(id)
);

-- Insert Users
INSERT INTO users (name, username, password, role) VALUES
  ('Brian Taylor', 'brian', 'password', 'lecturer'),
  ('Cathy V', 'cathy', 'password', 'lecturer'),
  ('Agus', 'agus', 'password', 'lecturer'),
  ('Enzo', 'enzo', 'password', 'student'),
  ('Robin', 'robin', 'password', 'student'),
  ('Alice Johnson', 'alice', 'password', 'student'),
  ('Bob Smith', 'bob', 'password', 'student'),
  ('Charlie Brown', 'charlie', 'password', 'student'),
  ('David Wilson', 'david', 'password', 'student'),
  ('Eva Davis', 'eva', 'password', 'student'),
  ('Frank Moore', 'frank', 'password', 'student'),
  ('Grace Taylor', 'grace', 'password', 'student'),
  ('Hannah White', 'hannah', 'password', 'student'),
  ('Ian Thompson', 'ian', 'password', 'student');

-- Insert courses
INSERT INTO courses (name, description, lecturer_id) VALUES
  ('Math 101', 'The basics of math.', 1), ('English Literature', 'The history of English literature.', 2), ('Physics', 'Fundamentals of physics.', 3), 
  ('Chemistry', 'The science of chemistry.', 1), ('Biology', 'The study of life.', 2), ('History', 'The history of the world.', 3), 
  ('Computer Science', 'Getting started with programming.', 1), ('Art', 'Analyzing famous paintings.', 2), ('Music', 'Understanding scales and chords.', 3), 
  ('Physical Education', 'Overview of physical fitness activities.', 1);

-- Insert student_course
INSERT INTO student_courses (student_id, course_id) VALUES
  (4, 3), (5, 4),
  (6, 5), (7, 6), (8, 7), (9, 8), (10, 9),
  (11, 10), (12, 1), (13, 2), (14, 3);

-- Insert post
INSERT INTO posts (lecturer_id, course_id, title, body) VALUES
  (1, 1, "Welcome to Math 101", "This is the first session of Math 101."),
  (2, 2, "Introduction to English Literature", "We will start with Shakespeare."),
  (3, 3, "Physics Fundamentals", "Understanding motion and forces."),
  (1, 4, "Chemistry Basics", "Today we discuss atoms."),
  (2, 5, "Biology 101", "Introduction to cellular biology."),
  (3, 6, "Historical Perspectives", "The Renaissance era."),
  (1, 7, "CS Intro", "Getting started with programming."),
  (2, 8, "Art Appreciation", "Analyzing famous paintings."),
  (3, 9, "Music Theory", "Understanding scales and chords."),
  (1, 10, "PE Introduction", "Overview of physical fitness activities.");

