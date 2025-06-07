package db

import (
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
		return "", err
	}

	houseId, err := houseRes.LastInsertId()
	if err != nil {
		return "", err
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
		return House{}, err
	}

	return House{
		Name: name,
	}, nil;
}

func GetHouse(house string) (House, error) {
	return GetHouseEx(db, house)
}
