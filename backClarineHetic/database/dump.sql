CREATE TABLE IF NOT EXISTS users (
    uuid UUID PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);
