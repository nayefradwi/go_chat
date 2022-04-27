CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    username VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    about VARCHAR(255),
    dp TEXT,
    last_online TIMESTAMP,
    created_on TIMESTAMP NOT NULL
)