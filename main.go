package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/gowdaganesh005/RSS-Aggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

// here we are craeting a struct to link the db and we are using sqlite3 as our database
type apiConfig struct {
	DB *database.Queries
}

func main() {
	// load environment variable mentioned in .env file or present in environment
	godotenv.Load()

	// get port from enviornment variables *PORT [1(1)] refer in main_docs.md
	portString := os.Getenv("PORT")
	//if not set return with error
	if portString == "" {
		log.Fatal("port is not found in the environment")
	}
	fmt.Println("PORT =", portString)

	dburl := os.Getenv("DB_URL")
	//if not set return with error
	if dburl == "" {
		log.Fatal("database is not found in the environment")
	}

	conn, err1 := sql.Open("sqlite3", dburl)
	if err1 != nil {
		log.Printf("Could not connect to database:%v", err1)
		return
	}
	apicn := apiConfig{
		DB: database.New(conn),
	}

	//a new chi router  Router [1(2)]
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		//AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// A Sub router for  v1 route
	v1router := chi.NewRouter()
	v1router.Get("/healthz", handler_readiness) // handler functions(handler_readiness) for GET http method
	v1router.Get("/err", errhandler_readiness)
	v1router.Post("/users", apicn.User_creating_handler)
	v1router.Get("/users", apicn.GetUserByAPI)
	router.Mount("/v1", v1router) // mounting over root router

	// determines the behaviour of the httpp server
	//[1(4)]
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	fmt.Printf("Server starting on port %v", portString)

	//err := srv.ListenAndServe(): Starts the HTTP server and logs any errors. The server will listen for incoming requests on the specified port
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
