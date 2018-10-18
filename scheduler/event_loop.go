package scheduler

import (
	"errors"
	"time"
	"sort"
	"github.com/Liar233/Scheduler/model"
)

type EventLoop struct {
	events   EventPool
	channels *ChannelPool
	running  bool
	add      chan *model.Event
	stop     chan interface{}
}

func NewEventLoop(channels *ChannelPool) *EventLoop {
	return &EventLoop{
		running:  false,
		channels: channels,
		events:   NewEventPool(),
	}
}

func (el EventLoop) AddChannel(ch model.ChannelInterface) error {
	return el.channels.Add(ch)
}

func (el *EventLoop) Push(event *model.Event) {
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
	el.add = make(chan *model.Event)
	el.stop = make(chan interface{})

	go el.run()
}

func (el *EventLoop) Snapshot() []*model.Event {
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

func (el *EventLoop) dispatch(event *model.Event) {
	go el.channels.DispatchEvent(event)
}
