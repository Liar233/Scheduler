package main

import (
	"errors"
	"time"
)

type EventLoop struct {
	events   []*Event
	channels map[string]ChannelInterface
	running  bool
	add      chan *Event
	stop     chan interface{}
}

func NewEventLoop() *EventLoop {
	return &EventLoop{
		running: false,
	}
}

func (e *EventLoop) AddChannel(ch ChannelInterface) error {
	if _, ok := e.channels[ch.Name()]; ok {
		errors.New("channel already exists")
	}

	e.channels[ch.Name()] = ch

	return nil
}

func (e *EventLoop) Push(event *Event) {
	if e.running == false {
		e.events = append(e.events, event)
	}

	e.add <- event
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
			close(e.add)
			e.running = false
			return
		}
	}
}

func (e *EventLoop) dispatch(event *Event) {
	_, err := e.channels[event.Channel]

	if !err  {
		// log error message channel not found

		return
	}

	timer := time.NewTimer(2 * time.Second)

	go e.send(event, timer.C)
}

func (e *EventLoop) send(event *Event, timer <-chan time.Time)  {
	<-timer
	err := e.channels[event.Channel].Process(event)

	if err != nil {
		// log event fail
	}

	e.events = e.events[:len(e.events)-1]
}
