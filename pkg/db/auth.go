package db

import (
	"github.com/PoulDev/roommates-api/pkg/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"context"
	"time"
)

type User struct {
	UID string
	Username string
	Email string
	House string
}

type dbUser struct {
	Id primitive.ObjectID		`bson:"_id,omitempty"`
	Email string				`bson:"email"`
	Username string				`bson:"username"`
	Password []byte				`bson:"pwd"`
	Salt []byte					`bson:"salt"`
	House primitive.ObjectID	`bson:"house"`
}

func Login(email string, password string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"email", email}}
	var dbuser dbUser

	err := users.FindOne(ctx, filter).Decode(&dbuser)
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

func Register(email string, username string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	hash, salt, err := auth.HashPassword(password)
	if (err != nil) {
		return "", err
	}

	var dbuser dbUser
	dbuser.Email = email
	dbuser.Username = username
	dbuser.Password = hash
	dbuser.Salt = salt

	result, err := users.InsertOne(ctx, dbuser)
	if (err != nil) {
		return "", err
	}

	// returnato come stringa perche' solo pkg/db deve avere
	// a che fare con i metodi & tipi di MongoDB.
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
