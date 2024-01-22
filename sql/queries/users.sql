-- name: Createuser :one
INSERT INTO users (id , created_at, updated_at, name,api_key)
VALUES (?,?,?,?,
lower(hex(randomblob(32)))
)
RETURNING *;

-- name: GetUserByAPI :one
SELECT * FROM users WHERE api_key=?;