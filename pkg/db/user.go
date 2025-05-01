package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"context"
	"time"
)

func getUser(id primitive.ObjectID) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var dbuser dbUser

	err := users.FindOne(ctx, bson.D{{"_id", id}}).Decode(&dbuser)
	if (err != nil) {
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

func GetUser(id string) (User, error) {
	objectid, err := primitive.ObjectIDFromHex(id)
	if (err != nil) {
		return User{}, err
	}

	return getUser(objectid)
}


func changeHouse(user primitive.ObjectID, house primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := users.UpdateOne(ctx,
		bson.M{"_id": user},
		bson.M{"$set": bson.M{"house": house}},
	)
	if (err != nil) {
		return err
	}
	return nil
}

func ChangeHouse(user string, house string) error {
	userid, err := primitive.ObjectIDFromHex(user)
	if (err != nil) {
		return err
	}

	houseid, err := primitive.ObjectIDFromHex(house)
	if (err != nil) {
		return err
	}

	return changeHouse(userid, houseid)
}


func getUserHouse(user primitive.ObjectID) (House, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var dbuser dbUser
	
	err := users.FindOne(ctx, bson.D{{"_id", user}}).Decode(&dbuser)
	if (err != nil) {
		return House{}, err
	}

	return getHouse(dbuser.House)
}

func GetUserHouse(user string) (House, error) {
	userid, err := primitive.ObjectIDFromHex(user)
	if (err != nil) {
		return House{}, err
	}

	return getUserHouse(userid)
}
