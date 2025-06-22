package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/PoulDev/roommates-api/pkg/avatar"
	"github.com/PoulDev/roommates-api/pkg/logger"
	_ "github.com/go-sql-driver/mysql"
)

func GetUserEx(exec Execer, id string) (User, error) {
	var (
		username string
		dbemail string
		avatar int64
		house sql.NullInt64
	)

	b_id, err := UUIDString2Bytes(id);
	if (err != nil) {
		logger.Logger.Error("getUser UUIDString2Bytes error", "err", err.Error(), "id", id)
		return User{}, fmt.Errorf("There's a problem with your user, please try again later")
	}

	err = exec.QueryRow(`
		SELECT name, email, house, avatar
		FROM users WHERE id = ?`, b_id).
		Scan(&username, &dbemail, &house, &avatar);

	if (err != nil) {
		logger.Logger.Error("select user error", "err", err.Error(), "id", id)
		return User{}, fmt.Errorf("There's a problem with your user, please try again later")
	}

	houseString := "null"
	if (house.Valid) {
		houseString = fmt.Sprintf("%d", house.Int64)
	}

	return User{
		UID: id,
		Username: username,
		Avatar: fmt.Sprintf("%d", avatar),
		Email: dbemail,
		House: houseString,
	}, nil;
}

func GetUser(id string) (User, error) {
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

func GetUserHouseEx(exec Execer, user string) (House, error) {
	var (
		tmp_houseid sql.NullInt64
		houseid int64
		name string
	)

	b_id, err := UUIDString2Bytes(user);
	if (err != nil) {
		logger.Logger.Error("UUIDString2Bytes error", "err", err.Error())
		return House{}, fmt.Errorf("There's a problem with your user, please try again later")
	}

	err = exec.QueryRow(`SELECT house FROM users WHERE id = ?`, b_id).Scan(&tmp_houseid);

	if (err != nil) {
		logger.Logger.Error("user house ID retrival error", "err", err.Error())
		return House{}, fmt.Errorf("Internal error, please try again later")
	}

	if (!tmp_houseid.Valid) {
		return House{}, errors.New("You don't have an house")
	} else {
		houseid = tmp_houseid.Int64
	}

	err = exec.QueryRow("SELECT name FROM houses WHERE id = ?", houseid).Scan(&name)
	if (err != nil) {
		logger.Logger.Error("user house name retrival error", "err", err.Error())
		return House{}, fmt.Errorf("Internal error, please try again later")
	}

	return House{
		Name: name,
	}, nil;
}

func GetUserHouse(user string) (House, error) {
	return GetUserHouseEx(db, user)
}

func GetAvatarEx(exec Execer, avatarid string) (avatar.Avatar, error) {
	var (
		bg_color string
		face_color string
		face_x float32
		face_y float32
		left_eye_x float32
		left_eye_y float32
		right_eye_x float32
		right_eye_y float32
		bezier string
	)

	err := exec.QueryRow(`
		SELECT bg_color, face_color, face_x, face_y, left_eye_x, left_eye_y, right_eye_x, right_eye_y, bezier 
		FROM avatars WHERE id = ?`, avatarid).
			Scan(&bg_color, &face_color, &face_x, &face_y, &left_eye_x, &left_eye_y, &right_eye_x, &right_eye_y, &bezier);
	
	if (err != nil) {
		logger.Logger.Error("user avatar retrival error", "err", err.Error())
		return avatar.Avatar{}, fmt.Errorf("Internal error, please try again later")
	}

	return avatar.Avatar{
		BgColor: bg_color,
		FaceColor: face_color,
		FaceX: face_x,
		FaceY: face_y,
		LeX: left_eye_x,
		LeY: left_eye_y,
		ReX: right_eye_x,
		ReY: right_eye_y,
		Bezier: bezier,
	}, nil
}

func GetAvatar(avatarid string) (avatar.Avatar, error) {
	return GetAvatarEx(db, avatarid)
}
