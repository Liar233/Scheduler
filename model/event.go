package model

import "time"

type Event struct {
	ID       string    `json:"id"`
	Channel  string    `json:"channel"`
	Payload  string    `json:"payload"`
	FireTime time.Time `json:"firetime"`
}

func (e *Event) Sub(t time.Time) time.Duration {
	return e.FireTime.Sub(t)
}
