package execers

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/zibbadies/homies/internal/homies/models"
	"github.com/zibbadies/homies/pkg/homies/auth"
)


func RegisterEx(exec Execer, username string, password string, avatar models.Avatar) (string, error) {
	hash, salt, err := auth.HashPassword(password)
	if err != nil {
		return "", err
	}

	var userId string
	err = exec.QueryRow(`
		INSERT INTO users (
			name, pwd_hash, pwd_salt,
			bg_color, face_color, face_x, face_y,
			left_eye_x, left_eye_y, right_eye_x, right_eye_y, bezier
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id`,
		username, hash, salt, 
		avatar.BgColor, avatar.FaceColor, 
		avatar.FaceX, avatar.FaceY, 
		avatar.LeX, avatar.LeY,
		avatar.ReX, avatar.ReY,
		avatar.Bezier,
	).Scan(&userId)

	if err != nil {
		return "", err
	}

	return userId, nil
}

func LoginEx(exec Execer, name string, password string) (models.DBUser, error) {
	var (
		ID string
		pwdHash []byte
		pwdSalt []byte
	)

	err := exec.QueryRow("SELECT id, pwd_hash, pwd_salt FROM users WHERE name = $1", name).Scan(&ID, &pwdHash, &pwdSalt)
	if (err != nil) {
		if (err == sql.ErrNoRows) {
			return models.DBUser{}, fmt.Errorf("Wrong username or password")
		}
		return models.DBUser{}, err;
	}

	if (!auth.CheckPassword(password, pwdHash, pwdSalt)) {
		return models.DBUser{}, errors.New("Wrong username or password");
	}

	return models.DBUser{ // ;-;
		Account: models.Account{
			User: models.User{
				UID: ID,
			},
		},
	}, nil;
}

