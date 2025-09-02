package checks

import (
	"fmt"
)

func CheckHouseName(name string) error {
	if len(name) < 4 {
		return fmt.Errorf("House name must be at least 4 characters.")
	}
	if len(name) > 12 {
		return fmt.Errorf("House name cannot exceed 12 characters.")
	}
	return nil
}
