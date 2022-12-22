package main

import (
	"booking-app/pkg/cli"
	"booking-app/pkg/helper"
	"booking-app/pkg/rest"
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
