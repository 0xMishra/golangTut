package advanced

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	count int64
}

func (ac *AtomicCounter) increment() {
	atomic.AddInt64(&ac.count, 1)
}

func (ac *AtomicCounter) getValue() int64 {
	return atomic.LoadInt64(&ac.count)
}

func atomicCounterFunc() {
	var wg sync.WaitGroup
	numGoroutines := 10

	counter := &AtomicCounter{}

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("final counter value: %d\n", counter.getValue())
}
