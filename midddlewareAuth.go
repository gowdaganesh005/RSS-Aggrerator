package main

import (
	"fmt"
	"net/http"

	"github.com/gowdaganesh005/RSS-Aggregator/internal/auth"
	"github.com/gowdaganesh005/RSS-Aggregator/internal/database"
)

type authHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apicn apiConfig) middlewareauth(authhandler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			Err_Response(w, 403, fmt.Sprintf("authorization failed and error:%v", err))
		}
		user, err := apicn.DB.GetUserByAPI(r.Context(), apikey)
		if err != nil {
			Err_Response(w, 400, fmt.Sprintf("could not get user %v", err))
		}
		authhandler(w, r, user)
	}
}
