package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	capacity   int
	tokens     float64
	rate       float64
	lastRefill time.Time
	mu         sync.Mutex
}

func NewTokenBucket(capacity int, rate float64) *TokenBucket {

	return &TokenBucket{
		capacity:   capacity,
		rate:       rate,
		tokens:     float64(capacity),
		lastRefill: time.Now(),
	}
}

func (tb *TokenBucket) Allow() {

	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()

	if tb.tokens >= 1 { // token is float, so partial request not allowed
		tb.tokens--
		fmt.Println("Request Allowed!")
	} else {
		fmt.Println("Request Rejected!")
	}

}

func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()

	newTokens := elapsed * tb.rate

	tb.tokens += newTokens

	if tb.tokens > float64(tb.capacity) {
		tb.tokens = float64(tb.capacity)
	}

	tb.lastRefill = time.Now()
}

func main() {

	myTokenBucket := NewTokenBucket(5, 1)
	i := 0
	for {

		fmt.Println(i)
		myTokenBucket.Allow()
		time.Sleep(10 * time.Millisecond)
		i++
	}
}
