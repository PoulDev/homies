package checks

import (
	"fmt"
	"strconv"
	"strings"
)

func BezierCheck(value string) error {
	for _, number := range strings.Split(value, " ") {
		if _, err := strconv.Atoi(number); err != nil {
			return fmt.Errorf("Your bezier is not valid! It must be a list of numbers separated by spaces.");
		}
	}
	return nil;
}
