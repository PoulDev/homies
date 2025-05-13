package db

import (
	"fmt"
)

type House struct {
	Name string
	Members []string
}

func NewHouse(name string) (string, error) {
	houseRes, err := db.Exec(`
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

// ** Mai utilizzata, potrebbe essere non necessaria
func GetHouse(house string) (House, error) {
	var (
		name string
	)

	err := db.QueryRow("SELECT name FROM houses WHERE id = ?", house).Scan(&name)

	if (err != nil) {
		return House{}, err
	}

	return House{
		Name: name,
	}, nil;
}


