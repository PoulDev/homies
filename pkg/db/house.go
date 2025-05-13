package db

import (
	"fmt"
)

type House struct {
	Name string
	Owner string
	Members []string
}


func NewHouse(name string, owner string) (string, error) {
	userRes, err := db.Exec(`
		INSERT INTO houses (name, owner_id)
		VALUES (?, ?)`,
		name, owner,
	)
	if err != nil {
		return "", err
	}

	userId, err := userRes.LastInsertId()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", userId), nil;
}

// ** Mai utilizzata, (forse) non necessaria
func GetHouse(house string) (House, error) {
	var (
		name string
		owner_id int
	)

	err := db.QueryRow("SELECT name, owner_id FROM houses WHERE id = ?", house).Scan(&name, &owner_id)

	if (err != nil) {
		return House{}, err
	}

	return House{
		Name: name,
		Owner: fmt.Sprintf("%d", owner_id),
	}, nil;
}


