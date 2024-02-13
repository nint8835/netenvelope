-- name: ListVlans :many
SELECT *
from vlans;

-- name: CreateVlan :one
INSERT INTO vlans (tag, name)
VALUES (?, ?)
RETURNING *;
