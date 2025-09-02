package db

import (
	"github.com/PoulDev/roommates-api/internal/homies/logger"
	"fmt"
)

type House struct {
	Name string
	Members []string
}

func NewHouseEx(exec Execer, name string) (string, error) {
	houseRes, err := exec.Exec(`
		INSERT INTO houses (name)
		VALUES (?)`, name,
	)
	if err != nil {
		logger.Logger.Error("house insert error", "err", err.Error())
		return "", fmt.Errorf("Internal error, please try again later")
	}

	houseId, err := houseRes.LastInsertId()
	if err != nil {
		logger.Logger.Error("house ID retrival error", "err", err.Error())
		return "", fmt.Errorf("Internal error, please try again later")
	}

	return fmt.Sprintf("%d", houseId), nil;
}

func NewHouse(name string) (string, error) {
	return NewHouseEx(db, name)
}


// ** Mai utilizzata, potrebbe essere non necessaria
func GetHouseEx(exec Execer, house string) (House, error) {
	var (
		name string
	)

	err := exec.QueryRow("SELECT name FROM houses WHERE id = ?", house).Scan(&name)
	if (err != nil) {
		logger.Logger.Error("house get error", "err", err.Error())
		return House{}, fmt.Errorf("Internal error, please try again later")
	}

	return House{
		Name: name,
	}, nil;
}

func GetHouse(house string) (House, error) {
	return GetHouseEx(db, house)
}
