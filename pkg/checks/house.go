package checks

import (
	"fmt"
)

func CheckHouseName(name string) error {
	if (len(name) <= 3) {
		return fmt.Errorf("Your house name must be at least 4 characters long!");
	} else if (len(name) >= 32) {
		return fmt.Errorf("Your house name is too long! max 12 characters.");
	}

	return nil;
}
