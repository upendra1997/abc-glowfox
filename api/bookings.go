package api

import (
	"abc/db"
	"abc/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

func BookClass(c *gin.Context) {
	var bookingRequest BookingRequest
	bodyReader := c.Request.Body
	body, err := io.ReadAll(bodyReader)
	if err != nil {
		util.HandleError(c, "Unable to read the request body", err)
		return
	}
	err = json.Unmarshal(body, &bookingRequest)
	if err != nil {
		util.HandleError(c, "Unable to parse the request body", err)
		return
	}
	userId := c.Param("user_id")
	if strings.TrimSpace(userId) == "" {
		util.HandleError(c, "No user id specified", nil)
		return
	}
	slog.Info("Parsed body is", "booking", bookingRequest, "user_id", userId)
	err = db.DB.AddBooking(db.Booking{
		User:  userId,
		Class: bookingRequest.Name,
		Date:  bookingRequest.Date,
	})
	if err != nil {
		util.HandleError(c, fmt.Sprintf("Unable to add the booking: %s", err), err)
		return
	}
	c.Status(http.StatusOK)
}

func Bookings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db.DB.GetBookings())
}

type BookingRequest struct {
	Name string    `json:"name"`
	Date util.Date `json:"date"`
}
