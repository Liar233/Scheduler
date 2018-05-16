package main

import (
	"time"
)

type Event struct {
	ID       string
	Schedule string
	Channel  string
	Payload  []byte
	FireTime time.Time
}

type EventPool struct {
	events []*Event
}

func (e *EventPool) Push(event *Event) {
	e.events = append(e.events, event)
}

func (e *EventPool) Snapshot() []*Event {
	snapshot := make([]*Event, 0)

	for _, event := range e.events {
		eventCopy := &Event{}

		*eventCopy = *event

		snapshot = append(snapshot, eventCopy)
	}

	return snapshot
}

func (e *EventPool) Len() int {
	return len(e.events)
}

func NewEventPool() EventPool {
	return EventPool{
		events: make([]*Event, 0),
	}
}
