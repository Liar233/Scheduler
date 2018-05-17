package main

import "testing"

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
