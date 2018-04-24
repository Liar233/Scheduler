package main

type ChannelInterface interface {
	Process(e *Event) error
	Name() string
}

type Channel struct {
	name   string
	Events []*Event
}

func (c *Channel) Process(e *Event) error {
	panic("implement me")
}

func (c *Channel) Name() string {
	panic("implement me")
}

func NewChannel(name string) *Channel {
	return &Channel{
		name:   name,
		Events: make([]*Event, 0),
	}
}
