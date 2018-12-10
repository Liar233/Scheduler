package scheduler

import (
	"github.com/Liar233/Scheduler/model"
	"time"
	"fmt"
)

type EventPool struct {
	events []*model.Event
}

func (ep *EventPool) Push(event *model.Event) error {
	if time.Now().Before(event.FireTime) {
		ep.events = append(ep.events, event)
	} else {
		return fmt.Errorf("Event %s is to late...", event.ID)
	}

	return nil
}

func (ep *EventPool) Snapshot() []*model.Event {
	snapshot := make([]*model.Event, 0)

	for _, event := range ep.events {
		eventCopy := &model.Event{}

		*eventCopy = *event

		snapshot = append(snapshot, eventCopy)
	}

	return snapshot
}

func (ep *EventPool) Get(i int) *model.Event {
	return ep.events[i]
}

func (ep *EventPool) Len() int {
	return len(ep.events)
}

func (ep *EventPool) Swap(i, j int) {
	ep.events[i], ep.events[j] = ep.events[j], ep.events[i]
}

func (ep *EventPool) Less(i, j int) bool {
	return ep.events[i].FireTime.Before(ep.events[j].FireTime)
}

func (ep *EventPool) Remove(event *model.Event) {
	for i, e := range ep.events {
		if e == event {
			ep.events = append(ep.events[:i], ep.events[i+1:]...)
		}
	}
}

func NewEventPool() EventPool {
	return EventPool{
		events: make([]*model.Event, 0),
	}
}
