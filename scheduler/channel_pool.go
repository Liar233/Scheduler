package scheduler

import (
	"errors"
	"fmt"
	"github.com/Liar233/Scheduler/model"
)

type ChannelPool struct {
	channels map[string]model.ChannelInterface
}

func (cp *ChannelPool) Add(channel model.ChannelInterface) error {
	if _, ok := cp.channels[channel.Name()]; ok {
		return errors.New(fmt.Sprintf("channel %s already exist", channel.Name()))
	}

	cp.channels[channel.Name()] = channel

	return nil
}

func (cp *ChannelPool) DispatchEvent(event *model.Event) error {
	channel, err := cp.channels[event.Channel]

	if err != true {
		return errors.New(fmt.Sprintf("channel %s not found", event.Channel))
	}

	return channel.Fire(event)
}

func NewChannelPool() *ChannelPool {
	return &ChannelPool{
		channels: make(map[string]model.ChannelInterface),
	}
}
