CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    email VARCHAR (128) UNIQUE NOT NULL,
    password_hash VARCHAR (128) NOT NULL
);
