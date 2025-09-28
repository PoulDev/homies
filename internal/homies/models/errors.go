package models

const (
	UserNotFound     = "user_not_found"
	HouseNotFound    = "house_not_found"
	InviteNotFound   = "invite_not_found"
	UserNotInHouse   = "user_not_in_house"
	WrongCredentials = "wrong_credentials"
	InternalError    = "internal_error"
	UsernameTaken    = "username_taken"
	ListNotFound     = "list_not_found"
	ItemNotFound     = "item_not_found"
	JsonFormatError  = "json_format_error"
	BasicCheckError  = "basic_check_error"
	UserInHouse      = "user_in_house"
	NotAuthorized    = "cant_access_this"
	NotAuthenticated = "not_authenticated"
)

type DBError struct {
	Message   string `json:"message"`
	ErrorCode string `json:"code"`
}

func (e DBError) Error() string {
	return e.Message
}

type CheckError struct {
	Message   string `json:"message"`
	ErrorCode string `json:"code"`
}

func (e CheckError) Error() string {
	return e.Message
}
