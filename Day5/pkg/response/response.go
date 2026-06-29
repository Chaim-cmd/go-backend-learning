package response

import (
	"net/http"

	errorsa "github.com/Chaim-cmd/go-backend-learning.git/Day5/pkg/errorsa"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
	Request_ID string      `json:"request_id,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 0, Msg: "success", Data: data,
		Request_ID: getRequestID(c),
	})
}

func Fail(c *gin.Context, httpCode int, msg interface{}) {
	var code1 int
	var message string
	switch v := msg.(type) {
	case *errorsa.AppError:
		httpCode = v.HTTPCode
		code1 = v.Code
		message = v.Message
	case string:
		code1 = httpCode
		message = v
	default:
		code1 = http.StatusInternalServerError
		message = "未知错误"
	}
	c.JSON(httpCode, Response{
		Code: code1, Msg: message, Data: nil, Request_ID: getRequestID(c),
	})

}

func getRequestID(c *gin.Context) string {
	if id, exists := c.Get("request_id"); exists {
		if s, ok := id.(string); ok {
			return s
		}
	}
	return ""
}
