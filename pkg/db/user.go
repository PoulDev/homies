package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/PoulDev/roommates-api/pkg/avatar"
	_ "github.com/go-sql-driver/mysql"
)

func GetUser(id string) (User, error) {
	var (
		username string
		dbemail string
		avatar int64
		house sql.NullInt64
	)

	b_id, err := UUIDString2Bytes(id);
	if (err != nil) { return User{}, err; }

	err = db.QueryRow(`
		SELECT name, email, house, avatar
		FROM users WHERE id = ?`, b_id).
		Scan(&username, &dbemail, &house, &avatar);

	if (err != nil) {
		return User{}, err
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

func ChangeHouse(user string, house string, make_owner bool) error {
	userid, err := UUIDString2Bytes(user)
	if (err != nil) { return err }

	houseid, err := strconv.Atoi(house)
	if (err != nil) { return err }

	if (make_owner) {
		_, err = db.Exec("UPDATE users SET house = ?, is_owner = TRUE WHERE id = ?", houseid, userid)
	} else {
		_, err = db.Exec("UPDATE users SET house = ? WHERE id = ?", houseid, userid)
	}

	log.Println("Making admin", user, userid)

	if (err != nil) { return err }

	return nil;
}

func MakeHouseOwner(user string, owner bool) error {
	userid, err := UUIDString2Bytes(user)
	if (err != nil) { return err }

	_, err = db.Exec("UPDATE users SET is_owner = ? WHERE id = ?", owner, userid)
	if (err != nil) { return err }

	return nil;
}

func GetUserHouse(user string) (House, error) {
	var (
		tmp_houseid sql.NullInt64
		houseid int64
		name string
	)

	b_id, err := UUIDString2Bytes(user);
	if (err != nil) { return House{}, err; }

	err = db.QueryRow(`SELECT house FROM users WHERE id = ?`, b_id).Scan(&tmp_houseid);

	if (err != nil) {
		return House{}, err
	}

	if (!tmp_houseid.Valid) {
		return House{}, errors.New("User doesn't have an house")
	} else {
		houseid = tmp_houseid.Int64
	}

	err = db.QueryRow("SELECT name FROM houses WHERE id = ?", houseid).Scan(&name)

	if (err != nil) {
		return House{}, err
	}

	return House{
		Name: name,
	}, nil;
}

func GetAvatar(avatarid string) (avatar.Avatar, error) {
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

	err := db.QueryRow(`
		SELECT bg_color, face_color, face_x, face_y, left_eye_x, left_eye_y, right_eye_x, right_eye_y, bezier 
		FROM avatars WHERE id = ?`, avatarid).
		Scan(&bg_color, &face_color, &face_x, &face_y, &left_eye_x, &left_eye_y, &right_eye_x, &right_eye_y, &bezier);
	
	if (err != nil) {
		return avatar.Avatar{}, err;
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
