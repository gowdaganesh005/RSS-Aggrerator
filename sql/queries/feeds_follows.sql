-- name: Createfeedfollow :one
INSERT INTO feeds_follows (id , created_at, updated_at, user_id,feed_id)
VALUES (?,?,?,?,?)
RETURNING *;

-- name: Getfeedfollow :many
SELECT * from feeds_follows where user_id=?;


