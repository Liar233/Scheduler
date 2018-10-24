package main

import (
	"go.uber.org/fx"
	"github.com/Liar233/Scheduler/config"
	"github.com/Liar233/Scheduler/scheduler"
	"github.com/Liar233/Scheduler/api"
	"github.com/Liar233/Scheduler/api/action"
	"github.com/Liar233/Scheduler/storage"
)

// Init new application
func NewApplication() *fx.App {
	return fx.New(
		fx.Provide(
			config.NewAppConfig,
			action.NewHealthCheck,
			api.NewRouterAdapter,
			api.NewWebServer,
			scheduler.NewChannelPool,
			scheduler.NewEventLoop,
			storage.NewEventStorage,
		),
		fx.Invoke(func(wsa *api.WebServerAdapter) {
			go wsa.ListenAndServe()
		}),
	)
}
