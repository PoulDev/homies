package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/PoulDev/homies/internal/homies/logger"
	"github.com/PoulDev/homies/internal/homies/models"
	_ "github.com/go-sql-driver/mysql"
)

func GetUserEx(exec Execer, id string) (models.DBUser, error) {
	var (
		username string
		avatar models.Avatar
		house sql.NullInt64
	)

	b_id, err := UUIDString2Bytes(id);
	if (err != nil) {
		logger.Logger.Error("getUser UUIDString2Bytes error", "err", err.Error(), "id", id)
		return models.DBUser{}, fmt.Errorf("There's a problem with your user, please try again later")
	}

	err = exec.QueryRow(`
		SELECT name, house, bg_color, face_color, face_x, face_y, left_eye_x, left_eye_y, right_eye_x, right_eye_y, bezier
		FROM users WHERE id = ?`, b_id).
		Scan(&username, &house, &avatar.BgColor, &avatar.FaceColor, &avatar.FaceX, &avatar.FaceY, &avatar.LeX, &avatar.LeY, &avatar.ReX, &avatar.ReY, &avatar.Bezier);

	if (err != nil) {
		logger.Logger.Error("select user error", "err", err.Error(), "id", id)
		return models.DBUser{}, fmt.Errorf("There's a problem with your user, please try again later")
	}

	houseString := "null"
	if (house.Valid) {
		houseString = fmt.Sprintf("%d", house.Int64)
	}

	return models.DBUser{
		Account: models.Account{
			User: models.User{
				UID: id,
				Username: username,
				Avatar: avatar,
			},
		},

		HouseId: houseString,
	}, nil;
}

func GetUser(id string) (models.DBUser, error) {
	return GetUserEx(db, id)
}


func ChangeHouseEx(exec Execer, user string, house string, make_owner bool) error {
	userid, err := UUIDString2Bytes(user)
	if (err != nil) { 
		logger.Logger.Error("ChangeHouse UUIDString2Bytes error", "err", err.Error(), "user", user)
		return fmt.Errorf("There's a problem with your user, please try again later")
	}

	houseid, err := strconv.Atoi(house)
	if (err != nil) { 
		logger.Logger.Error("ChangeHouse Atoi error", "err", err.Error(), "user", user)
		return fmt.Errorf("There's a problem with your house, please try again later")
	}

	if (make_owner) {
		_, err = exec.Exec("UPDATE users SET house = ?, is_owner = TRUE WHERE id = ?", houseid, userid)
	} else {
		_, err = exec.Exec("UPDATE users SET house = ? WHERE id = ?", houseid, userid)
	}

	if (err != nil) {
		logger.Logger.Error("ChangeHouse update error", "err", err.Error(), "user", user)
		return fmt.Errorf("Internal error, please try again later")
	}

	return nil;
}

func ChangeHouse(user string, house string, make_owner bool) error {
	return ChangeHouseEx(db, user, house, make_owner)
}

func HouseIDByInvite(invite string) (string, error) {
	var houseid int64
	
	err := db.QueryRow(`SELECT id FROM houses WHERE invite = ?`, invite).Scan(&houseid)
	if (err != nil) {
		logger.Logger.Error("FindHouseByInvite error", "err", err.Error(), "invite", invite)
		return "", fmt.Errorf("Internal error, please try again later")
	}
	
	return strconv.FormatInt(houseid, 10), nil;
}


func MakeHouseOwnerEx(exec Execer, user string, owner bool) error {
	userid, err := UUIDString2Bytes(user)
	if (err != nil) {
		logger.Logger.Error("MakeHouseOwner UUIDString2Bytes error", "err", err.Error(), "user", user)
		return fmt.Errorf("There's a problem with your user, please try again later")
	}

	_, err = exec.Exec("UPDATE users SET is_owner = ? WHERE id = ?", owner, userid)
	if (err != nil) {
		logger.Logger.Error("MakeHouseOwner update error", "err", err.Error(), "user", user)
		return fmt.Errorf("Internal error, please try again later")
	}

	return nil;
}

func MakeHouseOwner(user string, owner bool) error {
	return MakeHouseOwnerEx(db, user, owner)
}

// TODO: Use a JOIN query?
// good: more performant, faster
// bad:  If I need to update the house table, I'll have to update GetUserHouseEx and GetHouseEx
func GetUserHouseEx(exec Execer, user string) (models.House, error) {
	var (
		tmp_houseid sql.NullInt64
		houseid int64
	)

	b_id, err := UUIDString2Bytes(user);
	if (err != nil) {
		logger.Logger.Error("UUIDString2Bytes error", "err", err.Error())
		return models.House{}, fmt.Errorf("There's a problem with your user, please try again later")
	}

	err = exec.QueryRow(`SELECT house FROM users WHERE id = ?`, b_id).Scan(&tmp_houseid);

	if (err != nil) {
		logger.Logger.Error("user house ID retrival error", "err", err.Error())
		return models.House{}, fmt.Errorf("Internal error, please try again later")
	}

	if (!tmp_houseid.Valid) {
		return models.House{}, errors.New("You don't have an house")
	} else {
		houseid = tmp_houseid.Int64
	}
	
	house, err := GetHouse(strconv.FormatInt(houseid, 10), b_id)

	if (err != nil) {
		logger.Logger.Error("user house name retrival error", "err", err.Error())
		return models.House{}, fmt.Errorf("Internal error, please try again later")
	}

	return house, nil;
}

func GetUserHouse(user string) (models.House, error) {
	return GetUserHouseEx(db, user)
}

