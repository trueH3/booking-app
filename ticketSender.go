package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

func sendTicket(user helper.User) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v.", user.UserTickets, user.FirstName, user.LastName)
	fmt.Println("#########")
	fmt.Printf("Sending ticket: \n %v \n to email address %v\n", ticket, user.Email)
	fmt.Println("#########")

	// this decrements counter by 1 for workgroup
	WaitGroup.Done()
}
