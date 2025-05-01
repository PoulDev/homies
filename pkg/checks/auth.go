package checks

func CheckPassword(password string) string {
	if (len(password) <= 6) {
		return "Your password must be at least 6 characters long!";
	} else if (len(password) >= 100) {
		return "Your password is too long! max 100 characters.";
	}

	return ""
}

func CheckUsername(username string) string {
	if (len(username) <= 3) {
		return "Your username is too short!"
	} else if (len(username) >= 12) {
		return "Your username is too long!"
	}

	return ""
}
