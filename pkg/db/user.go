package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"context"
	"time"
)

func GetUser(id string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectid, err := primitive.ObjectIDFromHex(id)
	if (err != nil) {
		return User{}, err
	}

	filter := bson.D{{"_id", objectid}}
	var dbuser dbUser

	err = users.FindOne(ctx, filter).Decode(&dbuser)
	if (err != nil) {
		return User{}, err
	}

	return User{
		UID:      dbuser.Id.Hex(),
		Username: dbuser.Username,
		Email:    dbuser.Email,
		House:    dbuser.House.Hex(),
	}, nil
}

