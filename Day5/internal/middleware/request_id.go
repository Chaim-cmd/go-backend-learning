package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Request_id() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.GetHeader("X-Request-ID")
		if id == "" {
			id = uuid.New().String()
		}
		ctx.Set("request_id", id)
		ctx.Header("X-Request-ID", id)
		ctx.Next()
	}
}
