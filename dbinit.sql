CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    username VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    about VARCHAR(255) NOT NULL,
    dp TEXT NOT NULL,
    last_online TIMESTAMP,
    created_at TIMESTAMP NOT NULL
)