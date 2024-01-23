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
	ApiKey    string    `json:"api_key"`
}

func dbusertouser(dbuser database.User) User {
	return User{

		ID:        dbuser.ID,
		CreatedAt: dbuser.CreatedAt,
		UpdatedAt: dbuser.UpdatedAt,
		Name:      dbuser.Name,
		ApiKey:    dbuser.ApiKey,
	}

}
func dbfeedtofeed(dbfeed database.Feed) Feed {
	return Feed{
		ID:        dbfeed.ID,
		CreatedAt: dbfeed.CreatedAt,
		UpdatedAt: dbfeed.UpdatedAt,
		Name:      dbfeed.Name,
		Url:       dbfeed.Url,
		UserID:    dbfeed.UserID,
	}
}

func dbfeedstofeeds(dbfeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbfeed := range dbfeeds {
		feeds = append(feeds, dbfeedtofeed(dbfeed))
	}
	return feeds
}
