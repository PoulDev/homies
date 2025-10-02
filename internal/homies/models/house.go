package models

type House struct {
	Invite  string `json:"invite"`
	Name    string `json:"name"`
	Owner   string `json:"owner"`
	Members []User `json:"members"`
}
