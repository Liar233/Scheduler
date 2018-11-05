package scheduler

import (
	"testing"
	"time"
	"github.com/Liar233/Scheduler/model"
)

type ChannelStub struct {
	name    string
	Payload string
}

func (c *ChannelStub) Fire(e *model.Event) error {
	c.Payload = e.Payload

	return nil
}

func (c *ChannelStub) Name() string {
	return c.name
}

func TestEventLoop_Start(t *testing.T) {
	channels := NewChannelPool()
	eventLoop := NewEventLoop(channels)

	eventLoop.Start()

	if eventLoop.running != true {
		t.Fail()
		t.Error("Fail to start EventLoop")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
			t.Error("Fail. EventLoop start twice")
		}
	}()

	eventLoop.Start()
}

func TestEventLoopRun(t *testing.T) {
	channelPool := NewChannelPool()

	channel := &ChannelStub{
		name: "test_channel",
	}

	channelPool.Add(channel)

	now := time.Now()
	fireTime := now.Add(time.Duration(1) * time.Second)

	event := &model.Event{
		ID:       "TestEvent",
		Channel:  "test_channel",
		Payload:  "Test",
		FireTime: fireTime,
	}

	defer func() {
		if channel.Payload != "Test" {
			t.Fail()
			t.Error("Not valid event payload!")
		}
	}()

	eventLoop := NewEventLoop(channelPool)
	eventLoop.Start()

	eventLoop.Push(event)

	time.Sleep(time.Duration(2) * time.Second)
}
