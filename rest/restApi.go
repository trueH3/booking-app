package rest

import (
	"booking-app/helper"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"sync/atomic"
)

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
	context.IndentedJSON(http.StatusOK, helper.GetAllBookings())
}

func getInfo(context *gin.Context) {
	infoText := helper.GetGreetUsersAsString(getRemainingTickets()) + helper.GetNamesOfAllAttendantsAsString(helper.GetAllBookings())
	context.String(http.StatusOK, infoText)
}

func addBooking(context *gin.Context) {
	var bookingRequest helper.User

	err := context.BindJSON(&bookingRequest)
	if err != nil {
		return
	}

	if !helper.IsRequestValid(bookingRequest, getRemainingTickets()) {
		err := errors.New("please provide valid first name, last name, email, and be sure that we have enough tickets left").Error()
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	bookingRequest.Id = uuid.NewString()
	helper.SaveBooking(&bookingRequest)
	setRemainingTickets(getRemainingTickets() - bookingRequest.UserTickets)
	context.IndentedJSON(http.StatusCreated, bookingRequest)
	go helper.SendTicket(bookingRequest)
	helper.PrintConfirmation(bookingRequest, getRemainingTickets(), helper.GetAllBookings())
}
