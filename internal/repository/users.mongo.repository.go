package repository

import (
	"context"
	"fmt"
	"sturdy-winner-api/model"

	"go.mongodb.org/mongo-driver/bson"
)

func (h Handler) Users() ([]model.User, error) {
	var users []model.User
	filter := bson.M{}
	cursor, err := h.mongo.DB.Collection("users").Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	fmt.Println(users)
	return users, nil
}
