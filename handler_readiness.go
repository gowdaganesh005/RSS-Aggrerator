package main

import "net/http"

func handler_readiness(w http.ResponseWriter, r *http.Request) { //[2(1)]
	JSON_Response(w, 200, struct{}{})
}
