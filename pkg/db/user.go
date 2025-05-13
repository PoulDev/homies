package db

import (
	_ "fmt"
	_ "database/sql"

	_ "github.com/PoulDev/roommates-api/config"
	_ "github.com/go-sql-driver/mysql"
)

func GetUser(id string) (User, error) {
	return User{}, nil
}

func ChangeHouse(user string, house string) error {
	return nil;
}

func GetUserHouse(user string) (House, error) {
	return House{}, nil;
}
