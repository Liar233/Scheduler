package action

import (
	"net/http"
	"github.com/Liar233/Scheduler/scheduler"
)

type EventLoopAction struct {
	el *scheduler.EventLoop
}

func (ela *EventLoopAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	events := ela.el.Snapshot()

	WriteSuccessResponse(w, events)
}

func NewEventLoopAction(el * scheduler.EventLoop) *EventLoopAction {
	return &EventLoopAction{
		el: el,
	}
}
