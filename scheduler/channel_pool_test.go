package scheduler

import (
	"testing"
	"github.com/Liar233/Scheduler/model"
)

func TestChannelPool_Add(t *testing.T) {
	channel1 := model.NewChannel("channel1")
	channel2 := model.NewChannel("channel2")

	channelPool := NewChannelPool()

	channelPool.Add(channel1)
	channelPool.Add(channel2)

	if _, ok := channelPool.channels["channel1"]; !ok {
		t.Fail()
		t.Error("Channel1 not found in ChannelPool")
	}

	if _, ok := channelPool.channels["channel2"]; !ok {
		t.Fail()
		t.Error("Channel2 not found in ChannelPool")
	}
}

func TestChannelPool_DispatchEvent_Fire(t *testing.T) {
	event1 := model.Event{
		Channel: "channel1",
	}

	event2 := model.Event{
		Channel: "channel3",
	}

	channel1 := model.NewChannel("channel1")
	channel2 := model.NewChannel("channel2")

	channelPool := NewChannelPool()

	channelPool.Add(channel1)
	channelPool.Add(channel2)

	var err error

	err = channelPool.DispatchEvent(&event1)

	if err != nil {
		t.Fail()
		t.Error("Fail to dispatch event")
		t.Error(err)
	}

	err = channelPool.DispatchEvent(&event2)

	if err == nil {
		t.Fail()
		t.Error("Event dispatched on none existence channel")
	}

}
