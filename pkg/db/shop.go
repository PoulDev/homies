package db

import (
	_ "fmt"
	_ "database/sql"

	_ "github.com/PoulDev/roommates-api/config"
	_ "github.com/go-sql-driver/mysql"
)


type CartItem struct {
	Name string			`bson:"name"`
	Quantity uint16		`bson:"quantity"`
}

type Cart struct {
	Owner string		`bson:"owner"`
	Items []CartItem	`bson:"items"`
}

func GetCart(ownerid string) (Cart, error) {
	return Cart{}, nil
}

