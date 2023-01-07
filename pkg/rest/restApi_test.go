package rest

import (
	"booking-app/pkg/helper"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockUserRepositoryStruct struct{}

func (repository mockUserRepositoryStruct) SaveBooking(user *helper.User) {
	fmt.Println("SaveBooking from mocked repository")
}

func (repository mockUserRepositoryStruct) GetAllBookings() []*helper.User {
	return []*helper.User{
		{
			Id:          "0e1ef792-864f-407b-8f38-394fd84a544a",
			FirstName:   "testName",
			LastName:    "testLastName",
			Email:       "testEmail",
			UserTickets: 3,
		}, {
			Id:          "e5f7b0fb-e627-4b7e-937a-282c60dcabde",
			FirstName:   "testName1",
			LastName:    "testLastName2",
			Email:       "testEmail3",
			UserTickets: 2,
		}}
}

func TestGetAllBookings(t *testing.T) {
	router := produceRouterWithHandlers()
	requestUrl := "/booking-app"
	userRepository = mockUserRepositoryStruct{}
	requestPointer, _ := http.NewRequest(http.MethodGet, requestUrl, nil)
	responseRecorderPointer := httptest.NewRecorder()

	router.ServeHTTP(responseRecorderPointer, requestPointer)

	responseData, _ := io.ReadAll(responseRecorderPointer.Body)
	var mappedResponse []*helper.User
	json.Unmarshal(responseData, &mappedResponse)

	assert.Equal(t, http.StatusOK, responseRecorderPointer.Code)
	assert.ElementsMatch(t, mappedResponse, userRepository.GetAllBookings())
}

func TestGetInfo(t *testing.T) {
	t.Log()
	router := produceRouterWithHandlers()
	requestUrl := "/booking-app/info"
	userRepository = mockUserRepositoryStruct{}
	requestPointer, _ := http.NewRequest(http.MethodGet, requestUrl, nil)
	responseRecorderPointer := httptest.NewRecorder()

	router.ServeHTTP(responseRecorderPointer, requestPointer)

	responseData, _ := io.ReadAll(responseRecorderPointer.Body)
	var mappedResponse = string(responseData)
	fmt.Println(mappedResponse)

	expected := "Welcome to Go conference booking application.\nWe have total of 50 tickets and 50 are still available.\nGet your tickets here for attend.\nPeople that are coming to our conference: [testName testName1]\n"
	assert.Equal(t, http.StatusOK, responseRecorderPointer.Code)
	assert.Equal(t, mappedResponse, expected)
}

func TestAddBookingWithValidRequest(t *testing.T) {
	defer setRemainingTickets(helper.ConferenceTickets) // it will reset shared variable across all test to its initial value
	router := produceRouterWithHandlers()
	requestUrl := "/booking-app"
	userRepository = mockUserRepositoryStruct{}

	user := helper.User{
		FirstName:   "sz3",
		LastName:    "Lambo",
		Email:       "random@com",
		UserTickets: 3,
	}
	requestBody, _ := json.Marshal(user)
	requestPointer, _ := http.NewRequest(http.MethodPost, requestUrl, bytes.NewBuffer(requestBody))
	responseRecorderPointer := httptest.NewRecorder()
	router.ServeHTTP(responseRecorderPointer, requestPointer)

	responseData, _ := io.ReadAll(responseRecorderPointer.Body)
	var mappedResponse helper.User
	json.Unmarshal(responseData, &mappedResponse)

	assert.Equal(t, http.StatusCreated, responseRecorderPointer.Code)
	assert.Equal(t, user.FirstName, mappedResponse.FirstName)
	assert.Equal(t, user.LastName, mappedResponse.LastName)
	assert.Equal(t, user.Email, mappedResponse.Email)
	assert.Equal(t, user.UserTickets, mappedResponse.UserTickets)
	assert.NotNil(t, mappedResponse.Id)
	assert.Equal(t, getRemainingTickets(), helper.ConferenceTickets-user.UserTickets)
}

func TestAddBookingWithTooShortName(t *testing.T) {
	router := produceRouterWithHandlers()
	requestUrl := "/booking-app"
	userRepository = mockUserRepositoryStruct{}

	user := helper.User{
		FirstName:   "s",
		LastName:    "Lambo",
		Email:       "random@com",
		UserTickets: 3,
	}
	requestBody, _ := json.Marshal(user)
	requestPointer, _ := http.NewRequest(http.MethodPost, requestUrl, bytes.NewBuffer(requestBody))
	responseRecorderPointer := httptest.NewRecorder()
	router.ServeHTTP(responseRecorderPointer, requestPointer)

	responseData, _ := io.ReadAll(responseRecorderPointer.Body)
	mappedResponse := string(responseData)

	assert.Equal(t, http.StatusBadRequest, responseRecorderPointer.Code)
	assert.Contains(t, mappedResponse, "please provide valid first name, last name, email, and be sure that we have enough tickets left")
}

func TestAddBookingRemainingTicketAfterConcurrentUpdate(t *testing.T) {
	defer setRemainingTickets(helper.ConferenceTickets) // it will reset shared variable across all test to its initial value
	router := produceRouterWithHandlers()
	requestUrl := "/booking-app"
	userRepository = mockUserRepositoryStruct{}

	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()

			user := helper.User{
				FirstName:   "sz3",
				LastName:    "Lambo",
				Email:       "random@com",
				UserTickets: 2,
			}
			requestBody, _ := json.Marshal(user)
			requestPointer, _ := http.NewRequest(http.MethodPost, requestUrl, bytes.NewBuffer(requestBody))
			responseRecorderPointer := httptest.NewRecorder()
			router.ServeHTTP(responseRecorderPointer, requestPointer)

			responseData, _ := io.ReadAll(responseRecorderPointer.Body)
			var mappedResponse helper.User
			json.Unmarshal(responseData, &mappedResponse)

			assert.Equal(t, http.StatusCreated, responseRecorderPointer.Code)
		}()
	}

	wg.Wait()

	assert.Equal(t, getRemainingTickets(), uint32(10))
}
