// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"time"
)

type User struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}
