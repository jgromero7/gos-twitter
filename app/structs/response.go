package structs

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ReponseSignIn returns data SignIn
type ReponseSignIn struct {
	Token string `json:"token,omitempty"`
}

// ResponseQueryRelation true or false
type ResponseQueryRelation struct {
	Status bool `json:"status"`
}

// ResponseTweetsFollowers all tweets
type ResponseTweetsFollowers struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuerID      string             `bson:"userid" json:"userid,omitempty"`
	UserRelation string             `bson:"userrelation" json:"userrelation"`
	Tweet        struct {
		ID       string    `bson:"_id" json:"id,omitempty"`
		Message  string    `bson:"message" json:"message,omitempty"`
		CreateAt time.Time `bson:"createAt" json:"createAt,omitempty"`
	}
}
