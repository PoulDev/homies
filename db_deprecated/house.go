package db_deprecated

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type dbHouse struct {
	ID primitive.ObjectID			`bson:"_id,omitempty"`
	Name string 					`bson:"name"`
	Owner primitive.ObjectID		`bson:"owner"`
	Members []primitive.ObjectID 	`bson:"members"` // owner omitted 
}

type House struct {
	Name string
	Owner string
	Members []string
}

func newHouse(name string, owner primitive.ObjectID) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var dbhouse dbHouse;
	dbhouse.Name = name
	dbhouse.Owner = owner
	dbhouse.Members = []primitive.ObjectID{}

	result, err := houses.InsertOne(ctx, dbhouse)
	if (err != nil) {
		if (mongo.IsDuplicateKeyError(err)) {
			return "", errors.New("You already have an house!")
		}
		return "", err
	}
	changeHouse(owner, result.InsertedID.(primitive.ObjectID))

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func NewHouse(name string, owner string) (string, error) {
	ownerid, err := primitive.ObjectIDFromHex(owner)
	if (err != nil) {
		return "", err
	}

	return newHouse(name, ownerid)
}

func getHouse(house primitive.ObjectID) (House, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var dbhouse dbHouse

	err := houses.FindOne(ctx, bson.D{{"_id", house}}).Decode(&dbhouse)
	if (err != nil) {
		return House{}, err
	}

	var members []string
	for _, id := range dbhouse.Members {
		members = append(members, id.Hex())
	}

	return House{
		dbhouse.Name,
		dbhouse.Owner.Hex(),
		members,
	}, nil
}

func GetHouse(house string) (House, error) {
	houseid, err := primitive.ObjectIDFromHex(house)
	if (err != nil) {
		return House{}, err
	}
	return getHouse(houseid)
}

