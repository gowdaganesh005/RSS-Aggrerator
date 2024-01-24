-- name: Createfeed :one
INSERT INTO feeds (id , created_at, updated_at, name,url,user_id)
VALUES (?,?,?,?,?,?)
RETURNING *;

-- name: GetFeed :many
Select * from feeds;

-- name: GetNextFeedtoFetch :many
select * from feeds
order by lastfetched_at Asc nulls first
limit ?;

-- name: MarkAsFetched :one
update  feeds
set lastfetched_at=CURRENT_TIMESTAMP,
updated_at=CURRENT_TIMESTAMP
where id=? 
returning *;

