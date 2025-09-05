package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/PoulDev/homies/pkg/homies/auth"
	"github.com/PoulDev/homies/internal/homies/logger"
	"github.com/PoulDev/homies/internal/homies/models"
	"github.com/google/uuid"
	"github.com/go-sql-driver/mysql"
)


func RegisterEx(exec Execer, username string, password string, avatar models.Avatar) (string, error) {
	hash, salt, err := auth.HashPassword(password)
	if err != nil {
		return "", err
	}

	var mysqlErr *mysql.MySQLError
	userId := uuid.New();
	_, err = exec.Exec(`
		INSERT INTO users (
			id, name, pwd_hash, pwd_salt,
			bg_color, face_color, face_x, face_y,
			left_eye_x, left_eye_y, right_eye_x, right_eye_y, bezier
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		UUID2Bytes(userId), username, hash, salt, 
		avatar.BgColor, avatar.FaceColor, 
		avatar.FaceX, avatar.FaceY, 
		avatar.LeX, avatar.LeY,
		avatar.ReX, avatar.ReY,
		avatar.Bezier,
	)
	if err != nil {
		if (errors.As(err, &mysqlErr)) {
			if (mysqlErr.Number == 1062) {
				return "", fmt.Errorf("This username is already in use")
			}
		}
		logger.Logger.Error("user insert error", "err", err.Error())
		return "", fmt.Errorf("Internal error, please try again later")
	}

	return userId.String(), nil
}

func Register(username string, password string, avatar models.Avatar) (string, error) {
	return RegisterEx(db, username, password, avatar)
}

func LoginEx(exec Execer, name string, password string) (models.DBUser, error) {
	var (
		ID []byte
		pwdHash []byte
		pwdSalt []byte
	)

	err := exec.QueryRow("SELECT id, pwd_hash, pwd_salt FROM users WHERE name = ?", name).Scan(&ID, &pwdHash, &pwdSalt)
	if (err != nil) {
		if (err == sql.ErrNoRows) {
			return models.DBUser{}, fmt.Errorf("Wrong username or password")
		}
		return models.DBUser{}, err;
	}

	if (!auth.CheckPassword(password, pwdHash, pwdSalt)) {
		return models.DBUser{}, errors.New("Wrong username or password");
	}

	uid, err := UUIDBytes2String(ID)
	if (err != nil) {
		logger.Logger.Error("UUIDBytes2String error", "err", err.Error())
		return models.DBUser{}, fmt.Errorf("there's a problem with your user, please try again later")
	}

	return models.DBUser{ // ;-;
		Account: models.Account{
			User: models.User{
				UID: uid,
			},
		},
	}, nil;
}

func Login(name string, password string) (models.DBUser, error) {
	return LoginEx(db, name, password)
}
