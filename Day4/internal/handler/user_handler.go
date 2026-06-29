package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Chaim-cmd/go-backend-learning.git/Day4/internal/repository"
	"github.com/Chaim-cmd/go-backend-learning.git/Day4/internal/service"
	"github.com/Chaim-cmd/go-backend-learning.git/Day4/pkg/response"
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

// POST /users
func (h *UserHandler) Create(c *gin.Context) {
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, "参数错误："+err.Error())
		return
	}
	user, err := h.svc.CreateUser(req.Name, req.Email, req.Age)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, user)

}

// Get /users/:id
func (h *UserHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, "id必须是数字")
		return
	}

	user, err := h.svc.GetUser(uint(id))
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			response.Fail(c, http.StatusNotFound, "用户不存在")
			return

		}
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, user)
}

// PUT /users/:id
func (h *UserHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, "id必须是数字")
		return
	}
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, "参数错误:"+err.Error())
		return
	}
	user, err := h.svc.UpdateUser(uint(id), req.Name, req.Email, req.Age)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			response.Fail(c, http.StatusNotFound, "用户不存在")
			return
		}
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, user)
}

// DELETE /users/:id
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, "id必须是数字")
		return
	}
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, "参数错误:"+err.Error())
		return
	}
	if err := h.svc.DeleteUser(uint(id)); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			response.Fail(c, http.StatusNotFound, "用户不存在")
			return
		}
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}
