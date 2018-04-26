package main

import (
	"errors"
)

type EventLoop struct {
	events   []*Event
	channels map[string]*ChannelInterface
	running  bool
	add      chan *Event
	stop     chan interface{}
}

func NewEventLoop() *EventLoop {
	return &EventLoop{
		running: false,
	}
}

func (e *EventLoop) Push(event *Event) {
	if e.running == false {
		e.events = append(e.events, event)
	}

	e.add <- event
}

func (e *EventLoop) Add(ch ChannelInterface) error {
	if _, ok := e.channels[ch.Name()]; ok {
		errors.New("channel already exists")
	}

	e.channels[ch.Name()] = &ch

	return nil
}

func (e *EventLoop) Start() {
	if e.running == true {
		panic(errors.New("event loop already running"))
	}

	e.running = true

	go e.run()
}

func (e *EventLoop) Snapshot() []*Event {
	snapshot := make([]*Event, 0)

	for _, event := range e.events {
		eventCopy := &Event{}

		*eventCopy = *event

		snapshot = append(snapshot, eventCopy)
	}

	return snapshot
}

func (e *EventLoop) run() {

	for {
		select {
		case newEvent := <-e.add:
			e.events = append(e.events, newEvent)
		case <-e.stop:
			e.running = false
			return
		}
	}
}
