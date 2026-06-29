package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Chaim-cmd/go-backend-learning.git/Day5/internal/repository"
	"github.com/Chaim-cmd/go-backend-learning.git/Day5/internal/service"
	errorsa "github.com/Chaim-cmd/go-backend-learning.git/Day5/pkg/errorsa"
	"github.com/Chaim-cmd/go-backend-learning.git/Day5/pkg/response"
	"github.com/gin-gonic/gin"
)

// userHandler 只依赖service接口
type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

type UserRequest struct {
	Name  string `json:"name" binding:"required,min=2,max=60"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"required,gte=1,lte=120"`
}

func handlerErr(c *gin.Context, err error) {
	var errorsa *errorsa.AppError
	if errors.As(err, &errorsa) {
		response.Fail(c, 0, errorsa)
		return
	}
	response.Fail(c, http.StatusInternalServerError, "服务器内部错误")
}

func pareseID(c *gin.Context) (uint, bool) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, "id 必须为正整数")
		return 0, false
	}
	return uint(id), true
}

// POST /users
func (h *UserHandler) Create(c *gin.Context) {
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, "参数错误："+err.Error())
		return
	}
	user, err := h.svc.CreateUser(req.Name, req.Email, req.Age)
	if err != nil {
		handlerErr(c, err)
		return
	}
	response.Success(c, user)

}

// Get /users/:id
func (h *UserHandler) Get(c *gin.Context) {
	id, ok := pareseID(c)
	if !ok {
		return
	}

	user, err := h.svc.GetUser(uint(id))
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			handlerErr(c, err)
			return

		}
	}
	response.Success(c, user)
}

// PUT /users/:id
func (h *UserHandler) Update(c *gin.Context) {
	id, ok := pareseID(c)
	if !ok {
		return
	}
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, "参数错误:"+err.Error())
		return
	}
	user, err := h.svc.UpdateUser(uint(id), req.Name, req.Email, req.Age)
	if err != nil {
		handlerErr(c, err)
		return
	}
	response.Success(c, user)
}

// DELETE /users/:id
func (h *UserHandler) Delete(c *gin.Context) {
	id, ok := pareseID(c)
	if !ok {
		return
	}
	if err := h.svc.DeleteUser(id); err != nil {
		handlerErr(c, err)
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}
