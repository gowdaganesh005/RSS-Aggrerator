package main

import (
	"time"

	"github.com/gowdaganesh005/RSS-Aggregator/internal/database"
)

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func dbusertouser(dbuser database.User) User {
	return User{

		ID:        dbuser.ID,
		CreatedAt: dbuser.CreatedAt,
		UpdatedAt: dbuser.UpdatedAt,
		Name:      dbuser.Name,
	}

}
