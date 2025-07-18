package atomic

import (
	"sync"
	"sync/atomic"
)

type Counter struct {
	count int64
	mu    sync.Mutex
}

func (c *Counter) Increase() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count = 0
}

func (c *Counter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func (c *Counter) AtomicIncrease() {
	atomic.AddInt64(&c.count, 1)
}
