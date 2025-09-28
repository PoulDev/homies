package db

import (
	"github.com/lib/pq"
	"database/sql"

	"github.com/zibbadies/homies/internal/homies/db/execers"
	"github.com/zibbadies/homies/internal/homies/models"
	"github.com/zibbadies/homies/internal/homies/logger"
)


func Register(username string, password string, avatar models.Avatar) (string, *models.DBError) {
	userid, err := execers.RegisterEx(db, username, password, avatar)
	if (err == nil) {
		return userid, nil
	}

	if pqErr, ok := err.(*pq.Error); ok {
		if (pqErr.Code == "23505") {
			return "", &models.DBError{
				Message: "This username is already in use!",
				ErrorCode: models.UsernameTaken,
			}
		}
	} else {
		logger.Logger.Error("user insert error", "err", err.Error())
	}

	return "", &models.DBError{
		Message: "General error, please try again later!",
		ErrorCode: models.InternalError,
	}
}

func Login(name string, password string) (models.DBUser, *models.DBError) {
	dbuser, err := execers.LoginEx(db, name, password)

	if (err == nil) {
		return dbuser, nil
	}

	if err == sql.ErrNoRows {
		return models.DBUser{}, &models.DBError{
			Message: "Wrong username or password!",
			ErrorCode: models.WrongCredentials,
		}
	}

	logger.Logger.Error("user login error", "err", err.Error())
	return models.DBUser{}, &models.DBError{
		Message: "General error, please try again later!",
		ErrorCode: models.InternalError,
	}
}
