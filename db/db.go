package db

import "abc/util"

type Db interface {
	AddClassInventory(classes []ClassInventory) error
	GetClasses() []ClassInventory
	AddBooking(booking Booking) error
	GetBookings() []Booking
}

type Booking struct {
	User  string    `json:"user_name"`
	Class string    `json:"class_name"`
	Date  util.Date `json:"date"`
}

type ClassInventory struct {
	Name     string    `json:"name"`
	Date     util.Date `json:"date"`
	Capacity int       `json:"capacity"`
}

var DB Db
