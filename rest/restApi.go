package rest

import (
	"booking-app/helper"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
)

var bookings = []helper.User{}

// this needs to be thread safe, do not use it directly
var remainingTickets = helper.ConferenceTickets

func getRemainingTickets() uint32 {
	return atomic.LoadUint32(&remainingTickets)
}

func setRemainingTickets(valueToSet uint32) {
	atomic.StoreUint32(&remainingTickets, valueToSet)
}

func RunRestApp() {
	var router = gin.Default()
	router.GET("/booking-app", getAllBookings)
	router.POST("/booking-app", addBooking)
	router.GET("/booking-app/info", getInfo)
	router.Run("localhost:8080")
}

func getAllBookings(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, bookings)
}

func getInfo(context *gin.Context) {
	infoText := helper.GetGreetUsersAsString(getRemainingTickets()) + helper.GetNamesOfAllAttendantsAsString(bookings)
	context.String(http.StatusOK, infoText)
}

func addBooking(context *gin.Context) {
	var newBooking helper.User

	err := context.BindJSON(&newBooking)
	if err != nil {
		return
	}

	if !helper.IsRequestValid(newBooking, getRemainingTickets()) {
		err := errors.New("Please provide valid first name, last name, email, and be sure that we have enough tickets left").Error()
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	bookings = append(bookings, newBooking)
	setRemainingTickets(getRemainingTickets() - newBooking.UserTickets)
	context.IndentedJSON(http.StatusCreated, newBooking)
	go helper.SendTicket(newBooking)
	helper.PrintConfirmation(newBooking, getRemainingTickets(), bookings)
}
