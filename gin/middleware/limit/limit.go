package limit

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// MaxAllowed  router.Use(limit.MaxAllowed(20))
func MaxAllowed(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()

	}
}

//Limit function is gin middleware to limit current requests
func Limit(max int) gin.HandlerFunc {
	if max <= 0 {
		// log.Panic("max must be more than 0")
		max = 100
	}
	sema := make(chan struct{}, max)
	return func(c *gin.Context) {
		var called, fulled bool
		defer func() {
			if called == false && fulled == false {
				<-sema
			}
			if r := recover(); r != nil { // We don't handle panic
				panic(r)
			}
		}()

		select {
		case sema <- struct{}{}:
			c.Next()
			called = true
			<-sema
		default:
			fulled = true
			c.Status(http.StatusBadGateway)
		}
	}
}

var (
	ErrorLimitExceeded = errors.New("Limit exceeded")
)

// Drops (HTTP status 429) the request if the limit is reached.
func Limit(maxEventsPerSec float64, maxBurstSize int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(maxEventsPerSec), maxBurstSize)

	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()
			return
		}

		// Limit reached
		c.Error(ErrorLimitExceeded)
		c.AbortWithStatus(429)
	}
}
