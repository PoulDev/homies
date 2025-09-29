package db

import (
	"database/sql"
	"fmt"

	"github.com/zibbadies/homies/internal/homies/config"
	_ "github.com/lib/pq"
)

// Internal Variables
var (
	db *sql.DB
)

func ConnectDatabase() error {
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	newdb, err := sql.Open("postgres", psqlconn)
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
