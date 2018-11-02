package action

import (
	"github.com/Liar233/Scheduler/storage"
	"net/http"
)

type DeleteEventAction struct {
	es *storage.EventStorage
}

func (a *DeleteEventAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	event, err := a.es.Get(id)

	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	err = a.es.Delete(event)

	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	var stub interface{}
	WriteSuccessResponse(w, stub)
}

func NewDeleteEventAction(storage *storage.EventStorage) *DeleteEventAction {
	return &DeleteEventAction{
		es: storage,
	}
}
