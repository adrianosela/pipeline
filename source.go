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
	name   string
	action ingest
	out    chan interface{}
}

func newSource(name string, action ingest) *source {
	return &source{
		name:   name,
		action: action,
		out:    make(chan interface{}),
	}
}

func (s *source) run() {
	log.Printf("[%s] Starting...\n", s.name)

	for {
		recvd, err := s.action()
		if err != nil {
			if err == ErrorSourceFinished {
				log.Printf("[%s] Source finished. Closing output channel...\n", s.name)
				close(s.out)
				break
			}

			log.Printf("[%s] Ingestion error: %s.\n", s.name, err)
			continue
		}

		s.out <- recvd
	}
}
