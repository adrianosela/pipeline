package pipeline

import "sync"

// threaded runs the passed function on n threads
// and blocks until all threads have finished
func threaded(n int, fn func(threadID int)) {
	var wg sync.WaitGroup

	wg.Add(n)
	for t := 0; t < n; t++ {
		go func(threadID int) {
			defer wg.Done()
			fn(threadID)
		}(t)
	}
	wg.Wait()
}
