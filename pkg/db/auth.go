package db

import (
	_ "database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/PoulDev/roommates-api/pkg/avatar"
)

type User struct {
	UID string
	Username string
	Email string
	House string
	Avatar avatar.Avatar
}

func Register(email string, username string, password string, avatar avatar.Avatar) (string, error) {
	userIns, err := db.Prepare("INSERT INTO users (name, email, pwd_hash, pwd_salt, avatar) VALUES(?, ?, ?, ?, ?)");
	if (err != nil) {
		return "", err;
	}
	defer userIns.Close();

	avatarIns, err := db.Prepare("INSERT INTO avatars (bg_color, face_color, face_x, face_y, left_eye_x, left_eye_y, right_eye_x, right_eye_y, bezier) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)");
	if (err != nil) {
		return "", err;
	}
	defer avatarIns.Close();

	avatarRes, err := avatarIns.Exec("000000", "000000", 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, "2 2 5 1 6 0");
	if (err != nil) {
		return "", err;
	}

	avatarId, err := avatarRes.LastInsertId();
	if (err != nil) {
		return "", err;
	}


	userRes, err := userIns.Exec(username, email, "c69734f24b0781ebae4adb7a137c07705d6d9f6f0a68f20973f2a5d834cd55ae", "a98a31fd4804c7cb5f0c9b64ae6c4ba8", avatarId);
	if (err != nil) {
		return "", err;
	}

	userId, err := userRes.LastInsertId();
	if (err != nil) {
		return "", err;
	}

	return fmt.Sprintf("%d", userId), nil;
}


func Login(email string, password string) (User, error) {
	return User{}, nil;
}
