package model

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
