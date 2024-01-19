-- name: Createuser :one
INSERT INTO users (id , created_at, updated_at, name)
VALUES (?,?,?,?)
RETURNING *;

