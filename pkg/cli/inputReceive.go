package cli

import "fmt"

func getUserInput() (string, string, string, uint32) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint32

	fmt.Println("Please provide number of tickets")
	fmt.Scanln(&userTickets)
	fmt.Println("Please provide your name")

	// without & will not work, & stands for pointer, pointer in golang is also called special variable, pointer is a variable that stores memory addres of given variable
	fmt.Scanln(&firstName)

	fmt.Println("Please provide your last name")
	fmt.Scanln(&lastName)

	fmt.Println("Please provide your email")
	fmt.Scanln(&email)

	return firstName, lastName, email, userTickets
}
