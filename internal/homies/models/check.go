package models

type Check struct {
	FriendlyName string `json:"friendly_name"`
	MinLength    int    `json:"min_length"`
	MaxLength    int    `json:"max_length"`
}

type Checker struct {
	Check
	Checker func(string) error
}
