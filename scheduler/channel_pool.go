package scheduler

import (
	"errors"
	"fmt"
	"github.com/Liar233/Scheduler/model"
	"github.com/Liar233/Scheduler/config"
	"github.com/Liar233/Scheduler/drivers"
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

func (cp *ChannelPool) GetList() []string {
	var channelList []string

	for name := range cp.channels {
		channelList = append(channelList, name)
	}

	return channelList
}

func NewChannelPool(conf *config.AppConfig) *ChannelPool {
	cp := &ChannelPool{
		channels: make(map[string]model.ChannelInterface),
	}

	for name, channelConf := range conf.Channels {
		channel := drivers.NewTcpChannel(&channelConf, name)
		cp.Add(channel)
	}

	return cp
}
