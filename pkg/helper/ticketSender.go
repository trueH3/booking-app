package helper

import (
	"fmt"
	"time"
)

func SendTicket(user User) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v.", user.UserTickets, user.FirstName, user.LastName)
	fmt.Println("#########")
	fmt.Printf("Sending ticket: \n %v \n to email address %v\n", ticket, user.Email)
	fmt.Println("#########")
}
