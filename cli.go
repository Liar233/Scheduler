package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"time"
)

type CliAdapter struct {
	engine *cli.App
}

func NewCli(config *AppConfig) CliAdapter {
	app := CliAdapter{
		engine: cli.NewApp(),
	}

	app.engine.Name = "scheduler"
	app.engine.Usage = "Event scheduler service."
	app.engine.Version = "0.0.1"

	app.engine.Flags = []cli.Flag{
		cli.UintFlag{
			Name:   "port, p",
			Usage:  "`PORT` for http server",
			EnvVar: "SCHEDULER_PORT",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	app.engine.Action = func(c *cli.Context) error {
		configFilePath := c.String("config")

		if configFilePath == "" {
			return fmt.Errorf("--config parameter is empty")
		}

		validationErrors := LoadYamlConfig(configFilePath, config)

		if len(validationErrors) > 0 {
			for _, err := range validationErrors {
				fmt.Fprintf(os.Stderr, "%s %s \n", time.Now().Format("2006/01/02 15:04:05"), err)
			}

			return fmt.Errorf("scheduler can't start")
		}

		if port := c.Uint("port"); port != 0 {
			config.Port = c.Uint("port")
		}

		return nil
	}

	return app
}

func (cli *CliAdapter) Run(arguments []string) error {
	return cli.engine.Run(arguments)
}
