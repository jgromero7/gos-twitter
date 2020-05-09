package main

import (
	"log"

	"github.com/jgromero7/gos-twitter/database"
	"github.com/jgromero7/gos-twitter/handlers"
)

func main() {

	if database.CheckConnection() == 0 {
		log.Fatal("no connection to the database")
		return
	}

	handlers.Handlers()

}
