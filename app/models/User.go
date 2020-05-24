package models

import (
	"context"
	"time"

	"github.com/jgromero7/gos-twitter/app/services"
	"github.com/jgromero7/gos-twitter/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// User data structure
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	LastName  string             `bson:"lastName,omitempty" json:"lastName,omitempty"`
	BirthDate time.Time          `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Email     string             `bson:"email" json:"email,omitempty"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar,omitempty" json:"avatar,omitempty"`
	Banner    string             `bson:"banner,omitempty" json:"banner,omitempty"`
	Biography string             `bson:"biography,omitempty" json:"biography,omitempty"`
	Location  string             `bson:"location,omitempty" json:"location,omitempty"`
	WebSite   string             `bson:"webSite,omitempty" json:"webSite,omitempty"`
}

// ReadAllUsers get all user by contifion for parameters
func ReadAllUsers(ID string, page int64, search string, typeUser string) ([]*User, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := database.MongoCN.Database("gos-twitter")
	collection := database.Collection("user")

	var users []*User

	searchOption := options.Find()
	searchOption.SetSkip((page - 1) * 20)
	searchOption.SetLimit(20)

	condition := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := collection.Find(ctx, condition, searchOption)
	if err != nil {
		return users, false
	}

	var exist, include bool
	for cursor.Next(ctx) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			return users, false
		}

		var relation Relation
		relation.UserID = ID
		relation.UserRelation = user.ID.Hex()

		include = false
		exist, err = ReadRelation(relation)
		if typeUser == "new" && exist == false {
			include = true
		}
		if typeUser == "follow" && exist == true {
			include = true
		}
		if relation.UserRelation == ID {
			include = false
		}

		if include == true {
			user.Password = ""
			user.Biography = ""
			user.WebSite = ""
			user.Location = ""
			user.Banner = ""
			user.Email = ""

			users = append(users, &user)
		}
	}

	err = cursor.Err()
	if err != nil {
		return users, false
	}
	cursor.Close(ctx)

	return users, true

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

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
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

// CheckSignIn for auth user
func CheckSignIn(email string, password string) (User, bool) {

	user, exists, _ := ExistsUser(email)

	if exists == false {
		return user, false
	}

	auxPWD := []byte(password)
	auxPWDDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(auxPWDDB, auxPWD)

	if err != nil {
		return user, false
	}

	return user, true
}

// GetUser get information of one user
func GetUser(ID string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	database := database.MongoCN.Database("gos-twitter")
	collection := database.Collection("user")

	var user User
	objectID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objectID,
	}

	err := collection.FindOne(ctx, condition).Decode(&user)
	user.Password = ""
	if err != nil {
		return user, err
	}

	return user, nil
}

// UpdateUser updated information user
func UpdateUser(user User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	database := database.MongoCN.Database("gos-twitter")
	collection := database.Collection("user")

	auxUser := make(map[string]interface{})

	if len(user.Name) > 0 {
		auxUser["nombre"] = user.Name
	}

	if len(user.LastName) > 0 {
		auxUser["lastName"] = user.LastName
	}

	auxUser["birthDate"] = user.BirthDate

	if len(user.Avatar) > 0 {
		auxUser["avatar"] = user.Avatar
	}

	if len(user.Banner) > 0 {
		auxUser["banner"] = user.Banner
	}

	if len(user.Biography) > 0 {
		auxUser["biography"] = user.Biography
	}

	if len(user.Location) > 0 {
		auxUser["location"] = user.Location
	}

	if len(user.WebSite) > 0 {
		auxUser["webSite"] = user.WebSite
	}

	updatedString := bson.M{
		"$set": auxUser,
	}
	objectID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{"_id": bson.M{"$eq": objectID}}

	_, err := collection.UpdateOne(ctx, condition, updatedString)
	if err != nil {
		return false, err
	}

	return true, nil
}
