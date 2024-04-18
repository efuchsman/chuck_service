package main

import (
	config "chuck_service/config"
	chuckhandler "chuck_service/handlers/chuck_norris"
	chucknorris "chuck_service/internal/chuck_norris"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Starting the application")
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	var httpClient chucknorris.HTTPClient = &http.Client{}
	chuckService := chucknorris.NewChuckNorrisService(httpClient)

	// Setup the HTTP server and router
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the local Chuck Norris API")
	})

	chuckHandler := chuckhandler.NewChuckNorrisHandler(chuckService)
	router.HandleFunc("/random", chuckHandler.GetRandomJoke).Methods("GET")

	handler := config.SetCORS(router)

	port := 8000
	fmt.Printf("Server is running on :%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	fmt.Println("Application started successfully")
}
