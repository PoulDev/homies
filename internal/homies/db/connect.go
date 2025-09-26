package db

import (
	"database/sql"
	"fmt"

	"github.com/PoulDev/homies/internal/homies/config"
	_ "github.com/lib/pq"
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
