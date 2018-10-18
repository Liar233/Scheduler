package main

import (
	"go.uber.org/fx"
	"github.com/Liar233/Scheduler/cli"
	"github.com/Liar233/Scheduler/config"
	"github.com/Liar233/Scheduler/scheduler"
	"os"
)

// Init new application
func NewApplication() *fx.App {
	return fx.New(
		fx.Provide(
			config.NewAppConfig,
			cli.NewCli,
			scheduler.NewChannelPool,
			scheduler.NewEventLoop,
		),
		fx.Invoke(func(cli *cli.CliAdapter) {
			cli.Run(os.Args)
		}),
	)
}
