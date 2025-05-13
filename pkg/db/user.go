package db

import (
	"database/sql"
	_ "database/sql"
	"errors"
	"fmt"
	_ "fmt"

	_ "github.com/PoulDev/roommates-api/config"
	"github.com/PoulDev/roommates-api/pkg/avatar"
	_ "github.com/go-sql-driver/mysql"
)

func GetUser(id string) (User, error) {
	var (
		ID int
		username string
		dbemail string
		house sql.NullInt64
	)

	err := db.QueryRow(`
		SELECT id, name, email, house
		FROM users WHERE id = ?`, id).
		Scan(&ID, &username, &dbemail, &house);

	if (err != nil) {
		return User{}, err
	}

	houseString := "null"
	if (house.Valid) {
		houseString = fmt.Sprintf("%d", house.Int64)
	}

	return User{
		UID: fmt.Sprintf("%d", ID),
		Username: username,
		Email: dbemail,
		House: houseString,
	}, nil;
}

func ChangeHouse(user string, house string) error { // !! TODO
	return nil;
}

func GetUserHouse(user string) (House, error) {
	var (
		tmp_houseid sql.NullInt64
		houseid int64
		name string
		owner_id int64
	)
	err := db.QueryRow(`SELECT house FROM users WHERE id = ?`, user).Scan(&tmp_houseid);

	if (err != nil) {
		return House{}, err
	}

	if (!tmp_houseid.Valid) {
		return House{}, errors.New("User doesn't have an house")
	} else {
		houseid = tmp_houseid.Int64
	}

	err = db.QueryRow("SELECT name, owner_id FROM houses WHERE id = ?", houseid).Scan(&name, &owner_id)

	if (err != nil) {
		return House{}, err
	}

	return House{
		Name: name,
		Owner: fmt.Sprintf("%d", owner_id),
	}, nil;
}

func GetUserAvatar(avatarid string) (avatar.Avatar, error) {
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
