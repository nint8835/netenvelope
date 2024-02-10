-- name: CreateUser :one
INSERT INTO users (username, password_hash)
VALUES (?, ?)
RETURNING *;

-- name: GetUserById :one
SELECT *
FROM users
WHERE id = ?;

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = ?;

-- name: UpdateUser :one
UPDATE users
SET username      = coalesce(sqlc.narg('username'), username),
    password_hash = coalesce(sqlc.narg('password_hash'), password_hash)
WHERE id = sqlc.arg('id')
RETURNING *;