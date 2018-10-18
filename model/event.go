package model

import "time"

type Event struct {
	ID       string
	Channel  string
	Payload  []byte
	FireTime time.Time
}

func (e *Event) Sub(t time.Time) time.Duration {
	return e.FireTime.Sub(t)
}
