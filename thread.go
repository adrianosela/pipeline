package pipeline

import "sync"

// threaded runs the passed function on n threads
// and blocks until all threads have finished
// NOTE: fn must call Done on wait group when finished
func threaded(n int, fn func(threadId int, wg *sync.WaitGroup)) {
	var wg sync.WaitGroup

	wg.Add(n)
	for threadId := 0; threadId < n; threadId++ {
		go fn(threadId, &wg)
	}
	wg.Wait()
}
