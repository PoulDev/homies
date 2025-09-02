package checks

import (
	"fmt"
)

func CheckPassword(password string) error {
	if (len(password) <= 6) {
		return fmt.Errorf("Your password must be at least 6 characters long!");
	} else if (len(password) >= 100) {
		return fmt.Errorf("Your password is too long! max 100 characters.");
	}

	return nil;
}

func CheckUsername(username string) error {
	if (len(username) <= 3) {
		return fmt.Errorf("Your username must be at least 3 characters long!");
	} else if (len(username) >= 12) {
		return fmt.Errorf("Your username is too long! max 12 characters.");
	}

	return nil
}
