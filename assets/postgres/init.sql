CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name Varchar(255) NOT NULL,
    last_login TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    likes INT DEFAULT 0,
    created_at TIMESATMP
)