package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN export connection to database
var MongoCN = Connection()
var clientOptions = options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0-vevgu.mongodb.net/test?retryWrites=true&w=majority", os.Getenv("DB_USER"), os.Getenv("DB_PASS")))

// Connection method connecto to database
func Connection() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("DB Is Connected")

	return client
}

// CheckConnection verify the connection with database
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}

	return 1
}
