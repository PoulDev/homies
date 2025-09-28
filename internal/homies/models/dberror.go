package models;

const (
	UserNotFound = "user_not_found"
	HouseNotFound = "house_not_found"
	WrongCredentials = "wrong_credentials"
	InternalError = "internal_error"
	UsernameTaken = "username_taken"
)

type DBError struct {
	Message string
	ErrorCode string
}

func (e DBError) Error() string {
	return e.Message
}
