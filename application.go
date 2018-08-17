package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

type Application struct {
	Config      AppConfig
	Cli         CliAdapter
	ApiServer   WebServerAdapter
	ChannelPool ChannelPool

	wg sync.WaitGroup
}

// Bootstrap application
// Make application config and trying to fill it
func (app *Application) Bootstrap(arguments []string) error {
	app.Config = AppConfig{}
	app.Cli = NewCli(&app.Config)

	if err := app.Cli.Run(arguments); err != nil {
		return err
	}

	app.ChannelPool = NewChannelPool()

	for name := range app.Config.Channels {
		channel := NewChannel(name)
		if err := app.ChannelPool.Add(channel); err != nil {
			fmt.Errorf("Fail to add channel `%s` with error: %s\n", name, err)
			continue
		}
	}

	return nil
}

// Start server api server and event loop
func (app *Application) Run() error {
	app.wg.Add(1)
	go app.HandleClose()

	app.ApiServer = NewWebServer(&app.Config, &app.wg)

	if err:= app.ApiServer.ListenAndServe(); err != nil {
		return err
	}

	go app.HandleClose()

	println("Stopped!")

	return nil
}

// Waiting for an Interrupt signal to close goroutines
func (app *Application) HandleClose() {
	signchan := make(chan os.Signal, 1)
	signal.Notify(signchan, os.Interrupt, os.Kill)

	<-signchan

	println("Stopping...")

	app.ApiServer.Close()
	app.wg.Wait()

	return
}
