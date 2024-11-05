package model

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	ID    bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email string        `bson:"email,omitempty" json:"email,omitempty"`
	Name  string        `bson:"name,omitempty" json:"name,omitempty"`
}
