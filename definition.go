package main

import (
	"go.uber.org/fx"
	"github.com/Liar233/Scheduler/config"
	"github.com/Liar233/Scheduler/api"
	"github.com/Liar233/Scheduler/storage"
	"github.com/Liar233/Scheduler/scheduler"
	"github.com/Liar233/Scheduler/api/action"
)

// Init new application
func NewApplication() *fx.App {
	return fx.New(
		fx.Provide(
			config.NewAppConfig,

			storage.NewEventStorage,

			scheduler.NewChannelPool,
			scheduler.NewEventLoop,

			action.NewHealthCheck,
			action.NewGetEventAction,
			action.NewCreateEventAction,
			action.NewDeleteEventAction,

			api.NewRouterAdapter,
			api.NewWebServer,
		),
		fx.Invoke(func(wsa *api.WebServerAdapter) {
			go wsa.ListenAndServe()
		}),
	)
}
