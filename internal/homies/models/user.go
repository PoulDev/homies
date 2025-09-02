package models

type User struct {
	UID      string `json:"uid"`
	Username string `json:"username"`
	House    House  `json:"house"`
	Avatar   Avatar `json:"avatar"`
}

type HouseMember struct {
	UID      string `json:"uid"`
	Username string `json:"username"`
	Avatar   Avatar `json:"avatar"`
}
