package ratelimiter

import (
	"sync"
	"time"
)

// TokenBucket holds the rate limiting data per user.
type TokenBucket struct {
	tokens      int
	lastRefill  time.Time
	maxToken    int
	refillEvery time.Duration // tokens per second
	mu          sync.Mutex
}

// RateLimiter manages token buckets per user.
type RateLimiter struct {
	users       map[string]*TokenBucket
	maxToken    int
	refillEvery time.Duration
	mu          sync.Mutex
}

// NewRateLimiter ...
func NewRateLimiter(maxToken int, refillEvery time.Duration) *RateLimiter {
	return &RateLimiter{
		users:       make(map[string]*TokenBucket),
		maxToken:    maxToken,
		refillEvery: refillEvery,
	}
}

// Allow ...
func (rl *RateLimiter) Allow(userID string) bool {
	rl.mu.Lock()

	tb, exist := rl.users[userID]
	if !exist {
		tb = &TokenBucket{
			tokens:      rl.maxToken,
			maxToken:    rl.maxToken,
			lastRefill:  time.Now(),
			refillEvery: rl.refillEvery,
		}
		rl.users[userID] = tb
	}
	rl.mu.Unlock()

	return tb.Allow()
}

// Operations ...
type Operations interface {
	Allow(userID string) bool
}

// Helper function
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Allow returns true if the request is allowed, false if rate-limited.
func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)

	// Calculate how many tokens to add
	newTokens := int(elapsed / tb.refillEvery)

	if newTokens > 0 {
		tb.tokens = min(tb.maxToken, tb.tokens+newTokens)
		tb.lastRefill = now
	}

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}

	return false
}
