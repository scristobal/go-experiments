package pubsub

import (
	"sync"
)

type Subscriber[S any] struct {
	state  S
	states chan S
	mut    sync.RWMutex
}

func NewSubscriber[S any](state S) *Subscriber[S] {
	subscriber := &Subscriber[S]{}
	subscriber.states = make(chan S)
	return subscriber
}

func (s *Subscriber[S]) Signal(state S) {
	s.mut.Lock()
	s.state = state
	s.mut.Unlock()

	s.states <- state
}

func (s *Subscriber[S]) LastRead() S {
	s.mut.RLock()
	state := s.state
	defer s.mut.RUnlock()
	return state
}

func (s *Subscriber[S]) Read() S {
	return <-s.states
}
