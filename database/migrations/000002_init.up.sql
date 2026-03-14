CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(10),
    birth_date DATE,
    created_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO users (name, email, age, gender, birth_date) VALUES ('John Doe', 'john@example.com', 25, 'male', '1998-01-15');