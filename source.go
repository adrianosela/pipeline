package pipeline

import (
	"errors"
	"log"
)

// ErrorSourceFinished should be returned by the ingest
// action when the source has no input data left
var ErrorSourceFinished = errors.New("source finished")

type ingest func() (interface{}, error)

type source struct {
	name    string
	threads int
	action  ingest
	out     chan interface{}
}

func newSource(name string, threads int, chanSize int, action ingest) *source {
	return &source{
		name:    name,
		threads: threads,
		action:  action,
		out:     make(chan interface{}, chanSize),
	}
}

func (s *source) run() {
	log.Printf("[SOURCE:<%s>] Starting...\n", s.name)

	threaded(s.threads, func(threadID int) {
		for {
			recvd, err := s.action()
			if err != nil {
				if err == ErrorSourceFinished {
					log.Printf("[SOURCE:<%s-%d>] Source finished. Thread terminating...\n", s.name, threadID)
					break
				}

				log.Printf("[SOURCE:<%s-%d>] Ingestion error: %s.\n", s.name, threadID, err)
				continue
			}

			s.out <- recvd
		}
	})

	log.Printf("[SOURCE:<%s>] All threads terminated. Closing output channel...\n", s.name)
	close(s.out)
}
