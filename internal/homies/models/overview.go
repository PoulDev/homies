package models

type Overview struct {
	User  User   `json:"user"`
	Items []Item `json:"items"`
}
