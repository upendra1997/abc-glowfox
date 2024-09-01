package inmemory

import (
	schema "abc/db"
	"errors"
	"sync"
)

type DB struct {
	Classes  []schema.ClassInventory
	Bookings []schema.Booking
	lock     sync.Mutex
}

func (D DB) AddBooking(booking schema.Booking) error {
	db.lock.Lock()
	defer db.lock.Unlock()
	classIdx := -1
	for i, class := range db.Classes {
		if class.Name == booking.Class && class.Date.Time.Equal(booking.Date.Time) {
			classIdx = i
			break
		}
	}
	if classIdx == -1 {
		return errors.New("no such class found")
	}
	db.Classes[classIdx].Capacity -= 1
	db.Bookings = append(db.Bookings, booking)
	return nil
}

func (D DB) GetBookings() []schema.Booking {
	return db.Bookings
}

func (D DB) AddClassInventory(classes []schema.ClassInventory) error {
	db.lock.Lock()
	defer db.lock.Unlock()
	for _, class := range classes {
		classIdx := -1
		for i, dbClass := range db.Classes {
			if class.Name == dbClass.Name && dbClass.Date.Time.Equal(class.Date.Time) {
				classIdx = i
				db.Classes[i].Capacity += class.Capacity
			}
		}
		if classIdx == -1 {
			db.Classes = append(db.Classes, class)
		}
	}
	return nil
}

func (D DB) GetClasses() []schema.ClassInventory {
	return db.Classes
}

var db *DB
var lock sync.Once

func Get() *DB {
	lock.Do(func() {
		db = &DB{
			Classes:  []schema.ClassInventory{},
			Bookings: []schema.Booking{},
			lock:     sync.Mutex{},
		}
	})
	return db
}
