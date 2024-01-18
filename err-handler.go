package main

import "net/http"

func errhandler_readiness(w http.ResponseWriter, r *http.Request) { //[2(1)]
	Err_Response(w, 400, "Something went wrong")
}
