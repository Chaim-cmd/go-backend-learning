package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		//执行 handler
		c.Next()

		//handler执行完后记录日志
		duration := time.Since(start)
		request_id, _ := c.Get("request_id")

		status := c.Writer.Status()

		//根据状态码加颜色（终端可见）
		statusColor := colorForStatus(status)
		//重置样式
		reset := "\033[0m"

		fmt.Printf("[GIN] %s | %s%d%s | %12v |  %-7s %s | request_id=%s\n",
			time.Now().Format("2006/01/02 15:04:05"),
			statusColor, status, reset,
			duration,
			c.Request.Method,
			c.Request.URL.Path,
			&request_id,
		)
	}
}
func colorForStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "\033[32m" //绿色
	case code >= 400 && code < 500:
		return "\033[33m" //黄色
	default:
		return "\033[31m" //红色
	}
}
