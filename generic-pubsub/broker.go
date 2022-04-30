package pubsub

import (
	"sync"
)

type Broker[M any] struct {
	subscribers map[*Subscriber[M]]bool
	lastMsg     M
	mut         sync.RWMutex // to avoid reading subscribers while it is being updated
}

func NewBroker[M any](initMsg M) *Broker[M] {
	broker := &Broker[M]{}
	broker.subscribers = make(map[*Subscriber[M]]bool)
	broker.lastMsg = initMsg
	return broker
}

func (b *Broker[M]) NewSubscriber() *Subscriber[M] {
	b.mut.Lock()
	defer b.mut.Unlock()
	s := NewSubscriber(b.lastMsg)
	b.subscribers[s] = true
	go func() { s.Signal(b.lastMsg) }()
	return s
}

func (b *Broker[M]) RemoveSubscriber(s *Subscriber[M]) {
	b.mut.Lock()
	defer b.mut.Unlock()
	delete(b.subscribers, s)
}

func (b *Broker[M]) Broadcast(msg M) {
	b.mut.RLock()
	defer b.mut.RUnlock()
	b.lastMsg = msg
	for s := range b.subscribers {
		go func(s *Subscriber[M]) { s.Signal(msg) }(s)
	}
}
