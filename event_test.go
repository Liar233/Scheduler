package main

import (
	"testing"
	"reflect"
	"time"
)

func TestEventPool_Snapshot(t *testing.T) {
	pool := NewEventPool()

	events := getEvents()

	for _, event := range (events) {
		pool.Push(event)
	}

	if !reflect.DeepEqual(pool.Snapshot(), events) {
		t.Fail()
		t.Error("Events not equal")
	}
}

func TestEventPool_Less(t *testing.T) {
	pool := NewEventPool()

	events := getEvents()

	for _, event := range (events) {
		pool.Push(event)
	}

	if !pool.Less(0, 1) {
		t.Fail()
		t.Error("Events not less")
	}

	if pool.Less(1, 0) {
		t.Fail()
		t.Error("Events not less")
	}
}

func getEvents() []*Event {
	events := make([]*Event, 0)

	for i := 0; i < 3; i++ {
		eventID := "event_" + string(i)

		fireTime := time.Now()
		fireTime.Add(time.Duration(i*5) * time.Minute)

		event := &Event{
			FireTime: fireTime,
			ID:       eventID,
		}

		events = append(events, event)
	}

	return events
}
