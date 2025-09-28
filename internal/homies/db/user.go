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

func ChangeHouse(user string, house string) error {
	return execers.ChangeHouseEx(db, user, house)
}

func LeaveHouse(user string) error {
	return execers.LeaveHouseEx(db, user)
}

func HouseIDByInvite(invite string) (string, error) {
	return execers.HouseIDByInviteEx(db, invite)
}

func MakeHouseOwner(house string, user string, owner bool) error {
	return execers.MakeHouseOwnerEx(db, house, user, owner)
}

func GetUserHouse(user string) (models.House, error) {
	return execers.GetUserHouseEx(db, user)
}

func SetAvatar(user string, avatar models.Avatar) error {
	return execers.SetAvatarEx(db, user, avatar)
}
