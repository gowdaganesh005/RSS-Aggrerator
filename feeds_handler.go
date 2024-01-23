package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gowdaganesh005/RSS-Aggregator/internal/database"
)

type Feed struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    string    `json:"user_id"`
}

func (apicn *apiConfig) feeds_handler(w http.ResponseWriter, r *http.Request, user database.User) { //[2(1)]
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	param := parameters{}
	err := decoder.Decode(&param)
	if err != nil {
		Err_Response(w, 400, fmt.Sprintf("error parsing JSON data: %v", err))
		return
	}
	feed, err1 := apicn.DB.Createfeed(r.Context(), database.CreatefeedParams{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      param.Name,
		Url:       param.Url,
		UserID:    user.ID,
	})
	if err1 != nil {
		Err_Response(w, 400, fmt.Sprintf("could not create feed: %v", err1))
		return
	}
	JSON_Response(w, 201, dbfeedtofeed(feed))
}

func (apicn *apiConfig) Get_feeds_handler(w http.ResponseWriter, r *http.Request) { //[2(1)]

	feed, err1 := apicn.DB.GetFeed(r.Context())
	if err1 != nil {
		Err_Response(w, 400, fmt.Sprintf("could not get feed: %v", err1))
		return
	}
	JSON_Response(w, 201, dbfeedstofeeds(feed))
}
