package db

import (
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	_ "context"
	_ "time"
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
	ID primitive.ObjectID		`bson:"_id,omitempty"`
	Owner primitive.ObjectID 	`bson:"owner"`
	Items []CartItem			`bson:"items"`
}

func GetCart(ownerid string) (Cart, error) {


	return Cart{}, nil
}

func createCart(ownerid primitive.ObjectID) (string, error) {


	return "", nil
}
