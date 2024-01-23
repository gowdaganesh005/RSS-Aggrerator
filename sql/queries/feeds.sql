-- name: Createfeed :one
INSERT INTO feeds (id , created_at, updated_at, name,url,user_id)
VALUES (?,?,?,?,?,?)
RETURNING *;

-- name: GetFeed :many
Select * from feeds;
