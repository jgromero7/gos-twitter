package models

import (
	"context"
	"time"

	"github.com/jgromero7/gos-twitter/database"
	"go.mongodb.org/mongo-driver/bson"
)

// Relation of users
type Relation struct {
	UserID       string `bson:"userid" json:"userid"`
	UserRelation string `bson:"userrelation" json:"userrelation"`
}

// ReadRelation for users
func ReadRelation(relation Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := database.MongoCN.Database("gos-twitter")
	collection := database.Collection("relation")

	condition := bson.M{
		"userid":       relation.UserID,
		"userrelation": relation.UserRelation,
	}

	var realtions Relation
	err := collection.FindOne(ctx, condition).Decode(&realtions)
	if err != nil {
		return false, err
	}

	return true, nil
}

// CreateRelation for users
func CreateRelation(relation Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := database.MongoCN.Database("gos-twitter")
	collection := database.Collection("relation")

	_, err := collection.InsertOne(ctx, relation)
	if err != nil {
		return false, err
	}

	return true, nil
}

// DeleteRelation for users
func DeleteRelation(relation Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := database.MongoCN.Database("gos-twitter")
	collection := database.Collection("relation")

	condition := bson.M{
		"userid":       relation.UserID,
		"userrelation": relation.UserRelation,
	}

	_, err := collection.DeleteOne(ctx, condition)
	if err != nil {
		return false, err
	}

	return true, nil
}
