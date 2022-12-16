package helper

import "fmt"

const conferenceName = "Go conference"

// capital letter changes visibility to public for all packages
const ConferenceTickets uint32 = 50
const invitationalInfo = "Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here for attend.\n"
const attendantsInfo = "People that are coming to our conference: %v\n"

func GetGreetUsersAsString(remainingTickets uint32) string {
	return fmt.Sprintf(invitationalInfo, conferenceName, ConferenceTickets, remainingTickets)
}

func PrintGreetUsers(remainingTickets uint32) {
	fmt.Println(GetGreetUsersAsString(remainingTickets))
}

func printNamesOfAllAttendants(bookings []User) {
	fmt.Println(GetNamesOfAllAttendantsAsString(bookings))
}

func GetNamesOfAllAttendantsAsString(bookings []User) string {
	// range means get index and value for given collection, _ means ignore returning index cause i do not use it anywhere
	var firstNames = []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.FirstName)
	}
	return fmt.Sprintf(attendantsInfo, firstNames)
}

func PrintConfirmation(user User, remainingTickets uint32, bookings []User) {
	fmt.Printf("User %v %v, booked %v tickets. Confirmation will be sent on %v\n", user.FirstName, user.LastName, user.UserTickets, user.Email)
	fmt.Printf("%v tickets remiaining\n", remainingTickets)
	printNamesOfAllAttendants(bookings)

}
