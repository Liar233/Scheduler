package main

import (
	"errors"
)

type EventLoop struct {
	events   EventPool
	channels ChannelPool
	running  bool
	add      chan *Event
	stop     chan interface{}
}

func NewEventLoop(channels ChannelPool) *EventLoop {
	return &EventLoop{
		running:  false,
		channels: channels,
		events:   NewEventPool(),
	}
}

func (el *EventLoop) AddChannel(ch ChannelInterface) error {
	return el.channels.Add(ch)
}

func (el *EventLoop) Push(event *Event) {
	if el.running == false {
		el.events.Push(event)
	}

	el.add <- event
}

func (el *EventLoop) Start() {
	if el.running == true {
		panic(errors.New("event loop already running"))
	}

	el.running = true

	go el.run()
}

func (el *EventLoop) Snapshot() []*Event {
	return el.events.Snapshot()
}

func (el *EventLoop) run() {

	for {
		select {
		case newEvent := <-el.add:
			el.events.Push(newEvent)
		case <-el.stop:
			close(el.add)
			el.running = false
			return
		}
	}
}

func (el *EventLoop) dispatch(event *Event) {
	go el.channels.DispatchEvent(event)
}
