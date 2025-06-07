package db

import (
	"database/sql"
	"fmt"

	"github.com/PoulDev/roommates-api/config"
	_ "github.com/go-sql-driver/mysql"
)

// Internal Variables
var (
	db *sql.DB
)

type Execer interface {
    Exec(query string, args ...any) (sql.Result, error)
    Query(query string, args ...any) (*sql.Rows, error)
    QueryRow(query string, args ...any) *sql.Row
}

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
