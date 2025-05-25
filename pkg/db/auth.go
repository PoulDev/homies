package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/PoulDev/roommates-api/pkg/auth"
	"github.com/PoulDev/roommates-api/pkg/avatar"
	"github.com/google/uuid"
)

type User struct {
	UID string
	Username string
	Email string
	House string
	Avatar string
}

func Register(email string, username string, password string, avatar avatar.Avatar) (string, error) {
	hash, salt, err := auth.HashPassword(password)
	if err != nil {
		return "", err
	}

	avatarRes, err := db.Exec(`
		INSERT INTO avatars (
			bg_color, face_color, face_x, face_y,
			left_eye_x, left_eye_y, right_eye_x, right_eye_y, bezier
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		avatar.BgColor, avatar.FaceColor, avatar.FaceX, avatar.FaceY, avatar.LeX, avatar.LeY, avatar.ReX, avatar.ReY, avatar.Bezier,
	)
	if err != nil {
		return "", err
	}

	avatarId, err := avatarRes.LastInsertId()
	if err != nil {
		return "", err
	}

	userId := uuid.New();

	_, err = db.Exec(`
		INSERT INTO users (id, name, email, pwd_hash, pwd_salt, avatar)
		VALUES (?, ?, ?, ?, ?, ?)`,
		UUID2Bytes(userId), username, email, hash, salt, avatarId,
	)
	if err != nil {
		return "", err
	}

	return userId.String(), nil
}


func Login(email string, password string) (User, error) {
	var (
		ID []byte
		username string
		dbemail string
		house sql.NullInt64
		pwdHash []byte
		pwdSalt []byte
	)

	err := db.QueryRow("SELECT id, name, email, house, pwd_hash, pwd_salt FROM users WHERE email = ?", email).Scan(&ID, &username, &dbemail, &house, &pwdHash, &pwdSalt)
	if (err != nil) {
		return User{}, err;
	}

	if (!auth.CheckPassword(password, pwdHash, pwdSalt)) {
		return User{}, errors.New("password mismatch");
	}

	var houseString string;
	if (house.Valid) {
		houseString = fmt.Sprintf("%d", house.Int64)
	} else {
		houseString = "null";
	}

	uid, err := UUIDBytes2String(ID)
	if (err != nil) {
		return User{}, err;
	}

	return User{
		UID: uid,
		Username: username,
		Email: dbemail,
		House: houseString,
	}, nil;
}
