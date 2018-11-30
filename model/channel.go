package model

import "github.com/Liar233/Scheduler/config"

type ChannelInterface interface {
	Fire(e *Event) error
	Name() string
	Init(config *config.ChannelConfig, name string)
	Connect() error
}
