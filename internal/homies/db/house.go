package db

import (
	"github.com/lib/pq"
	"database/sql"

	"github.com/zibbadies/homies/internal/homies/db/execers"
	"github.com/zibbadies/homies/internal/homies/models"
	"github.com/zibbadies/homies/internal/homies/logger"
)

// Returns the house ID & invite code
func NewHouse(name string, owner string) (string, string, error) {
	houseid, invite, err := execers.NewHouseEx(db, name, owner)
	
	if (err == nil) {
		return houseid, invite, nil
	}

	if pqErr, ok := err.(*pq.Error); ok {
		logger.Logger.Error("house insert error", "err", err.Error(), "sql_err", pqErr.Code)
	} else {
		logger.Logger.Error("house insert error", "err", err.Error())
	}

	return "", "", &models.DBError{
		Message: "General error, please try again later!",
		ErrorCode: models.InternalError,
	}
}

func GetHouse(house string, skipUser string) (models.House, error) {
	dbhouse, err := execers.GetHouseEx(db, house, skipUser)

	if (err == nil) {
		return dbhouse, nil
	}

	if err == sql.ErrNoRows {
		return models.House{}, &models.DBError{
			Message: "We didn't find your house in the database!",
			ErrorCode: models.HouseNotFound,
		}
	}

	if pqErr, ok := err.(*pq.Error); ok {
		logger.Logger.Error("house get error", "err", err.Error(), "sql_err", pqErr.Code)
	} else {
		logger.Logger.Error("house get error", "err", err.Error())
	}

	return models.House{}, &models.DBError{
		Message: "General error, please try again later!",
		ErrorCode: models.InternalError,
	}
}
