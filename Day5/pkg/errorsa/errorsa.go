package errorsa

import (
	"fmt"
	"net/http"
)

// AppError 业务错误，携带 http状态码 + 对外展示的消息
// 和原生 go 兼容 ，可以被errors.As提取
type AppError struct {
	HTTPCode int    //返回给客户端的 HTTP 状态码
	Code     int    //业务错误码
	Message  string //对外展示的消息
	Err      error  //原始错误
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%d] %s:%v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// Unwrap 支持errors.Is / errors.As解包
func (e *AppError) Unwrap() error { return e.Err }

// 业务错误码定义
const (
	CodeBadRequest   = 400001
	CodeUnauthorized = 401001
	CodeNotFound     = 404001
	CodeConflict     = 409001
	CodeInternal     = 500001
)

// 快捷构造函数
func BadRequest(msg string, err error) *AppError {
	return &AppError{
		HTTPCode: http.StatusBadRequest,
		Code:     CodeBadRequest,
		Message:  msg,
		Err:      err,
	}

}

func NotFound(msg string, err error) *AppError {
	return &AppError{
		HTTPCode: http.StatusNotFound,
		Code:     CodeNotFound,
		Message:  msg,
		Err:      err,
	}

}

func Conflict(msg string, err error) *AppError {
	return &AppError{
		HTTPCode: http.StatusConflict,
		Code:     CodeConflict,
		Message:  msg,
		Err:      err,
	}

}
func Internal(msg string, err error) *AppError {
	return &AppError{
		HTTPCode: http.StatusInternalServerError,
		Code:     CodeInternal,
		Message:  msg,
		Err:      err,
	}

}
