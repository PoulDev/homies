package middlewares

// provides rate limiting for the homies API
// based on https://github.com/gin-gonic/examples/tree/master/ratelimiter

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit" // thx uber?
)

func GetLimiter(limit ratelimit.Limiter) gin.HandlerFunc {
	prev := time.Now()
	return func(ctx *gin.Context) {
		now := limit.Take()
		log.Print(now.Sub(prev))
		prev = now
	}
}


