package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestID 为每个请求生成唯一 ID，注入 context 和响应头
func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//优先使用客户端传过来的 ID
		requestID := ctx.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		//存入context,后续handler和日志都能拿到
		ctx.Set("request_id", requestID)

		//写入响应头,客户端能看到
		ctx.Header("X-Request-ID", requestID)

		ctx.Next()
	}
}
