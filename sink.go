package pipeline

import "log"

type publish func(interface{}) error

type sink struct {
	name   string
	action publish
	in     chan interface{}
	done   chan interface{}
}

func newSink(name string, action publish, in chan interface{}) *sink {
	return &sink{
		name:   name,
		action: action,
		in:     in,
		done:   make(chan interface{}),
	}
}

func (s *sink) run() {
	log.Printf("[%s] Starting...\n", s.name)

	for {
		received, ok := <-s.in
		if !ok {
			log.Printf("[%s] Input channel closed. Quitting...\n", s.name)
			s.done <- true
			break
		}

		if err := s.action(received); err != nil {
			log.Printf("[%s] Committing error: %s.\n", s.name, err)
			continue
		}
	}
}
