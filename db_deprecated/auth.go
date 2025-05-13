package db_deprecated

import (
	"errors"

	"github.com/PoulDev/roommates-api/pkg/auth"
	"github.com/PoulDev/roommates-api/pkg/avatar"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"context"
	"time"
)

type User struct {
	UID string
	Username string
	Email string
	House string
	Avatar avatar.Avatar
}

type dbUser struct {
	ID primitive.ObjectID		`bson:"_id,omitempty"`
	Email string				`bson:"email"`
	Username string				`bson:"username"`
	Avatar avatar.Avatar		`bson:"avatar"`
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
		if (err == mongo.ErrNoDocuments) {
			return User{}, errors.New("No account associated with this email!")
		}
		return User{}, err
	}

	return User{
		UID:      dbuser.ID.Hex(),
		Username: dbuser.Username,
		Email:    dbuser.Email,
		House:    dbuser.House.Hex(),
		Avatar:   dbuser.Avatar,
	}, nil
}

func Register(email string, username string, password string, avatar avatar.Avatar) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	hash, salt, err := auth.HashPassword(password)
	if (err != nil) {
		return "", err
	}

	var dbuser dbUser
	dbuser.Email = email
	dbuser.Username = username
	dbuser.Avatar = avatar
	dbuser.Password = hash
	dbuser.Salt = salt

	result, err := users.InsertOne(ctx, dbuser)
	if (err != nil) {
		if (mongo.IsDuplicateKeyError(err)) {
			return "", errors.New("This email is already taken!")
		}

		return "", err
	}

	// returnato come stringa perche' solo pkg/db deve avere
	// a che fare con i metodi & tipi di MongoDB.
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
