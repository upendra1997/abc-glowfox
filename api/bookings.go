package api

import (
	"abc/db"
	"abc/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

func BookClass(c *gin.Context) {
	var bookingRequest BookingRequest
	bodyReader, err := c.Request.GetBody()
	if err != nil {
		util.HandelError(c, "Request body is not found", err)
		return
	}
	body, err := io.ReadAll(bodyReader)
	if err != nil {
		util.HandelError(c, "Unable to read the request body", err)
		return
	}
	err = json.Unmarshal(body, &bookingRequest)
	if err != nil {
		util.HandelError(c, "Unable to parse the request body", err)
		return
	}
	userId := c.Param("user_id")
	if strings.TrimSpace(userId) == "" {
		util.HandelError(c, "No user id specified", nil)
		return
	}
	slog.Info("Parsed body is", "booking", bookingRequest, "user_id", userId)
	err = db.DB.AddBooking(db.Booking{
		User:  userId,
		Class: bookingRequest.Name,
		Date:  bookingRequest.Date,
	})
	if err != nil {
		util.HandelError(c, "Unable to add the booking", err)
		return
	}
	c.Status(http.StatusOK)
}

func Bookings(c *gin.Context) {
	userId := c.Param("user_id")
	if strings.TrimSpace(userId) == "" {
		util.HandelError(c, "No user id specified", nil)
	}
	slog.Info("Parsed body is", "user_id", userId)
	c.IndentedJSON(http.StatusOK, db.DB.GetBookings())
}

type BookingRequest struct {
	Name string    `json:"name"`
	Date util.Date `json:"date"`
}
