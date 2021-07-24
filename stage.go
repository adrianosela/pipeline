package pipeline

import (
	"log"
	"sync"
)

type transform func(interface{}) (interface{}, error)

type stage struct {
	name      string
	threads   int
	transform transform
	in        chan interface{}
	out       chan interface{}
}

func newStage(name string, threads int, action transform, in chan interface{}) *stage {
	return &stage{
		name:      name,
		threads:   threads,
		transform: action,
		in:        in,
		out:       make(chan interface{}),
	}
}

func (s *stage) run() {
	log.Printf("[STAGE:<%s>] Starting...\n", s.name)

	var wg sync.WaitGroup

	wg.Add(s.threads)
	for t := 0; t < s.threads; t++ {
		go func(threadId int) {
			defer wg.Done()
			for {
				received, ok := <-s.in
				if !ok {
					log.Printf("[STAGE:<%s-%d>] Input channel closed. Thread terminating...\n", s.name, threadId)
					break
				}

				processed, err := s.transform(received)
				if err != nil {
					log.Printf("[STAGE:<%s-%d>] Processing error: %s.\n", s.name, threadId, err)
					continue
				}

				s.out <- processed
			}
		}(t)
	}

	wg.Wait()
	log.Printf("[STAGE:<%s>] All threads terminated. Closing output channel...\n", s.name)
	close(s.out)
}
