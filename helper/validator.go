package helper

import "strings"

// capitalizing function name changes it access from package to public
func IsConferenceSoldOut(remainingTickets uint, userTickets uint) bool {
	return remainingTickets < userTickets
}

func IsValidName(firstName string, lastName string) bool {
	return len(firstName) >= 2 && len(lastName) >= 2
}

func IsValidEmail(email string) (isValid bool) {
	return strings.Contains(email, "@")
}
