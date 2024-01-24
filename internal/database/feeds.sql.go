// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: feeds.sql

package database

import (
	"context"
	"time"
)

const createfeed = `-- name: Createfeed :one
INSERT INTO feeds (id , created_at, updated_at, name,url,user_id)
VALUES (?,?,?,?,?,?)
RETURNING id, created_at, updated_at, name, url, user_id
`

type CreatefeedParams struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    string
}

func (q *Queries) Createfeed(ctx context.Context, arg CreatefeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createfeed,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Url,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
	)
	return i, err
}

const getFeed = `-- name: GetFeed :many
Select id, created_at, updated_at, name, url, user_id from feeds
`

func (q *Queries) GetFeed(ctx context.Context) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getFeed)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Url,
			&i.UserID,
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