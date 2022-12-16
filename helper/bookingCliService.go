package helper

import (
	"fmt"
	"sync"
)

var waitGroup = sync.WaitGroup{}

func RunCliApp() {
	var remainingTickets uint = conferenceTickets
	// this is array, defining array requires size , if we don't know the size then we can use slice , slice is an array with dynamic size
	// var bookings = [conferenceTickets]string{}
	// this is a slice
	var bookings = []User{}

	greetUsers(remainingTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()

		if isConferenceSoldOut(remainingTickets, userTickets) {
			fmt.Printf("Our conference is sold out, %v tickets remains\n", remainingTickets)
			break
		}

		if !isValidName(firstName, lastName) {
			println("Please provide valid name and surname")
			continue
		}

		if !isValidEmail(email) {
			println("Please provide valid email ")
			continue
		}

		var bookedUser = produceUser(firstName, lastName, userTickets, email)
		bookings = append(bookings, bookedUser)
		remainingTickets = remainingTickets - bookedUser.UserTickets
		printConfirmation(bookedUser, remainingTickets, bookings)

		// to run code in separate thread just use 'go'
		// but without thread synchro when main thread will be finished before SendTicket, then ticket won't be sent. Thats why we define waitGroup
		waitGroup.Add(1) // means that untill counter becomes 0 then this thread - main thread ( go routine) caan't be finished and need to wait
		go sendTicket(bookedUser)
	}
	waitGroup.Wait() // this waits when counter reaches 0 before releasing routine
}

func produceUser(firstName string, lastName string, userTickets uint, email string) User {

	//this is how we add element to array
	//bookings[0] = firstName + " " + lastName
	//this is howe we add eleement to slice ( comment from me but we create each time new slice in this way)
	//bookings = append(bookings, firstName+" "+lastName)

	return User{
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		UserTickets: userTickets,
	}
}
