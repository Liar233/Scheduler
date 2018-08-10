package main

import (
	"github.com/urfave/cli"
)

type CliAdapter struct {
	engine *cli.App
}

func NewCli(config *AppConfig) *CliAdapter {
	app :=  &CliAdapter{
		engine: cli.NewApp(),
	}

	app.engine.Name = "scheduler"
	app.engine.Usage = "Event scheduler service."
	app.engine.Version = "0.0.1"

	app.engine.Flags = []cli.Flag{
		cli.UintFlag{
			Name: "port, p",
			Usage: "`PORT` for http server",
			Value: 2337,
			EnvVar: "SCHEDULER_PORT",
		},
		cli.StringFlag{
			Name: "config, c",
			Usage: "Load configuration from `FILE`",
		},
		cli.BoolFlag{
			Name: "master, m",
			Usage: "Start as Master server",
		},
	}

	app.engine.Action = func(c *cli.Context) error {
		config.Port = c.GlobalUint("port")
		config.Master = c.Bool("master")

		return nil
	}

	return app
}

func (cli *CliAdapter) Run(arguments []string) error {
	return cli.engine.Run(arguments)
}
