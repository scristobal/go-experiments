package pubsub

import (
	"testing"
)

func TestBroadcast(t *testing.T) {

	broker := NewBroker("")

	subscriber := broker.NewSubscriber()

	broker.Broadcast("test single")

	if msg := subscriber.Read(); msg != "test single" {
		t.Errorf("single subscriber expected 'test single' got %s", msg)
	}

}

func TestRead(t *testing.T) {

	broker := NewBroker("")

	subscriber1 := broker.NewSubscriber()
	subscriber2 := broker.NewSubscriber()

	broker.Broadcast("test multiple")

	if msg := subscriber1.Read(); msg != "test multiple" {
		t.Errorf("subscriber1 expected 'test multiple' got %s", msg)
	}

	if msg := subscriber2.Read(); msg != "test multiple" {
		t.Errorf("subscriber2 expected 'test multiple' got %s", msg)
	}

}

func TestLastRead(t *testing.T) {

	broker := NewBroker("")
	subscriber := broker.NewSubscriber()

	broker.Broadcast("test last message")

	if msg := subscriber.Read(); msg != "test last message" {
		t.Errorf("subscriber expected 'test last message' got %s", msg)
	}

	broker.Broadcast("new message")

	if msg := subscriber.LastRead(); msg != "test last message" {
		t.Errorf("subscriber expected 'test last message' got %s", msg)
	}

	subscriber.Read()

	if msg := subscriber.LastRead(); msg != "new message" {
		t.Errorf("subscriber expected 'new message' got %s", msg)
	}

}
