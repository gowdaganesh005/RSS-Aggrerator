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

func dbfeedfollowtofeedfollow(dbfeedfollow database.FeedsFollow) FeedFollows {
	return FeedFollows{
		ID:        dbfeedfollow.ID,
		CreatedAt: dbfeedfollow.CreatedAt,
		UpdatedAt: dbfeedfollow.UpdatedAt,

		UserID: dbfeedfollow.UserID,
		FeedID: dbfeedfollow.FeedID,
	}
}
func dbfeedsfollowstofeedsfollows(dbfeedfollows []database.FeedsFollow) []FeedFollows {
	feedsfollow := []FeedFollows{}
	for _, dbfeedfollow := range dbfeedfollows {
		feedsfollow = append(feedsfollow, dbfeedfollowtofeedfollow(dbfeedfollow))
	}
	return feedsfollow
}

type Post struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      string    `json:"feed_id"`
}

func dbposttopost(dbpost database.Post) Post {
	var description *string
	if dbpost.Description.Valid {
		description = &dbpost.Description.String
	}
	return Post{
		ID:          dbpost.ID,
		CreatedAt:   dbpost.CreatedAt,
		UpdatedAt:   dbpost.UpdatedAt,
		Title:       dbpost.Title,
		Description: description,
		PublishedAt: dbpost.PublishedAt,
		Url:         dbpost.Url,
		FeedID:      dbpost.FeedID,
	}
}
func dbpoststoposts(dbposts []database.Post) []Post {
	posts := []Post{}
	for _, post := range dbposts {
		posts = append(posts, dbposttopost(post))
	}
	return posts
}
