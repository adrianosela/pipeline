package pipeline

import "log"

type stage struct {
	name   string
	action action
	in     chan interface{}
	out    chan interface{}
}

type action func(interface{}) (interface{}, error)

func newStage(name string, action action, in chan interface{}) *stage {
	return &stage{
		name:   name,
		action: action,
		in:     in,
		out:    make(chan interface{}),
	}
}

func (s *stage) run() {
	for {
		received, ok := <-s.in
		if !ok {
			log.Printf("[%s] input channel closed... quitting.\n", s.name)
			close(s.out)
			break
		}

		processed, err := s.action(received)
		if err != nil {
			log.Printf("[%s] processing error: %s\n", s.name, err)
			continue
		}

		s.out <- processed
	}
}
