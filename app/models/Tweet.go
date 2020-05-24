package models

import (
	"context"
	"time"

	"github.com/jgromero7/gos-twitter/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Tweet model tweets
type Tweet struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserID   string             `bson:"userid" json:"userid,omitempty"`
	Message  string             `bson:"message" json:"message,omitempty"`
	CreateAt time.Time          `bson:"createAt" json:"createAt,omitempty"`
}

// Create stores a record
func Create(tweet Tweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := database.MongoCN.Database("gos-twitter")
	collection := database.Collection("tweet")

	auxTweet := bson.M{
		"userid":    tweet.UserID,
		"message":   tweet.Message,
		"createdAt": tweet.CreateAt,
	}

	result, err := collection.InsertOne(ctx, auxTweet)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}

// Read roturn all tweets
func Read(userID string, page int64) ([]*Tweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := database.MongoCN.Database("gos-twitter")
	collection := database.Collection("tweet")

	var tweets []*Tweet

	condition := bson.M{
		"userid": userID,
	}

	searchOption := options.Find()
	searchOption.SetLimit(20)
	searchOption.SetSort(bson.D{{Key: "creatAt", Value: -1}})
	searchOption.SetSkip((page - 1) * 20)

	cursor, err := collection.Find(ctx, condition, searchOption)
	if err != nil {
		return tweets, false
	}

	for cursor.Next(context.TODO()) {
		var tweet Tweet
		err := cursor.Decode(&tweet)
		if err != nil {
			return tweets, false
		}

		tweets = append(tweets, &tweet)
	}

	return tweets, true
}

// Delete delete a register tweet
func Delete(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := database.MongoCN.Database("gos-twitter")
	collection := database.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	_, err := collection.DeleteOne(ctx, condition)

	return err
}
