package scheduler

import (
	"errors"
	"time"
	"sort"
	"github.com/Liar233/Scheduler/model"
)

type EventLoop struct {
	eventPool EventPool
	channels  *ChannelPool
	running   bool
	add       chan *model.Event
	stop      chan interface{}
}

func NewEventLoop(channels *ChannelPool) *EventLoop {
	return &EventLoop{
		running:   false,
		channels:  channels,
		eventPool: NewEventPool(),
	}
}

func (el EventLoop) AddChannel(ch model.ChannelInterface) error {
	return el.channels.Add(ch)
}

func (el *EventLoop) Push(event *model.Event) {
	if el.running == false {
		el.eventPool.Push(event)
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
	return el.eventPool.Snapshot()
}

func (el *EventLoop) run() {
	var now time.Time

	for {
		sort.Sort(&el.eventPool)

		var timer *time.Timer

		if el.eventPool.Len() == 0 {
			timer = time.NewTimer(time.Duration(1000000) * time.Hour)
		} else {
			timer = time.NewTimer(el.eventPool.Get(0).Sub(time.Now()))
		}

		for {
			select {
			case now = <-timer.C:
				for _, event := range el.eventPool.Snapshot() {
					if now.After(event.FireTime) {
						el.dispatch(event)
					}
				}
			case newEvent := <-el.add:
				timer.Stop()
				el.eventPool.Push(newEvent)
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
