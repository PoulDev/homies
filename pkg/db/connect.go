package db

import (
	"fmt"
	"database/sql"

	"github.com/PoulDev/roommates-api/config"
	_ "github.com/go-sql-driver/mysql"
)

// Internal Variables
var (
	db *sql.DB
)

func ConnectDatabase() error {
	newdb, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBName))
	db = newdb;
	if (err != nil) {
		return err;
	}

	err = db.Ping()
	if (err != nil) {
		return err;
	}

	return nil;
}

func CheckConnection() error {
	return db.Ping();
}
