package main

import (
	"testing"
	"reflect"
)

func TestEventPool_Snapshot(t *testing.T) {
	pool := NewEventPool()

	events := getEvents()

	for _, event := range(events)  {
		pool.Push(event)
	}

	if !reflect.DeepEqual(pool.Snapshot(), events) {
		t.Fail()
		t.Error("Events not equal")
	}
}

func getEvents() []*Event {
	events := make([]*Event, 0)

	for i := 0; i < 3; i++ {
		eventID := "event_" + string(i)

		event := &Event{
			ID: eventID,
		}

		events = append(events, event)
	}

	return events
}
