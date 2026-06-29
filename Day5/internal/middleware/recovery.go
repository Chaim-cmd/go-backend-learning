package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/Chaim-cmd/go-backend-learning.git/Day2/pkg/response"
	errorsa "github.com/Chaim-cmd/go-backend-learning.git/Day5/pkg/errorsa"
	"github.com/Chaim-cmd/go-backend-learning.git/Day5/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Recovery 替换gin.Recovery(), panic 时打结构化日志 + 返回统一 JSON
func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				//打完整堆到日志，但不暴露给客户端
				logger.Error("panic recovered",
					zap.Any("error", r),
					zap.String("stack", string(debug.Stack())),
					zap.String("path", ctx.Request.URL.Path),
				)

				response.Fail(ctx, http.StatusInternalServerError, errorsa.Internal("服务器内部错误", nil).Error())
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
