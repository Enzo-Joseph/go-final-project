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
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (`id`)
);

-- Create courses table
CREATE TABLE courses (
  id            INT AUTO_INCREMENT NOT NULL,
  name          VARCHAR(255) NOT NULL,
  description   TEXT,
  lecturer_id    INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`lecturer_id`) REFERENCES users(id)
);

-- Create student_course table
CREATE TABLE student_courses (
  id            INT AUTO_INCREMENT NOT NULL,
  student_id    INT NOT NULL,
  course_id     INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL,
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
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL,
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
  ('Ian Thompson', 'ian', 'password', 'student'),
  ('Jack Harris', 'jack', 'password', 'student'),
  ('Katie Brown', 'katie', 'password', 'student'),
  ('Liam Johnson', 'liam', 'password', 'student'),
  ('Mia Davis', 'mia', 'password', 'student'),
  ('Noah Wilson', 'noah', 'password', 'student'),
  ('Olivia Moore', 'olivia', 'password', 'student');

-- Insert courses
INSERT INTO courses (name, description, lecturer_id) VALUES
('Database', 'Introduction to Databases', 1),
('OOP', 'Object-Oriented Programming', 2),
('Web', 'Web Development Basics', 1),
('Big Data', 'All About Big Data Analytics', 2);

-- Insert student_course
INSERT INTO student_courses (student_id, course_id) VALUES
(1, 1), (2, 1), (3, 2), (4, 3), (5, 3), (6, 2), (7, 1), (8, 1), (9, 4), (10, 2), (11, 1), (12, 4), (13, 4), (14, 2), (15, 2), (16, 3), (17, 2), (18, 1), (19, 2), (20, 4);

-- Insert post
INSERT INTO posts (lecturer_id, course_id, title, body) VALUES
(1, 1, 'What is a Database?', 'A database is a structured collection of data. In this post, we discuss the fundamentals of databases and their role in organizing and managing information.'),
(1, 1, 'Types of Databases', 'Databases come in different types, including relational, non-relational, and distributed. This post explores the different types and their use cases.'),
(1, 1, 'Basic SQL Commands', 'SQL is the language used to interact with databases. Here, we cover the basic SQL commands like SELECT, INSERT, UPDATE, and DELETE.'),
(1, 1, 'Database Normalization', 'Normalization is the process of organizing data to avoid redundancy. This post explains the principles and techniques of database normalization.');

INSERT INTO posts (lecturer_id, course_id, title, body) VALUES
(2, 2, 'What is Object-Oriented Programming?', 'Object-Oriented Programming (OOP) is a programming paradigm based on the concept of objects. This post introduces the core principles of OOP, including classes and objects.'),
(2, 2, 'Key Concepts of OOP: Encapsulation, Inheritance, and Polymorphism', 'In this post, we dive into the key concepts of OOP, such as encapsulation, inheritance, and polymorphism, and how they help structure code effectively.'),
(2, 2, 'Understanding Classes and Objects in OOP', 'This post explains the relationship between classes and objects, demonstrating how classes are blueprints for creating objects in OOP.'),
(2, 2, 'Benefits of Object-Oriented Programming', 'OOP offers many advantages, including code reusability and easier maintenance. This post discusses the benefits of using OOP in modern software development.');

INSERT INTO posts (lecturer_id, course_id, title, body) VALUES
(1, 3, 'Introduction to Web Development', 'Web development involves creating websites and applications. This post covers the basics of web development, including the difference between front-end and back-end development.'),
(1, 3, 'HTML: The Structure of Web Pages', 'HTML is the foundation of web pages. In this post, we explore the role of HTML in structuring web content and introduce essential HTML tags and elements.'),
(1, 3, 'CSS: Styling Your Web Pages', 'CSS is used to style HTML elements. This post discusses the basics of CSS, including how to use selectors, properties, and values to enhance the appearance of your website.'),
(1, 3, 'Introduction to JavaScript for Web Development', 'JavaScript is the language that adds interactivity to web pages. Learn the basics of JavaScript and how it interacts with HTML and CSS to create dynamic websites.'),
(1, 3, 'Responsive Web Design: Making Websites Mobile-Friendly', 'In today’s mobile-first world, responsive design is essential. This post introduces the principles of responsive web design and techniques for creating websites that look great on any device.'),
(1, 3, 'Web Development Tools and Frameworks', 'Web developers use various tools and frameworks to streamline their work. This post highlights popular tools and frameworks like Visual Studio Code, Bootstrap, and React for efficient web development.');

INSERT INTO posts (lecturer_id, course_id, title, body) VALUES
(2, 4, 'Introduction to Big Data Analytics', 'Big data analytics involves processing large volumes of data to uncover hidden patterns, trends, and insights. This post introduces the concept of big data and its importance in modern analytics.'),
(2, 4, 'Tools for Big Data Analytics: Hadoop and Spark', 'Hadoop and Spark are two popular tools for handling big data. This post explores their architecture and how they are used to process and analyze large datasets efficiently.'),
(2, 4, 'Data Mining Techniques in Big Data', 'Data mining plays a critical role in big data analytics. This post covers common data mining techniques like clustering, classification, and regression, and their applications in big data environments.');