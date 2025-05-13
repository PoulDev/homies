package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/PoulDev/roommates-api/pkg/auth"
	"github.com/PoulDev/roommates-api/pkg/avatar"
)

type User struct {
	UID string
	Username string
	Email string
	House string
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

	userRes, err := db.Exec(`
		INSERT INTO users (name, email, pwd_hash, pwd_salt, avatar)
		VALUES (?, ?, ?, ?, ?)`,
		username, email, hash, salt, avatarId,
	)
	if err != nil {
		return "", err
	}

	userId, err := userRes.LastInsertId()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", userId), nil
}


func Login(email string, password string) (User, error) {
	var (
		ID int
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
		return User{}, errors.New("Password mismatch");
	}

	log.Println("Login Successful");

	var houseString string;
	if (house.Valid) {
		houseString = fmt.Sprintf("%d", house.Int64)
	} else {
		houseString = "null";
	}

	return User{
		UID: fmt.Sprintf("%d", ID),
		Username: username,
		Email: dbemail,
		House: houseString,
	}, nil;
}
