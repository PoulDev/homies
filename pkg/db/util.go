package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func IdOrnull(ID string) string {
	if (ID == "000000000000000000000000") {
		return "null"
	}
	return ID
}

func PreattyError(err error) string {
	if mongo.IsDuplicateKeyError(err) {
		return "This email address is already taken!"
	}

	if mongo.IsTimeout(err) {
		return "It seems like the database request took too long. Please try again later."
	}

	if mongo.IsNetworkError(err) {
		return "We had trouble connecting to the database."
	}

	return "Oops! Something went wrong, please try again later."
}
