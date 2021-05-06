package counter_test

import (
	"testing"

	"github.com/p4thfinderr/sync-pool/counter"
)

func BenchmarkWithoutPool(b *testing.B) {
	c := counter.NewCounter()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			// b.StartTimer()
			c.Increment()
			// b.StopTimer()
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	c := counter.NewCounterPool()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			// b.StartTimer()
			c.Increment()
			// b.StopTimer()
		}
	}
}
