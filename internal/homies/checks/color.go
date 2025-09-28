package checks

import (
	"strconv"
	"strings"

	"github.com/zibbadies/homies/internal/homies/models"
)

func BezierCheck(value string) error {
	for _, number := range strings.Split(value, " ") {
		if _, err := strconv.Atoi(number); err != nil {
			return &models.CheckError{
				Message: "Your bezier is not valid! It must be a list of numbers separated by spaces.",
				ErrorCode: models.BasicCheckError,
			}
		}
	}
	return nil;
}
