package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/JCKFinland/jck-connect/backend/pkg/logger"
)

// Logger logs every incoming HTTP request.
func Logger(log *logger.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		latency := time.Since(start)

		log.Info(
			"HTTP Request",
			logger.String("method", c.Request.Method),
			logger.String("path", c.Request.URL.Path),
			logger.Int("status", c.Writer.Status()),
			logger.String("client_ip", c.ClientIP()),
			logger.String("user_agent", c.Request.UserAgent()),
			logger.String("latency", latency.String()),
		)
	}
}
