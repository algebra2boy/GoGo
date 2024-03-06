package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, peopleInFrontQueue int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidNumber := peopleInFrontQueue >= 0

	return isValidName, isValidEmail, isValidNumber
}
