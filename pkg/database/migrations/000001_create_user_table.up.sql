CREATE TABLE users
(
    id            INTEGER PRIMARY KEY,
    username      TEXT NOT NULL UNIQUE,
    password_hash BLOB NOT NULL
)