package db

import (
	"database/sql"
	
	"github.com/lib/pq"

	"github.com/zibbadies/homies/internal/homies/logger"
	"github.com/zibbadies/homies/internal/homies/models"
	"github.com/zibbadies/homies/internal/homies/db/execers"
)

func GetUser(id string) (models.DBUser, error) {
	dbuser, err := execers.GetUserEx(db, id)
	if (err == nil) {
		return dbuser, nil
	}

	if (err == sql.ErrNoRows) {
		return models.DBUser{}, &models.DBError{
			Message: "We didn't find your user in the database!",
			ErrorCode: models.UserNotFound,
		}
	}

	if pqErr, ok := err.(*pq.Error); ok {
		logger.Logger.Error("user get error", "err", err.Error(), "id", id, "sql_err", pqErr.Code)
	} else {
		logger.Logger.Error("user get error", "err", err.Error(), "id", id)
	}

	return models.DBUser{}, &models.DBError{
		Message: "There's a problem with your user, please try again later",
		ErrorCode: "internal_error",
	}
}

// Same as GetUser, but if the user is not found
// it returns a DBError NotAuthenticated instead of UserNotFound
func GetUserMe(user string) (models.DBUser, error) {
	dbuser, err := GetUser(user)
	if (err != nil) {
		if dberr, ok := err.(models.DBError); ok {
			if dberr.ErrorCode == models.UserNotFound {
				return models.DBUser{}, &models.DBError{
					Message: "Your user was not found in the database!",
					ErrorCode: models.NotAuthenticated,
				}
			}
		}
		return models.DBUser{}, err
	}
	return dbuser, nil
}

func ChangeHouse(user string, house string) error {
	err := execers.ChangeHouseEx(db, user, house)

	if err == nil {
		return nil
	}

	if pqErr, ok := err.(*pq.Error); ok {
		logger.Logger.Error("change house error", "err", err.Error(), "sql_err", pqErr.Code)
	} else {
		logger.Logger.Error("change house error", "err", err.Error())
	}

	return &models.DBError{
		Message:   "General error, please try again later!",
		ErrorCode: models.InternalError,
	}
}

func LeaveHouse(user string) error {
	err := execers.LeaveHouseEx(db, user)

	if err == nil {
		return nil
	}

	if pqErr, ok := err.(*pq.Error); ok {
		logger.Logger.Error("leave house error", "err", err.Error(), "sql_err", pqErr.Code)
	} else {
		logger.Logger.Error("leave house error", "err", err.Error())
	}

	return &models.DBError{
		Message:   "General error, please try again later!",
		ErrorCode: models.InternalError,
	}
}

func HouseIDByInvite(invite string) (string, error) {
	houseID, err := execers.HouseIDByInviteEx(db, invite)

	if err == nil {
		return houseID, nil
	}

	if err == sql.ErrNoRows {
		return "", &models.DBError{
			Message:   "Invalid or expired invite code!",
			ErrorCode: models.InviteNotFound,
		}
	}

	if pqErr, ok := err.(*pq.Error); ok {
		logger.Logger.Error("house by invite error", "err", err.Error(), "sql_err", pqErr.Code)
	} else {
		logger.Logger.Error("house by invite error", "err", err.Error())
	}

	return "", &models.DBError{
		Message:   "General error, please try again later!",
		ErrorCode: models.InternalError,
	}
}

func MakeHouseOwner(house string, user string) error {
	err := execers.MakeHouseOwnerEx(db, house, user)

	if err == nil {
		return nil
	}

	if pqErr, ok := err.(*pq.Error); ok {
		logger.Logger.Error("make house owner error", "err", err.Error(), "sql_err", pqErr.Code)
	} else {
		logger.Logger.Error("make house owner error", "err", err.Error())
	}

	return &models.DBError{
		Message:   "General error, please try again later!",
		ErrorCode: models.InternalError,
	}
}

func GetUserHouse(user string) (models.House, error) {
	house, err := execers.GetUserHouseEx(db, user)

	if err == nil {
		return house, nil
	}

	if err.Error() == models.UserNotInHouse {
		return models.House{}, &models.DBError{
			Message:   "This user is not in a house!",
			ErrorCode: models.UserNotInHouse,
		}
	}

	if err.Error() == models.UserNotFound {
		return models.House{}, &models.DBError{
			Message:   "User not found in the database!",
			ErrorCode: models.UserNotFound,
		}
	}

	if err == sql.ErrNoRows {
		return models.House{}, &models.DBError{
			Message: "House not found in the database!",
			ErrorCode: models.HouseNotFound,
		}
	}

	if pqErr, ok := err.(*pq.Error); ok {
		logger.Logger.Error("get user house error", "err", err.Error(), "sql_err", pqErr.Code)
	} else {
		logger.Logger.Error("get user house error", "err", err.Error())
	}

	return models.House{}, &models.DBError{
		Message:   "General error, please try again later!",
		ErrorCode: models.InternalError,
	}
}

func SetAvatar(user string, avatar models.Avatar) error {
	err := execers.SetAvatarEx(db, user, avatar)

	if err == nil {
		return nil
	}

	if pqErr, ok := err.(*pq.Error); ok {
		logger.Logger.Error("set avatar error", "err", err.Error(), "sql_err", pqErr.Code)
	} else {
		logger.Logger.Error("set avatar error", "err", err.Error())
	}

	return &models.DBError{
		Message:   "General error, please try again later!",
		ErrorCode: models.InternalError,
	}
}

