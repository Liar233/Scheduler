package main

import (
	"github.com/urfave/cli"
	"fmt"
)

type CliAdapter struct {
	engine *cli.App
}

func NewCli() *CliAdapter {
	app := cli.NewApp()

	app.Name = "scheduler"
	app.Usage = "Event scheduler service."
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
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
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("Started...")
		return nil
	}

	return &CliAdapter{
		engine: app,
	}
}

func (cli *CliAdapter) Run(arguments []string) error {
	return cli.engine.Run(arguments)
}
