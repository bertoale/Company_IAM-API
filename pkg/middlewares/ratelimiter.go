package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type clientData struct {
	lastRequest  time.Time
	requestCount int
}

type RateLimiter struct {
	requests int
	window   time.Duration
	clients  map[string]*clientData
	mu       sync.Mutex
}

func NewRateLimiter(requests int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: requests,
		window:   window,
		clients:  make(map[string]*clientData),
	}
}

func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		rl.mu.Lock()
		data, exists := rl.clients[clientIP]
		now := time.Now()
		if !exists || now.Sub(data.lastRequest) > rl.window {
			// Reset window
			rl.clients[clientIP] = &clientData{lastRequest: now, requestCount: 1}
			rl.mu.Unlock()
			c.Next()
			return
		}
		if data.requestCount >= rl.requests {
			rl.mu.Unlock()
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"success": false,
				"message": "Rate limit exceeded. Please try again later.",
			})
			return
		}
		data.requestCount++
		data.lastRequest = now
		rl.mu.Unlock()
		c.Next()
	}
}
