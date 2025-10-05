package advanced

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct {
	capacity int
	leakRate time.Duration
	tokens   int
	lastLeak time.Time
	mu       sync.Mutex
}

func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		tokens:   capacity,
		lastLeak: time.Now(),
	}
}

func (lb *LeakyBucket) allowLB() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	now := time.Now()
	elapsedTime := now.Sub(lb.lastLeak)

	tokensToAdd := int(elapsedTime / lb.leakRate)
	lb.tokens += tokensToAdd

	if lb.tokens > lb.capacity {
		lb.tokens = lb.capacity
	}

	lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate)
	/*
		lb.lastLeak = lb.lastLeak.Add(elapsedTime) - WRONG APPROACH, in this manner the lastLeak will always keep increasing
		even if we haven't added any tokens- as it only depends on elapsedTime and not on tokensToAdd or leakRate
	*/

	if lb.tokens > 0 {
		lb.tokens--
		return true
	}

	return false
}

func LeakyBucketFunc() {
	leakyBucket := NewLeakyBucket(5, time.Duration(500)*time.Millisecond)
	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)

		go func() {
			defer wg.Done()
			if leakyBucket.allowLB() {
				fmt.Println("request allowed")
			} else {
				fmt.Println("request denied")
			}
			time.Sleep(time.Duration(200) * time.Millisecond)
		}()
	}

	wg.Wait()
}
