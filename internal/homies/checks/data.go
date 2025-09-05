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
			MinLength: 6,
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
	CheckersData["home_name"] = models.Checker{
		Check: models.Check{
			FriendlyName: "home name",
			MinLength: 4,
			MaxLength: 25,
		},
		Checker: BasicStringCheck("home_name"),
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
		if (len(value) <= CheckersData[key].MinLength) {
			return fmt.Errorf("Your %s must be at least %d characters long!", CheckersData[key].FriendlyName, CheckersData[key].MinLength);
		} else if (len(value) >= CheckersData[key].MaxLength) {
			return fmt.Errorf("Your %s is too long! max %d characters.", CheckersData[key].FriendlyName, CheckersData[key].MaxLength);
		}

		return nil;
	}
}

