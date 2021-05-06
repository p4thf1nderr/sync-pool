package main

import (
	"fmt"

	"github.com/p4thfinderr/sync-pool/counter"
)

func main() {
	c := &counter.Counter{}

	for i := 0; i < 100000; i++ {
		c.Increment()
	}

	fmt.Printf("%+v", c)
}
