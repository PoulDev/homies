package models

type Overview struct {
	User  User   `json:"user"`
	House House  `json:"house"`
	Items []Item `json:"items"`
}
