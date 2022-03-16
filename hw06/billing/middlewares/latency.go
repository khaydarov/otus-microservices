package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func LatencyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// before request
		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
