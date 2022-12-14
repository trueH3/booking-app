package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
)

var WaitGroup = sync.WaitGroup{}

func main() {

	var remainingTickets uint = helper.ConferenceTickets
	// this is array, defining array requires size , if we don't know the size then we can use slice , slice is an array with dynamic size
	// var bookings = [conferenceTickets]string{}
	// this is a slice
	var bookings = []helper.User{}

	helper.GreetUsers(remainingTickets)

	for {
		firstName, lastName, email, userTickets := helper.GetUserInput()

		if helper.IsConferenceSoldOut(remainingTickets, userTickets) {
			fmt.Printf("Our conference is sold out, %v tickets remains\n", remainingTickets)
			break
		}

		if !helper.IsValidName(firstName, lastName) {
			println("Please provide valid name and surname")
			continue
		}

		if !helper.IsValidEmail(email) {
			println("Please provide valid email ")
			continue
		}

		var bookedUser = bookUser(firstName, lastName, userTickets, email, remainingTickets, bookings)
		// to run code in separate thread just use 'go'
		// but without thread synchro when main thread will be finished before SendTicket, then ticket won't be sent. Thats why we define waitGroup
		WaitGroup.Add(1) // means that untill counter becomes 0 then this thread - main thread ( go routine) caan't be finished and need to wait
		go sendTicket(bookedUser)
	}
	WaitGroup.Wait() // this waits when counter reaches 0 before releasing routine
}

func bookUser(firstName string, lastName string, userTickets uint, email string, remainingTickets uint, bookings []helper.User) helper.User {
	remainingTickets = remainingTickets - userTickets

	//this is how we add element to array
	//bookings[0] = firstName + " " + lastName
	//this is howe we add eleement to slice ( comment from me but we create each time new slice in this way)
	//bookings = append(bookings, firstName+" "+lastName)

	var bookedUser = helper.User{
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		UserTickets: userTickets,
	}
	bookings = append(bookings, bookedUser)

	helper.PrintConfirmation(bookedUser, remainingTickets, bookings)
	return bookedUser
}
