package models

type User struct {
	UID      string `json:"uid"`
	Username string `json:"username"`
	Avatar   Avatar `json:"avatar"`
}

type DBUser struct {
	User
	HouseId string `json:"house_id"`
}
