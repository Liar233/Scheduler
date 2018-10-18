package scheduler

import (
	"github.com/Liar233/Scheduler/model"
)

type EventPool struct {
	events []*model.Event
}

func (e *EventPool) Push(event *model.Event) {
	e.events = append(e.events, event)
}

func (e *EventPool) Snapshot() []*model.Event {
	snapshot := make([]*model.Event, 0)

	for _, event := range e.events {
		eventCopy := &model.Event{}

		*eventCopy = *event

		snapshot = append(snapshot, eventCopy)
	}

	return snapshot
}

func (e *EventPool) Get(i int) *model.Event {
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
		events: make([]*model.Event, 0),
	}
}
