package db

import (
	_ "fmt"
	_ "database/sql"

	_ "github.com/PoulDev/roommates-api/config"
	_ "github.com/go-sql-driver/mysql"
)

type House struct {
	Name string
	Owner string
	Members []string
}


func NewHouse(name string, owner string) (string, error) {
	return "", nil
}

func GetHouse(house string) (House, error) {
	return House{}, nil
}

