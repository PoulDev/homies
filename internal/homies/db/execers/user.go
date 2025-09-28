package execers

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/zibbadies/homies/internal/homies/logger"
	"github.com/zibbadies/homies/internal/homies/models"
)

type UserError struct {
	Message string
	HttpCode int
	HomiesErrorCode int
}

func (e UserError) Error() string {
	return e.Message
}

func GetUserEx(exec Execer, id string) (models.DBUser, error) {
	var (
		username string
		avatar models.Avatar
		house sql.NullInt64
	)

	err := exec.QueryRow(`
		SELECT name, house, bg_color, face_color, face_x, face_y, left_eye_x, left_eye_y, right_eye_x, right_eye_y, bezier
		FROM users WHERE id = $1`, id).
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

func ChangeHouseEx(exec Execer, userId string, house string) error {
	houseid, err := strconv.Atoi(house)
	if (err != nil) { 
		logger.Logger.Error("ChangeHouse Atoi error", "err", err.Error(), "user", userId)
		return fmt.Errorf("There's a problem with your house, please try again later")
	}


	_, err = exec.Exec("UPDATE users SET house = $1 WHERE id = $2", houseid, userId)

	if (err != nil) {
		logger.Logger.Error("ChangeHouse update error", "err", err.Error(), "user", userId)
		return fmt.Errorf("Internal error, please try again later")
	}

	return nil;
}

func LeaveHouseEx(exec Execer, userId string) error {
	_, err := exec.Exec("UPDATE users SET house = NULL WHERE id = $1", userId)
	if (err != nil) {
		logger.Logger.Error("LeaveHouse update error", "err", err.Error(), "user", userId)
		return fmt.Errorf("Internal error, please try again later")
	}

	return nil;
}

func HouseIDByInviteEx(exec Execer, invite string) (string, error) {
	var houseid int64
	
	err := exec.QueryRow(`SELECT id FROM houses WHERE invite = $1`, invite).Scan(&houseid)
	if (err != nil) {
		logger.Logger.Error("FindHouseByInvite error", "err", err.Error(), "invite", invite)
		return "", fmt.Errorf("Internal error, please try again later")
	}
	
	return strconv.FormatInt(houseid, 10), nil;
}


func MakeHouseOwnerEx(exec Execer, houseId string, userId string) error {
	_, err := exec.Exec("UPDATE houses SET owner = $1 WHERE id = $2", userId, houseId)
	if (err != nil) {
		logger.Logger.Error("MakeHouseOwner update error", "err", err.Error(), "user", userId)
		return fmt.Errorf("Internal error, please try again later")
	}

	return nil;
}

func GetUserHouseEx(exec Execer, user string) (models.House, error) {
	var (
		tmp_houseid sql.NullInt64
		houseid int64
	)

	err := exec.QueryRow(`SELECT house FROM users WHERE id = $1`, user).Scan(&tmp_houseid);

	if (err != nil) {
		logger.Logger.Error("user house ID retrival error", "err", err.Error())
		return models.House{}, fmt.Errorf("Internal error, please try again later")
	}

	if (!tmp_houseid.Valid) {
		return models.House{}, errors.New("no_house")
	} else {
		houseid = tmp_houseid.Int64
	}

	// TODO: Use a JOIN query?
	// good: more performant, faster
	// bad:  If I need to update the house table, I'll have to update GetUserHouseEx and GetHouseEx
	house, err := GetHouseEx(exec, strconv.FormatInt(houseid, 10), user)

	if (err != nil) {
		logger.Logger.Error("user house name retrival error", "err", err.Error())
		return models.House{}, fmt.Errorf("Internal error, please try again later")
	}

	return house, nil;
}

func SetAvatarEx(exec Execer, userId string, avatar models.Avatar) error {
	_, err := exec.Exec("UPDATE users SET bg_color = $1, face_color = $2, face_x = $3, face_y = $4, left_eye_x = $5, left_eye_y = $6, right_eye_x = $7, right_eye_y = $8, bezier = $9 WHERE id = $10", avatar.BgColor, avatar.FaceColor, avatar.FaceX, avatar.FaceY, avatar.LeX, avatar.LeY, avatar.ReX, avatar.ReY, avatar.Bezier, userId)
	if (err != nil) {
		logger.Logger.Error("SetAvatar update error", "err", err.Error(), "user", userId)
		return fmt.Errorf("Internal error, please try again later")
	}

	return nil;
}

