package bug1

import (
	"sync"
)

// Counter stores a count.
type Counter struct {
	n int64
	mux sync.Mutex
}

// Inc increments the count in the Counter.
func (c *Counter) Inc() {
	c.mux.Lock()
	c.n++
	c.mux.Unlock()
}

//not concurrency safe because 2 pointers to the same counter could be passed
//plan: make a mutex 