-- name: Createuser :one
INSERT INTO users (id , created_at, updated_at, name,api_key)
VALUES (?,?,?,?,
encode(hex(randomblob(32)))
)
RETURNING *;

