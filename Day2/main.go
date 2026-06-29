package main

import (
	"net/http"

	"github.com/Chaim-cmd/go-backend-learning.git/Day2/internal/middleware"
	"github.com/Chaim-cmd/go-backend-learning.git/Day2/pkg/response"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required,min=2,max=20"`
	Email string `json:"Email" binding:"required,email"`
	Age   int    `json:"age" binding:"required,gte=1,lte=12"`
}

// GET /ping  返回pong
func handlePing(c *gin.Context) {
	response.Success(c, gin.H{"message": "pong"})
}

// GET /users/:id 返回路径参数
func handleGetUser(c *gin.Context) {
	id := c.Param("id")
	name := c.Query("name")

	response.Success(c, gin.H{
		"id":   id,
		"name": name,
		"tip":  "这是路径演示",
	})

}

// POST /users 读JSON body
func handleCreateUser(c *gin.Context) {
	var req CreateUserRequest

	//shouldBindJSON 绑定数据
	if err := c.ShouldBindJSON(req); err != nil {
		response.Fail(c, http.StatusBadRequest, "参数错误："+err.Error())
		return
	}

	response.Success(c, gin.H{
		"id":    1001,
		"name":  req.Name,
		"email": req.Email,
		"age":   req.Age,
	})

}

func main() {
	//使用gin.New() 而不是gin.Default(),手动挂中间件
	r := gin.New()

	//挂载中间件
	r.Use(middleware.RequestID()) //先请求request_id
	r.Use(middleware.Logger())    //在打印日志

	r.Use(gin.Recovery()) //panic放到最后

	//路由注册
	r.GET("/ping", handlePing)
	user := r.Group("/users")
	{
		user.GET("/:id", handleGetUser)
		user.POST("", handleCreateUser)
	}
	r.Run(":8080")

}
