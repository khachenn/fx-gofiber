package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email string             `bson:"email,omitempty" json:"email,omitempty"`
	Name  string             `bson:"name,omitempty" json:"name,omitempty"`
}
