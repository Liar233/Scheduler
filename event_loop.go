package main

import (
	"errors"
	"time"
	"sort"
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

func (el EventLoop) AddChannel(ch ChannelInterface) error {
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
	el.add = make(chan *Event)
	el.stop = make(chan interface{})

	go el.run()
}

func (el *EventLoop) Snapshot() []*Event {
	return el.events.Snapshot()
}

func (el *EventLoop) run() {
	var now time.Time

	for {
		sort.Sort(&el.events)

		var timer *time.Timer

		if el.events.Len() == 0 {
			timer = time.NewTimer(time.Duration(1000000) * time.Hour)
		} else {
			timer = time.NewTimer(el.events.Get(0).Sub(time.Now()))
		}

		for {
			select {
			case now = <-timer.C:
				for _, event := range el.events.Snapshot() {
					if now.After(event.FireTime) {
						el.dispatch(event)
					}
				}
			case newEvent := <-el.add:
				timer.Stop()
				el.events.Push(newEvent)
			case <-el.stop:
				close(el.add)
				el.running = false
				return
			}
			break
		}
	}
}

func (el *EventLoop) dispatch(event *Event) {
	go el.channels.DispatchEvent(event)
}
