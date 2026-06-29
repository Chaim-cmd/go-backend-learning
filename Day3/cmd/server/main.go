package main

import (
	"github.com/Chaim-cmd/go-backend-learning.git/Day3/internal/handler"
	"github.com/Chaim-cmd/go-backend-learning.git/Day3/internal/repository"
	"github.com/Chaim-cmd/go-backend-learning.git/Day3/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	//依赖注入链：repository->service->handler
	//这是控制反转的核心，上层不自己创建下层，而是被传入
	userRepo := repository.NewMemoryUserRepository()
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	r := gin.Default()
	users := r.Group("/users")
	{
		users.POST("", userHandler.Create)
		users.GET("/:id", userHandler.Get)
		users.PUT("/:id", userHandler.Update)
		users.DELETE("/:id", userHandler.Delete)
	}
	r.Run(":8080")
}
