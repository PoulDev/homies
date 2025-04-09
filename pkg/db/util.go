package db

import (
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func IdOrnull(ID string) string {
	if (ID == "000000000000000000000000") {
		return "null"
	}
	return ID
}

func PreattyError(err error) error {
	log.Println("Preattifing the following error: ", err.Error());

	if mongo.IsTimeout(err) {
		return errors.New("It seems like the database request took too long. Please try again later.")
	}

	if mongo.IsNetworkError(err) {
		return errors.New("We had trouble connecting to the database.")
	}

	return err
}
