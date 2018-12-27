package scheduler

import (
	"github.com/Liar233/Scheduler/storage"
	"errors"
	"time"
)

type EventProcessor struct {
	es            *storage.EventStorage
	el            *EventLoop
	running       bool
	freezeTimeout uint
	eventStep     uint
	stop          chan interface{}
}

func (ep *EventProcessor) Start() {
	if ep.running == true {
		panic(errors.New("event loop already running"))
	}

	ep.running = true

	go ep.run()
}

func (ep *EventProcessor) run() {
	for {
		var timer *time.Timer

		timer = time.NewTimer(time.Duration(ep.eventStep) * time.Minute)

		select {
		case <-timer.C:
			now := time.Now()

			params := make(map[string]map[string]interface{})
			params["firetime"][">="] = now
			params["firetime"]["<="] = now.Add(time.Minute * time.Duration(ep.freezeTimeout))

			events, err := ep.es.Query(params)

			if err != nil {
				println(err)
				continue
			}

			for _, event := range events  {
				ep.el.Push(event)
			}

			break
		case <-ep.stop:
			ep.running = false
			return
		}
	}
}
