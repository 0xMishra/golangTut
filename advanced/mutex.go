package advanced

import (
	"fmt"
	"sync"
)

type counter struct {
	mu    sync.Mutex
	count int
}

func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func mutex() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	numGoroutines := 5
	wg.Add(numGoroutines)
	increment := func() {
		defer wg.Done()
		for range 1000 {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}

	for range numGoroutines {
		go increment()
	}

	wg.Wait()
	fmt.Printf("final value of counter: %d\n", counter)
}

func mutexWithStruct() {
	var wg sync.WaitGroup
	ctr := &counter{}

	numGoroutines := 10

	// wg.Add(numGoroutines)

	for range numGoroutines {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for range 1000 {
				ctr.increment()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("final counter value %d\n", ctr.getValue())
}
