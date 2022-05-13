CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    username VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    about VARCHAR(255) NOT NULL,
    dp TEXT NOT NULL,
    last_online TIMESTAMP,
    created_at TIMESTAMP NOT NULL
);
CREATE UNIQUE INDEX idx_email ON users(email);


CREATE TABLE IF NOT EXISTS friend_requests (
    id serial PRIMARY KEY,
    user_requesting_id INT NOT NULL,
    user_requested_id  INT NOT NULL,
    status_id   INT DEFAULT 0,
    created_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_requesting_id)
        REFERENCES users(id),
    FOREIGN KEY (user_requested_id)
        REFERENCES users(id)
);
CREATE UNIQUE INDEX user_request ON friend_requests(user_requesting_id, user_requested_id)