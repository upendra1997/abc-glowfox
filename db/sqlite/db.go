package sqlite

import (
	schema "abc/db"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log/slog"
	"sync"
)

type DB struct {
	lock sync.Once
	db   *gorm.DB
}

func (db DB) AddBooking(booking schema.Booking) error {
	//TODO implement me
	panic("implement me")
}

func (db DB) GetBookings() []schema.Booking {
	//TODO implement me
	panic("implement me")
}

func (db DB) AddClassInventory(classes []schema.ClassInventory) error {
	//TODO implement me
	panic("implement me")
}

func (db DB) GetClasses() []schema.ClassInventory {
	//TODO implement me
	panic("implement me")
}

var db *DB

func Get() *DB {
	db.lock.Do(func() {
		db_, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
		if err != nil {
			slog.Error("Unable to connect to DB")
		}
		db.db = db_
	})
	return db
}
