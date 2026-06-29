package main

import (
	"log"

	"github.com/Chaim-cmd/go-backend-learning.git/Day4/internal/config"
	"github.com/Chaim-cmd/go-backend-learning.git/Day4/internal/handler"
	"github.com/Chaim-cmd/go-backend-learning.git/Day4/internal/repository"
	"github.com/Chaim-cmd/go-backend-learning.git/Day4/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	//加载配置
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("加载配置失败:%v", err)
	}

	//连接数据库
	db, err := repository.NewDB(&cfg.Database)
	if err != nil {
		log.Fatalf("数据库初始化失败:%v", err)
	}

	//依赖注入链：repository->service->handler
	//这是控制反转的核心，上层不自己创建下层，而是被传入
	userRepo := repository.NewMysqlUserRepository(db)
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
	r.Run(":" + cfg.Server.Port)
}
