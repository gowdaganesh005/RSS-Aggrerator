// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: posts.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createpost = `-- name: Createpost :one
INSERT INTO posts (id , created_at, updated_at,title,description,published_at,url,feed_id)
VALUES (?,?,?,?,?,?,?,?)
RETURNING id, created_at, updated_at, title, description, published_at, url, feed_id
`

type CreatepostParams struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Description sql.NullString
	PublishedAt time.Time
	Url         string
	FeedID      string
}

func (q *Queries) Createpost(ctx context.Context, arg CreatepostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createpost,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		arg.Description,
		arg.PublishedAt,
		arg.Url,
		arg.FeedID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Description,
		&i.PublishedAt,
		&i.Url,
		&i.FeedID,
	)
	return i, err
}

const getpostsForUser = `-- name: GetpostsForUser :many
Select posts.id, posts.created_at, posts.updated_at, posts.title, posts.description, posts.published_at, posts.url, posts.feed_id FROM posts
join feeds_follows on posts.feed_id=feeds_follows.feed_id
where feeds_follows.user_id=?
order by posts.published_at desc 
limit ?
`

type GetpostsForUserParams struct {
	UserID string
	Limit  int64
}

func (q *Queries) GetpostsForUser(ctx context.Context, arg GetpostsForUserParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getpostsForUser, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			&i.Description,
			&i.PublishedAt,
			&i.Url,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
