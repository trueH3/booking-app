package helper_test

import (
	"booking-app/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsConferenceSoldOutShouldReturnValidResult(t *testing.T) {
	testCases := []struct {
		testName         string
		remainingTickets uint32
		userTickets      uint32
		expected         bool
	}{
		{"conference is not sold out", 30, 12, false},
		{"conference is sold out", 12, 30, true},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			result := helper.IsConferenceSoldOut(testCase.remainingTickets, testCase.userTickets)
			assert.Equal(t, testCase.expected, result)
		})
	}
}

func TestIsValidNameAndSurnameReturnValidResult(t *testing.T) {
	testCases := []struct {
		testName  string
		firstName string
		lastName  string
		expected  bool
	}{
		{"too short name", "s", "lambert", false},
		{"too short first and last name", "u", "w", false},
		{"too short last name", "Joseph", "w", false},
		{"valid first and last name", "Jo", "ho", true},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			result := helper.IsValidNameAndSurname(testCase.firstName, testCase.lastName)
			assert.Equal(t, testCase.expected, result)
		})
	}
}

func TestIsValidEmailReturnValidResult(t *testing.T) {
	testCases := []struct {
		testName string
		email    string
		expected bool
	}{
		{"email is invalid", "szalala.yahoo.com", false},
		{"email is valid", "szalala@yahoo.com", true},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			result := helper.IsValidEmail(testCase.email)
			assert.Equal(t, testCase.expected, result)
		})
	}
}

func TestIsRequestValidReturnValidResult(t *testing.T) {
	testCases := []struct {
		testName         string
		user             helper.User
		remainingTickets uint32
		expected         bool
	}{
		{"first name too short", helper.User{FirstName: "s", LastName: "lala", Email: "email@yahoo.com", UserTickets: 2}, 32, false},
		{"last name too short", helper.User{FirstName: "szala", LastName: "l", Email: "email@yahoo.com", UserTickets: 2}, 32, false},
		{"email invalid", helper.User{FirstName: "szala", LastName: "lala", Email: "emailyahoo.com", UserTickets: 2}, 32, false},
		{"conference is sold out", helper.User{FirstName: "szala", LastName: "lala", Email: "email@yahoo.com", UserTickets: 20}, 2, false},
		{"valid request", helper.User{FirstName: "szala", LastName: "lala", Email: "email@yahoo.com", UserTickets: 2}, 20, true},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			result := helper.IsRequestValid(testCase.user, testCase.remainingTickets)
			assert.Equal(t, testCase.expected, result)
		})
	}
}
