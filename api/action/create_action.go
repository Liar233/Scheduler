package action

import (
	"github.com/Liar233/Scheduler/storage"
	"net/http"
	"encoding/json"
	"github.com/Liar233/Scheduler/model"
)

type CreateEventAction struct {
	es *storage.EventStorage
}

func (a *CreateEventAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var event model.Event
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&event)

	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	event, err = a.es.Create(event)

	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	WriteSuccessResponse(w, event)
}

func NewCreateEventAction(storage *storage.EventStorage) *CreateEventAction {
	return &CreateEventAction{
		es: storage,
	}
}
