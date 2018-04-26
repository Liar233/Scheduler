package main

import "time"

type Event struct {
	ID       string
	Schedule string
	Payload  []byte
	FireTime time.Time
}
