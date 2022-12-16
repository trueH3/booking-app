package helper

import "strings"

// capitalizing function name changes it access from package to public
func isConferenceSoldOut(remainingTickets uint, userTickets uint) bool {
	return remainingTickets < userTickets
}

func isValidName(firstName string, lastName string) bool {
	return len(firstName) >= 2 && len(lastName) >= 2
}

func isValidEmail(email string) (isValid bool) {
	return strings.Contains(email, "@")
}
