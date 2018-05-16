package main

import (
	"errors"
	"fmt"
)

type ChannelInterface interface {
	Fire(e *Event) error
	Name() string
}

type Channel struct {
	name string
}

func (c *Channel) Fire(e *Event) error {
	return nil
}

func (c *Channel) Name() string {
	return c.name
}

func NewChannel(name string) *Channel {
	return &Channel{
		name: name,
	}
}

type ChannelPool struct {
	channels map[string]ChannelInterface
}

func (cp *ChannelPool) Add(channel ChannelInterface) error {
	if _, ok := cp.channels[channel.Name()]; ok {
		return errors.New(fmt.Sprintf("channel %s already exist", channel.Name()))
	}

	cp.channels[channel.Name()] = channel

	return nil
}

func (cp *ChannelPool) DispatchEvent(event *Event) error {
	channel, err := cp.channels[event.Channel]

	if err != true {
		return errors.New(fmt.Sprintf("channel %s not found", event.Channel))
	}

	return channel.Fire(event)
}

func NewChannelPool() ChannelPool {
	return ChannelPool{
		channels: make(map[string]ChannelInterface),
	}
}
