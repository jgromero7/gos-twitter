package models

import (
	"context"
	"time"

	"github.com/jgromero7/gos-twitter/app/services"
	"github.com/jgromero7/gos-twitter/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User data structure
type User struct {
	ID        primitive.ObjectID `bson: "_id,omitempty" json:"id"`
	Name      string             `bson: "name,omitempty" json:"name,omitempty"`
	LastName  string             `bson: "lastName,omitempty" json:"lastName,omitempty"`
	BirthDate time.Time          `bson: "birthDate,omitempty" json:"birthDate,omitempty"`
	Email     string             `bson: "email" json:"email"`
	Password  string             `bson: "password" json:"password,omitempty"`
	Avatar    string             `bson: "avatar,omitempty" json:"avatar,omitempty"`
	Banner    string             `bson: "banner,omitempty" json:"banner,omitempty"`
	Location  string             `bson: "location,omitempty" json:"location,omitempty"`
	WebSite   string             `bson: "webSite,omitempty" json:"webSite,omitempty"`
}

// CreateUser store a user
func CreateUser(user User) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := database.MongoCN.Database("gos-twitter")
	collection := database.Collection("user")

	var err error
	user.Password, err = services.EncryptPassword(user.Password)

	if err != nil {
		return "", false, err
	}

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}

// ExistsUser verify if a user exists
func ExistsUser(email string) (User, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := database.MongoCN.Database("gos-twitter")
	collection := database.Collection("user")

	condition := bson.M{"email": email}

	var user User

	err := collection.FindOne(ctx, condition).Decode(&user)
	ID := user.ID.Hex()
	if err != nil {
		return user, false, ID
	}

	return user, true, ID
}
