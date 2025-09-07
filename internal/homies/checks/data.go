package checks;

import (
	"fmt"

	"github.com/PoulDev/homies/internal/homies/models"
)

var CheckersData map[string]models.Checker;

func init() {
	fmt.Println("Initializing checks...")
	CheckersData = make(map[string]models.Checker);

	// Password
	CheckersData["password"] = models.Checker{
		Check: models.Check{
			FriendlyName: "password",
			MinLength: 8,
			MaxLength: 100,
		},
		Checker: BasicStringCheck("password"),
	}

	// Username
	CheckersData["username"] = models.Checker{
		Check: models.Check{
			FriendlyName: "username",
			MinLength: 3,
			MaxLength: 12,
		},
		Checker: BasicStringCheck("username"),
	}

	// Home name
	CheckersData["house_name"] = models.Checker{
		Check: models.Check{
			FriendlyName: "home name",
			MinLength: 4,
			MaxLength: 25,
		},
		Checker: BasicStringCheck("house_name"),
	}

	// List Item text
	CheckersData["list_item_text"] = models.Checker{
		Check: models.Check{
			FriendlyName: "item",
			MinLength: 3,
			MaxLength: 512,
		},
		Checker: BasicStringCheck("list_item_text"),
	}
}

func Check(key string, value string) error {
	return CheckersData[key].Checker(value)
}

func BasicStringCheck(key string) func(string) error {
	return func(value string) error {
		checker, ok := CheckersData[key]
		if (!ok) {
			return fmt.Errorf("Internal error: Checker %s not found!", key)
		}

		if (len(value) < checker.MinLength) {
			return fmt.Errorf("Your %s must be at least %d characters long!", checker.FriendlyName, checker.MinLength);
		} else if (len(value) >= checker.MaxLength) {
			return fmt.Errorf("Your %s is too long! max %d characters.", checker.FriendlyName, checker.MaxLength);
		}

		return nil;
	}
}

