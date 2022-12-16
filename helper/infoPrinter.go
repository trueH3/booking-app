package helper

import "fmt"

const conferenceName = "Go conference"

// capital letter changes visibility to public for all packages
const conferenceTickets = 50

func greetUsers(remainingTickets uint) {
	fmt.Println("Welcome to", conferenceName, "booking application.")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here for attend")
}

func printConfirmation(user User, remainingTickets uint, bookings []User) {
	fmt.Printf("User %v %v, booked %v tickets. Confirmation will be sent on %v\n", user.FirstName, user.LastName, user.UserTickets, user.Email)
	fmt.Printf("%v tickets remiaining\n", remainingTickets)
	printNamesOfAllAttendants(bookings, user)

}

func printNamesOfAllAttendants(bookings []User, user User) {
	// how to get collection of only names using for each loop
	// range means get index and vqlue for given collection
	// identyfikator _ - normalnie dałbym tam wartość index czy coś takiego ale wtedy kompilator mówi ze mam uzyć tej wartości zatem jest po to
	// by dać mu znać ze moze olać tą wartość
	var firstNames = []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.FirstName)
	}
	fmt.Printf("People that are coming to our conference: %v\n", firstNames)
}
