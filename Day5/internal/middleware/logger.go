package middleware

import (
	"time"

	"github.com/Chaim-cmd/go-backend-learning.git/Day5/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Logger 用zap打印结构化日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		request_id, _ := c.Get("request_id")
		status := c.Writer.Status()

		//根据状态码选日志级别
		fields := []zap.Field{
			zap.String("request_id", request_id.(string)),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", status),
			zap.Duration("latency", time.Since(start)),
			zap.String("client_ip", c.ClientIP()),
		}
		if status >= 500 {
			logger.Error("server error", fields...)
		} else if status >= 400 {
			logger.Warn("client Error", fields...)
		} else {
			logger.Info("request", fields...)
		}

	}
}
