package action

import (
	"github.com/Liar233/Scheduler/storage"
	"net/http"
)

type GetEventAction struct {
	es *storage.EventStorage
}

func (a *GetEventAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	event, err := a.es.Get(id)

	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	WriteSuccessResponse(w, event)
}

func NewGetEventAction(storage *storage.EventStorage) *GetEventAction {
	return &GetEventAction{
		es: storage,
	}
}
