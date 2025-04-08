package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"context"
	"time"
)

type CartItem struct {
	Name string			`bson:"name"`
	Quantity uint16		`bson:"quantity"`
}

type Cart struct {
	Owner string		`bson:"owner"`
	Items []CartItem	`bson:"items"`
}

type dbCart struct {
	Id primitive.ObjectID		`bson:"_id,omitempty"`
	Owner primitive.ObjectID 	`bson:"owner"`
	Items []CartItem			`bson:"items"`
}

func GetCart(ownerid string) (Cart, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectid, err := primitive.ObjectIDFromHex(ownerid)
	if (err != nil) {
		return Cart{}, err
	}

	filter := bson.D{{"owner", objectid}}
	var dbcart dbCart

	err = carts.FindOne(ctx, filter).Decode(&dbcart)
	if (err != nil) {
		return Cart{}, err
	}

	return Cart{}, nil
}

func createCart(ownerid primitive.ObjectID) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var dbcart dbCart
	dbcart.Owner = ownerid
	dbcart.Items = make([]CartItem, 0)

	result, err := carts.InsertOne(ctx, dbcart)
	if (err != nil) {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
