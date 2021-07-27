package pipeline

import (
	"log"
	"sync"
)

type publish func(interface{}) error

type sink struct {
	name    string
	threads int
	action  publish
	in      chan interface{}
	done    chan interface{}
}

func newSink(name string, threads int, action publish, in chan interface{}) *sink {
	return &sink{
		name:    name,
		threads: threads,
		action:  action,
		in:      in,
		done:    make(chan interface{}),
	}
}

func (s *sink) run() {
	log.Printf("[SINK:<%s>] Starting...\n", s.name)

	threaded(s.threads, func(threadId int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			received, ok := <-s.in
			if !ok {
				log.Printf("[SINK:<%s-%d>] Input channel closed. Quitting...\n", s.name, threadId)
				break
			}

			if err := s.action(received); err != nil {
				log.Printf("[SINK:<%s-%d>] Committing error: %s.\n", s.name, threadId, err)
				continue
			}
		}
	})

	log.Printf("[SINK:<%s>] All threads terminated. Closing output channel...\n", s.name)
	s.done <- true
}
