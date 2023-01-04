// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package generated

// Booking defines model for Booking.
type Booking struct {
	// Email email of the person. Must contains '@'
	Email string `json:"Email"`

	// FirstName name of the person. Must be longer than 2
	FirstName string `json:"FirstName"`

	// LastName last name of the person. Must be longer than 2
	LastName string `json:"LastName"`

	// UserTickets number of ordered tickets. Must be greater than 0
	UserTickets int32 `json:"UserTickets"`
}

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// CreateBookingJSONRequestBody defines body for CreateBooking for application/json ContentType.
type CreateBookingJSONRequestBody = Booking
