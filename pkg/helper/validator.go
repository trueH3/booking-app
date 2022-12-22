package helper

import "strings"

// capitalizing function name changes it access from package to public
func IsConferenceSoldOut(remainingTickets uint32, userTickets uint32) bool {
	return remainingTickets < userTickets
}

func IsValidNameAndSurname(firstName string, lastName string) bool {
	return len(firstName) >= 2 && len(lastName) >= 2
}

func IsValidEmail(email string) (isValid bool) {
	return strings.Contains(email, "@")
}

func IsRequestValid(user User, remainingTickets uint32) bool {
	return !IsConferenceSoldOut(remainingTickets, user.UserTickets) && IsValidNameAndSurname(user.FirstName, user.LastName) && IsValidEmail(user.Email)

}
