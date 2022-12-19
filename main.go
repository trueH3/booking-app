package main

import (
	"booking-app/cli"
	"booking-app/helper"
	"booking-app/rest"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Booking App. Please choose app mode cli/rest. If other value will be chosen then rest will be launched")

	helper.InitDb()
	var mode string

	fmt.Scanln(&mode)

	if mode == "cli" {
		cli.RunCliApp()
	}
	rest.RunRestApp()
}
