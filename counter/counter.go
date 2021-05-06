package counter

import (
	"fmt"
	"sync"
)

type Counter struct {
	A      int
	B      int
	buffer map[string]string
}

func NewCounter() *Counter {
	return &Counter{
		A:      0,
		B:      0,
		buffer: make(map[string]string),
	}
}

var counterFree = sync.Pool{
	New: func() interface{} { return new(Counter) },
}

func NewCounterPool() *Counter {

	c := counterFree.Get().(*Counter)
	c.A = 0
	c.B = 0
	c.buffer = make(map[string]string)

	return c
}

func (c *Counter) Increment() {
	c.A++
	c.B++

	for i := 0; i < 100; i++ {
		c.buffer[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
	}

}
