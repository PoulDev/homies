package db

import (
	"fmt"
	"strconv"
	"crypto/rand"
	"math/big"
	"database/sql"
	"github.com/go-sql-driver/mysql"

	"github.com/PoulDev/homies/internal/homies/logger"
	"github.com/PoulDev/homies/internal/homies/models"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func GenerateCode() (string, error) {
	code := make([]byte, 6)
	for i := range code {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		code[i] = charset[num.Int64()]
	}
	return string(code), nil
}

// Returns the house ID & invite code
func NewHouse(name string) (string, string, error) {
    tx, err := db.Begin()
    if err != nil {return "", "", err}
    defer func (){
		if p := recover(); p != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var houseRes sql.Result
	var invite string
	for range 5 {
		invite, err = GenerateCode()
		if (err != nil) {
			logger.Logger.Error("house insert error", "err", err.Error())
			return "", "", fmt.Errorf("Internal error, please try again later")
		}

		houseRes, err = tx.Exec(`
			INSERT INTO houses (name, invite)
			VALUES (?, ?)`, name, invite,
		)

		if err == nil {
			break;
		}


		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			logger.Logger.Error("invite duplicate!?", "err", err.Error())
			continue
		}

		logger.Logger.Error("house insert error", "err", err.Error())
		break;
	}

	houseId, err := houseRes.LastInsertId()
	if err != nil {
		logger.Logger.Error("house ID retrival error", "err", err.Error())
		return "", "", fmt.Errorf("Internal error, please try again later")
	}

	err = NewListEx(tx, houseId, "shopping");
	if err != nil {
		logger.Logger.Error("shopping list insert error", "err", err.Error())
		return "", "", fmt.Errorf("Internal error, please try again later")
	}

	err = NewListEx(tx, houseId, "todo");
	if err != nil {
		logger.Logger.Error("todo list insert error", "err", err.Error())
		return "", "", fmt.Errorf("Internal error, please try again later")
	}

	return strconv.FormatInt(houseId, 10), invite, nil;
}

func GetHouseEx(exec Execer, house string, skipUser []byte) (models.House, error) {
	// TODO: Retrive house Members
	var resHouse models.House
	var houseid int64

	if (house == "null") {
		return models.House{}, fmt.Errorf("You don't have an house")
	}

	logger.Logger.Info("get house", "house", house)
	err := exec.QueryRow("SELECT id, name FROM houses WHERE id = ?", house).Scan(&houseid, &resHouse.Name)
	if (err != nil) {
		logger.Logger.Error("house get error", "err", err.Error())
		return models.House{}, fmt.Errorf("Internal error, please try again later")
	}

	resHouse.ID = strconv.FormatInt(houseid, 10)

	// Retrive house Members
	// Why this must be a mess every fucking time?
	
	rows, err := exec.Query(`SELECT id, name, bg_color, face_color, face_x, face_y, left_eye_x, left_eye_y, right_eye_x, right_eye_y, bezier FROM users WHERE house = ? AND id != ?`, houseid, skipUser);
	//rows, err := exec.Query(`SELECT id, name, bg_color, face_color, face_x, face_y, left_eye_x, left_eye_y, right_eye_x, right_eye_y, bezier FROM users WHERE house = ?`, houseid);
	defer rows.Close()

	if err != nil {
		logger.Logger.Error("list DB select error", "err", err.Error(), "houseId", houseid)
		return models.House{}, fmt.Errorf("Internal error, please try again later")
	}

	var users []models.User = make([]models.User, 0);
	for rows.Next() {
		var user models.User;
		var uid []byte;

		if err := rows.Scan(&uid, &user.Username, &user.Avatar.BgColor, &user.Avatar.FaceColor, &user.Avatar.FaceX, &user.Avatar.FaceY, &user.Avatar.LeX, &user.Avatar.LeY, &user.Avatar.ReX, &user.Avatar.ReY, &user.Avatar.Bezier); err != nil {
			logger.Logger.Error("list row scan error", "err", err.Error(), "houseId", houseid)
			return models.House{}, fmt.Errorf("There's a problem with your house, please try again later")
		}
		
		user.UID, err = UUIDBytes2String(uid)
		if (err != nil) {
			logger.Logger.Error("UUIDBytes2String error", "err", err.Error())
			return models.House{}, fmt.Errorf("There's a problem with your user, please try again later")
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		logger.Logger.Error("list rows error", "err", err.Error(), "houseId", houseid)
		return models.House{}, fmt.Errorf("There's a problem with your house, please try again later")
	}

	resHouse.Members = users

	return resHouse, nil;
}

func GetHouse(house string, skipUser []byte) (models.House, error) {
	return GetHouseEx(db, house, skipUser)
}
