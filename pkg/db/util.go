package db

import (
	_ "fmt"
	_ "database/sql"

	_ "github.com/PoulDev/roommates-api/config"
	_ "github.com/go-sql-driver/mysql"
)

func IdOrnull(ID string) string { // !! TODO
	if (ID == "000000000000000000000000") {
		return "null"
	}
	return ID
}

func PreattyError(err error) error { // !! TODO
	return err
}
