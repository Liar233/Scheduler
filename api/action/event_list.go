package action

import (
	"github.com/Liar233/Scheduler/storage"
	"net/http"
	"time"
)

type EventListAction struct {
	es *storage.EventStorage
}

func (a *EventListAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var params = make(map[string]map[string]interface{})

	now := time.Now()
	params["firetime"] = make(map[string]interface{})
	params["firetime"][">="] = now

	events, err := a.es.Query(params)

	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	WriteSuccessResponse(w, events)
}

func NewEventListAction(storage *storage.EventStorage) *EventListAction {
	return &EventListAction{
		es: storage,
	}
}
