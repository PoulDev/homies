package models

type House struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Members []User `json:"members"`
}
