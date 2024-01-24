package main

import (
	"context"
	"log"
	"sync"
	"time"

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
		log.Println("Found post ", item.Title, "on feed", feed.Name)
	}
	log.Printf("Found %s collected %v posts found", feed.Name, len(rssfeed.Channel.Item))
}
