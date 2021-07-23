package pipeline

import "log"

type stage struct {
	name      string
	transform transform
	in        chan interface{}
	out       chan interface{}
}

func newStage(name string, action transform, in chan interface{}) *stage {
	return &stage{
		name:      name,
		transform: action,
		in:        in,
		out:       make(chan interface{}),
	}
}

func (s *stage) run() {
	log.Printf("[%s] Starting...\n", s.name)

	for {
		received, ok := <-s.in
		if !ok {
			log.Printf("[%s] Input channel closed. Quitting...\n", s.name)
			close(s.out)
			break
		}

		processed, err := s.transform(received)
		if err != nil {
			log.Printf("[%s] Processing error: %s.\n", s.name, err)
			continue
		}

		s.out <- processed
	}
}
