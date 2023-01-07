package rest

import (
	"booking-app/generated"
	"booking-app/pkg/helper"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"sync/atomic"
)

// this needs to be thread safe, do not use it directly, use atomic operation
var remainingTickets = helper.ConferenceTickets
var userRepository helper.IUserRepository = helper.UserRepositoryStruct{}

type bookingAppServer struct {
}

func (bookingAppServer) GetBookings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, userRepository.GetAllBookings())
}

func (bookingAppServer) GetInfo(c *gin.Context) {
	infoText := helper.GetGreetUsersAsString(getRemainingTickets()) + helper.GetNamesOfAllAttendantsAsString(userRepository.GetAllBookings())
	c.String(http.StatusOK, infoText)
}

func (bookingAppServer) CreateBooking(c *gin.Context) {
	var bookingRequest helper.User

	err := c.BindJSON(&bookingRequest)
	if err != nil {
		return
	}

	if !helper.IsRequestValid(bookingRequest, getRemainingTickets()) {
		err := errors.New("please provide valid first name, last name, email, and be sure that we have enough tickets left").Error()
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	bookingRequest.Id = uuid.NewString()
	userRepository.SaveBooking(&bookingRequest)
	setRemainingTickets(getRemainingTickets() - bookingRequest.UserTickets)
	c.IndentedJSON(http.StatusCreated, bookingRequest)
	go helper.SendTicket(bookingRequest)
	helper.PrintConfirmation(bookingRequest, getRemainingTickets(), userRepository.GetAllBookings())
}

func getRemainingTickets() uint32 {
	return atomic.LoadUint32(&remainingTickets)
}

func setRemainingTickets(valueToSet uint32) {
	atomic.StoreUint32(&remainingTickets, valueToSet)
}

func RunRestApp() {
	router := produceRouterWithHandlers()
	router.Run("localhost:8080")
}

func produceRouterWithHandlers() *gin.Engine {
	var router = gin.Default()
	bookingServer := bookingAppServer{}
	router = generated.RegisterHandlers(router, bookingServer)
	return router
}
