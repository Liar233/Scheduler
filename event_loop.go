package main

import "errors"

type EventLoop struct {
	events   []*Event
	channels map[string]*ChannelInterface
	running  bool
}

func NewEventLoop() *EventLoop {
	return &EventLoop{
		running: false,
	}
}

func (e *EventLoop) Push(event *Event) {
	e.events = append(e.events, event)
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

	go e.process()
}

func (e *EventLoop) process() {

}