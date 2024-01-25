-- name: Createpost :one
INSERT INTO posts (id , created_at, updated_at,title,description,published_at,url,feed_id)
VALUES (?,?,?,?,?,?,?,?)
RETURNING *;

-- name: GetpostsForUser :many
Select posts.* FROM posts
join feeds_follows on posts.feed_id=feeds_follows.feed_id
where feeds_follows.user_id=?
order by posts.published_at desc 
limit ?;

