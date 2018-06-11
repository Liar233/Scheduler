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

func (e *Event) Sub(t time.Time) time.Duration {
	return e.FireTime.Sub(t)
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

func (e *EventPool) Get(i int) *Event {
	return e.events[i]
}

func (e *EventPool) Len() int {
	return len(e.events)
}

func (e *EventPool) Swap(i, j int) {
	e.events[i], e.events[j] = e.events[j], e.events[i]
}

func (e *EventPool) Less(i, j int) bool {
	return e.events[i].FireTime.Before(e.events[j].FireTime)
}

func NewEventPool() EventPool {
	return EventPool{
		events: make([]*Event, 0),
	}
}
