package advanced

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiterWC struct {
	mu        sync.Mutex
	count     int
	limit     int
	window    time.Duration
	resetTime time.Time
}

func NewRateLimiterWC(limit int, window time.Duration) *RateLimiterWC {
	return &RateLimiterWC{
		limit:  limit,
		window: window,
	}
}

func (rl *RateLimiterWC) allowWC() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.window)
		rl.count = 0
	}

	if rl.count < rl.limit {
		rl.count++
		return true
	}
	return false
}

func windoCounter() {
	rateLimiter := NewRateLimiterWC(3, time.Second)

	for range 10 {
		if rateLimiter.allowWC() {
			fmt.Println("request allowed")
		} else {
			fmt.Println("request denied")
		}
		time.Sleep(time.Duration(200) * time.Millisecond)
	}
}
