package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gowdaganesh005/RSS-Aggregator/internal/database"
)

func startscraping(
	db *database.Queries,
	concurrency int,
	timebetweenRequest time.Duration,

) {
	log.Printf("Scraping on %v goroutines every %v duration", concurrency, timebetweenRequest)
	ticker := time.NewTicker(timebetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedtoFetch(context.Background(), int64(concurrency))
		if err != nil {
			log.Println("error fetching feeds:", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapefeed(db, wg, feed)
		}
		wg.Wait()

	}

}
func scrapefeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched:", err)
		return
	}
	rssfeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed", err)
	}

	for _, item := range rssfeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		pubat, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Could not parse date %v with err %v", item.PubDate, err)
		}
		_, err1 := db.Createpost(context.Background(), database.CreatepostParams{
			ID:          uuid.New().String(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: description,
			PublishedAt: pubat,
			Url:         item.Link,
			FeedID:      feed.ID,
		})
		if err1 != nil {
			if strings.Contains(err1.Error(), "UNIQUE constraint failed") {
				continue
			}
			log.Println("failed to create a post", err1)
		}
	}
	log.Printf("Found %s collected %v posts found", feed.Name, len(rssfeed.Channel.Item))
}
