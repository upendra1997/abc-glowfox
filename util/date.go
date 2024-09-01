package util

import (
	"errors"
	"fmt"
	"time"
)

type Date struct{ time.Time }

func (d *Date) UnmarshalJSON(data []byte) error {
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("not a json string")
	}

	data = data[1 : len(data)-1]
	time, err := time.Parse(time.DateOnly, string(data))
	if err != nil {
		return err
	}
	*d = Date{time}
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", d.Format(time.DateOnly))), nil
}
