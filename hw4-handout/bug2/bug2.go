package bug2

import (
	"sync"
)

func bug2(n int, foo func(int) int, out chan int) {
	var wg sync.WaitGroup // i is being written to multiple times
	for i := 0; i < n; i++ {
		wg.Add(1) //wait group adds goroutine
		go func(i int) {
			out <- foo(i) //channel receiver
			wg.Done() // wait group calls done when finished
		}(i)
	}
	wg.Wait() // waitgroup waits until all goroutines have finished running
	close(out) //channel closed
}
