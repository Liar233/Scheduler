package api

import (
	"github.com/gorilla/mux"
	"github.com/Liar233/Scheduler/api/action"
	"go.uber.org/fx"
	"net/http"
)

type RouterAdapter struct {
	mux.Router
}

type ActionParams struct {
	fx.In

	HealthCheck *action.HealthCheck
	Create      *action.CreateEventAction
	Get         *action.GetEventAction
	Delete      *action.DeleteEventAction
}

func NewRouterAdapter(actions ActionParams) *RouterAdapter {
	r := &RouterAdapter{
		*mux.NewRouter(),
	}

	r.Handle("/health-check", actions.HealthCheck).Methods(http.MethodGet)
	r.Handle("/event/{id:[0-9]+}", actions.Get).Methods(http.MethodGet)
	r.Handle("/event", actions.Create).Methods(http.MethodPost)
	r.Handle("/event", actions.Delete).Methods(http.MethodDelete)

	return r
}
