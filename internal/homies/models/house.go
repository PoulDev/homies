package models

type House struct {
	Invite  string `json:"invite"`
	Name    string `json:"name"`
	Members []User `json:"members"`
}
