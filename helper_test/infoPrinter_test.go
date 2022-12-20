package helper_test

import (
	"booking-app/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGreetIsersAsStringShouldReturnCorrectString(t *testing.T) {
	result := helper.GetGreetUsersAsString(23)
	expectedVal := "Welcome to Go conference booking application.\nWe have total of 50 tickets and 23 are still available.\nGet your tickets here for attend.\n"

	assert.Equal(t, expectedVal, result)
}

func TestGetNamesOfAllAttendantsAsStringShouldReturnCorrectString(t *testing.T) {
	var users = []*helper.User{
		{FirstName: "szala", LastName: "lululu"},
		{FirstName: "ababa", LastName: "dibidi"},
	}
	result := helper.GetNamesOfAllAttendantsAsString(users)
	expected := "People that are coming to our conference: [" + users[0].FirstName + " " + users[1].FirstName + "]\n"

	assert.Equal(t, expected, result)
}
