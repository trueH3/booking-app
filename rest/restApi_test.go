package rest

import (
	//	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func initGin() *gin.Engine {
	return gin.Default()
}

func TestGetAllBookings(t *testing.T) {
	//	mockResponse := `{"message": "all bookings"}`
	router := initGin()

	router.GET("/booking-app", getAllBookings)
	requestPointer, _ := http.NewRequest(http.MethodGet, "/booking-app", nil)
	responseRecorderPointer := httptest.NewRecorder()

	router.ServeHTTP(responseRecorderPointer, requestPointer)

	//	responseData, _ := ioutil.ReadAll(requestPointer.Body)

	// it will fail because I didn't mocked getAllBookings method and it's trying to reach in memory db
	assert.Equal(t, http.StatusOK, responseRecorderPointer.Code)

}
