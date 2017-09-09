package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"./handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading environment variables.")
	}
}

func main() {
	ApiHost := os.Getenv("API_HOST")

	app := mux.NewRouter()

	app.HandleFunc("/posts/{page}/", v2handlers.PostsHandler).Methods("GET")
	app.HandleFunc("/cat/{catSlug}/{page}/", v2handlers.PostsByCatHandler).Methods("GET")
	app.HandleFunc("/cats/{page}/", v2handlers.CategoriesHandler).Methods("GET")
	app.HandleFunc("/post/{postSllug}/", v2handlers.PostHandler).Methods("GET")
	app.HandleFunc("/search/{page}/{keyword}/", v2handlers.SearchHandler).Methods("GET")

	server := &http.Server{
		Handler:      app,
		Addr:         ApiHost + ":" + os.Getenv("API_PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())

}
