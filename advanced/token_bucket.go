package advanced

import (
	"fmt"
	"time"
)

type RateLimiterTB struct {
	tokens     chan struct{}
	refillTime time.Duration
}

func NewRateLimiterTB(rateLimit int, refillTime time.Duration) *RateLimiterTB {
	rl := &RateLimiterTB{
		tokens:     make(chan struct{}, rateLimit),
		refillTime: refillTime,
	}
	for range rateLimit {
		rl.tokens <- struct{}{}
	}
	go rl.startRefill()
	return rl
}

func (rl *RateLimiterTB) startRefill() {
	ticker := time.NewTicker(rl.refillTime)

	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			select {
			case rl.tokens <- struct{}{}:
			default:
			}
		}
	}
}

func (rl *RateLimiterTB) allowTB() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

func tokenBucketFunc() {
	rateLimiter := NewRateLimiterTB(5, time.Second)
	for range 10 {
		if rateLimiter.allowTB() {
			fmt.Println("request allowed")
		} else {
			fmt.Println("request denied")
		}
		time.Sleep(time.Duration(200) * time.Millisecond)
	}
}
