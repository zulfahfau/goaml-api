package main

import (
	"goaml-api/config"
	"goaml-api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to the database
	config.ConnectDB()

	// Initialize the router
	r := mux.NewRouter()

	// Define the endpoints
	r.HandleFunc("/session/login", controllers.Login).Methods("POST")
	r.HandleFunc("/session/logout", controllers.Logout).Methods("POST")

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
