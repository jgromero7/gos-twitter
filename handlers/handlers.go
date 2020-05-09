package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/jgromero7/gos-twitter/routes"
	"github.com/rs/cors"
)

// Handlers config server
func Handlers() {
	router := routes.Routes()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
