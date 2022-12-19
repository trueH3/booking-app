package cli

import (
	"booking-app/helper"
	"fmt"
	"github.com/google/uuid"
)

func RunCliApp() {
	var remainingTickets uint32 = helper.ConferenceTickets

	helper.PrintGreetUsers(remainingTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()

		if helper.IsConferenceSoldOut(remainingTickets, userTickets) {
			fmt.Printf("Our conference is sold out, %v tickets remains\n", remainingTickets)
			break
		}

		if !helper.IsValidNameAndSurname(firstName, lastName) {
			println("Please provide valid name and surname")
			continue
		}

		if !helper.IsValidEmail(email) {
			println("Please provide valid email ")
			continue
		}

		var bookedUser = produceUser(firstName, lastName, userTickets, email)
		helper.SaveBooking(&bookedUser)
		remainingTickets = remainingTickets - bookedUser.UserTickets
		helper.PrintConfirmation(bookedUser, remainingTickets, helper.GetAllBookings())

		go helper.SendTicket(bookedUser)
	}
}

func produceUser(firstName string, lastName string, userTickets uint32, email string) helper.User {

	//this is how we add element to array
	//bookings[0] = firstName + " " + lastName
	//this is howe we add eleement to slice ( comment from me but we create each time new slice in this way)
	//bookings = append(bookings, firstName+" "+lastName)

	return helper.User{
		Id:          uuid.NewString(),
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		UserTickets: userTickets,
	}
}
