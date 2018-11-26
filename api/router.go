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
	List        *action.EventListAction
	EventLoop   *action.EventLoopAction
	ChannelList *action.ChannelListAction
}

func ParametriseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		q := r.URL.Query()

		for key, val := range vars {
			q.Add(key, val)
		}

		r.URL.RawQuery = q.Encode()

		next.ServeHTTP(w, r)
	})
}

func NewRouterAdapter(actions ActionParams) *RouterAdapter {
	r := &RouterAdapter{
		*mux.NewRouter(),
	}

	r.Use(ParametriseMiddleware)
	r.Handle("/health-check", actions.HealthCheck).Methods(http.MethodGet)
	r.Handle("/event/{id}", actions.Get).Methods(http.MethodGet)
	r.Handle("/event/{id}", actions.Delete).Methods(http.MethodDelete)
	r.Handle("/event", actions.Create).Methods(http.MethodPost)
	r.Handle("/event", actions.List).Methods(http.MethodGet)
	r.Handle("/event-loop", actions.EventLoop).Methods(http.MethodGet)
	r.Handle("/channel", actions.ChannelList).Methods(http.MethodGet)
	return r
}
