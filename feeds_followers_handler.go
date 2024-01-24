package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gowdaganesh005/RSS-Aggregator/internal/database"
)

type FeedFollows struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    string    `json:"user_id"`
	FeedID    string    `json:"feed_id"`
}

func (apicn *apiConfig) feeds_follow_handler(w http.ResponseWriter, r *http.Request, user database.User) { //[2(1)]
	type parameters struct {
		FeedID string `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	param := parameters{}
	err := decoder.Decode(&param)
	if err != nil {
		Err_Response(w, 400, fmt.Sprintf("error parsing JSON data: %v", err))
		return
	}
	feed_follow, err1 := apicn.DB.Createfeedfollow(r.Context(), database.CreatefeedfollowParams{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    param.FeedID,
	})
	if err1 != nil {
		Err_Response(w, 400, fmt.Sprintf("could not create feed follow: %v", err1))
		return
	}
	JSON_Response(w, 201, dbfeedfollowtofeedfollow(feed_follow))
}
func (apicn *apiConfig) Get_feeds_follow_handler(w http.ResponseWriter, r *http.Request, user database.User) { //[2(1)]

	feed_follow, err1 := apicn.DB.Getfeedfollow(r.Context(), user.ID)
	if err1 != nil {
		Err_Response(w, 400, fmt.Sprintf("could not get feed follows: %v", err1))
		return
	}
	JSON_Response(w, 201, dbfeedsfollowstofeedsfollows(feed_follow))
}
func (apicn *apiConfig) Delete_feeds_follow_handler(w http.ResponseWriter, r *http.Request, user database.User) { //[2(1)]
	feedfollowIDstr := chi.URLParam(r, "feedfollowID")

	err := apicn.DB.Deletefeedfollow(r.Context(), database.DeletefeedfollowParams{
		ID:     feedfollowIDstr,
		UserID: user.ID,
	})
	if err != nil {
		Err_Response(w, 400, fmt.Sprintf("Could not delete the feed follow:%v", err))
	}
	JSON_Response(w, 200, struct{}{})

}
