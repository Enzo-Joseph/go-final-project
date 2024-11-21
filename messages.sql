DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id      INT AUTO_INCREMENT NOT NULL,
    username    VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS messages;
CREATE TABLE messages (
  id         INT AUTO_INCREMENT NOT NULL,
  user_id    INT NOT NULL,
  body       TEXT NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES users(id)
);

INSERT INTO users
  (username)
VALUES
  ('Enzo'),
  ('Robin');

INSERT INTO messages
  (user_id, body)
VALUES
  (1, 'First message'),
  (2, 'Another message');