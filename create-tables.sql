DROP TABLE IF EXISTS Students;
CREATE TABLE Students (
    id      INT AUTO_INCREMENT NOT NULL,
    name    VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS Teachers;
CREATE TABLE Teachers (
    id      INT AUTO_INCREMENT NOT NULL,
    name    VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS Courses;
CREATE TABLE Courses (
  id            INT AUTO_INCREMENT NOT NULL,
  name          VARCHAR(255) NOT NULL,
  teacher_id    INT NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`teacher_id`) REFERENCES Teachers(id)
);

DROP TABLE IF EXISTS StudentCourse;
CREATE TABLE StudentCourse (
  id            INT AUTO_INCREMENT NOT NULL,
  student_id    INT NOT NULL,
  course_id     INT NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`student_id`) REFERENCES Students(id),
  FOREIGN KEY (`course_id`) REFERENCES Courses(id)
);

DROP TABLE IF EXISTS Posts;
CREATE TABLE Posts (
  id            INT AUTO_INCREMENT NOT NULL,
  teacher_id    INT NOT NULL,
  title         VARCHAR(255) NOT NULL,
  body          TEXT NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`teacher_id`) REFERENCES Teachers(id)
);

INSERT INTO Students
  (name)
VALUES
  ('Enzo'),
  ('Robin');

INSERT INTO Teachers
  (name)
VALUES
  ('Agus');

INSERT INTO Courses
  (name, teacher_id)
VALUES
  ('Framework Based Programming', 1);

INSERT INTO StudentCourse
  (student_id, course_id)
VALUES
  (1, 1),
  (2, 1);

INSERT INTO Posts
  (teacher_id, title, body)
VALUES
  (1, "Session 16", "This session will be the presentation of the final projects.");