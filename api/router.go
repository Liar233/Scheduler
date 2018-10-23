package api

import (
	"github.com/gorilla/mux"
	"github.com/Liar233/Scheduler/api/action"
)

type RouterAdapter struct {
	mux.Router
}

func NewRouterAdapter(hc *action.HealthCheck) *RouterAdapter {
	r := &RouterAdapter{
		*mux.NewRouter(),
	}

	r.Handle("/health-check", hc).Methods("GET")

	return r
}
