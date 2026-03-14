INSERT INTO users (name, email, age, gender, birth_date) VALUES
('Alice Johnson', 'alice@example.com', 25, 'female', '1999-03-15'),
('Bob Smith', 'bob@example.com', 30, 'male', '1994-07-22'),
('Charlie Brown', 'charlie@example.com', 28, 'male', '1996-11-05'),
('Diana Prince', 'diana@example.com', 27, 'female', '1997-09-10'),
('Eve Adams', 'eve@example.com', 32, 'female', '1992-02-28'),
('Frank Miller', 'frank@example.com', 29, 'male', '1995-06-18'),
('Grace Lee', 'grace@example.com', 26, 'female', '1998-12-30'),
('Henry Wilson', 'henry@example.com', 31, 'male', '1993-04-25'),
('Iris Taylor', 'iris@example.com', 24, 'female', '2000-08-14'),
('Jack Davis', 'jack@example.com', 33, 'male', '1991-10-08'),
('Kate Moore', 'kate@example.com', 22, 'female', '2002-01-20'),
('Liam White', 'liam@example.com', 35, 'male', '1989-05-12'),
('Mia Clark', 'mia@example.com', 23, 'female', '2001-07-07'),
('Noah Harris', 'noah@example.com', 34, 'male', '1990-03-03'),
('Olivia Martin', 'olivia@example.com', 21, 'female', '2003-11-17'),
('Paul Garcia', 'paul@example.com', 36, 'male', '1988-09-09'),
('Quinn Robinson', 'quinn@example.com', 20, 'non-binary', '2004-04-04'),
('Rachel Lewis', 'rachel@example.com', 37, 'female', '1987-12-12'),
('Sam Walker', 'sam@example.com', 38, 'male', '1986-06-06');

CREATE TABLE IF NOT EXISTS user_friends (
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    friend_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, friend_id),
    CHECK (user_id != friend_id)
);

INSERT INTO user_friends (user_id, friend_id) VALUES
(2, 4), (4, 2),
(2, 5), (5, 2),
(2, 6), (6, 2),
(3, 4), (4, 3),
(3, 5), (5, 3),
(3, 6), (6, 3),
(2, 7), (7, 2),
(3, 8), (8, 3),
(4, 9), (9, 4),
(5, 10), (10, 5);