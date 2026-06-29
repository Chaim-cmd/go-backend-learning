package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code      int         `json:"code"`
	Msg       string      `jsonL:"msg"`
	Data      interface{} `json:"data"`
	RequestID string      `json:"request_id,omitempty"`
}

// Success 返回成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:      0,
		Msg:       "success",
		Data:      data,
		RequestID: getRequestID(c),
	})
}

// Fail 返回失败响应
func Fail(c *gin.Context, httpCode int, msg string) {
	c.JSON(httpCode, Response{
		Code:      httpCode,
		Msg:       msg,
		Data:      nil,
		RequestID: getRequestID(c),
	})
}

func getRequestID(c *gin.Context) string {
	if id, exists := c.Get("request_id"); exists {
		return id.(string)
	}
	return ""
}
