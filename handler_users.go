package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gowdaganesh005/RSS-Aggregator/internal/database"
)

func (apicn *apiConfig) User_creating_handler(w http.ResponseWriter, r *http.Request) { //[2(1)]
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	param := parameters{}
	err := decoder.Decode(&param)
	if err != nil {
		Err_Response(w, 400, fmt.Sprintf("error parsing JSON data: %v", err))
		return
	}
	user, err1 := apicn.DB.Createuser(r.Context(), database.CreateuserParams{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      param.Name,
	})
	if err1 != nil {
		Err_Response(w, 400, fmt.Sprintf("could not create user: %v", err1))
	}
	JSON_Response(w, 201, dbusertouser(user))
}

func (apicn *apiConfig) GetUserByAPI(w http.ResponseWriter, r *http.Request, user database.User) { //[2(1)]

	JSON_Response(w, 200, dbusertouser(user))
}

func (apicn *apiConfig) user_post_handler(w http.ResponseWriter, r *http.Request, user database.User) { //[2(1)]
	posts, err := apicn.DB.GetpostsForUser(r.Context(), database.GetpostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		log.Println("Could not get posts for the user: ", err)
	}
	JSON_Response(w, 200, dbpoststoposts(posts))

}
