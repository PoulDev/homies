package models

type List struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
}

type Item struct {
	Id        string `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
	Author    string `json:"author"`
}
