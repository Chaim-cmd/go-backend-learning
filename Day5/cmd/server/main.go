package main

import (
	"log"

	"github.com/Chaim-cmd/go-backend-learning.git/Day5/internal/config"
	"github.com/Chaim-cmd/go-backend-learning.git/Day5/internal/handler"
	"github.com/Chaim-cmd/go-backend-learning.git/Day5/internal/middleware"
	"github.com/Chaim-cmd/go-backend-learning.git/Day5/internal/repository"
	"github.com/Chaim-cmd/go-backend-learning.git/Day5/internal/service"
	"github.com/Chaim-cmd/go-backend-learning.git/Day5/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	//日志初始化
	logger.Init("development")
	logger.Info("服务启动中")
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

	gin.SetMode(gin.DebugMode)
	r := gin.New()
	//中间件
	r.Use(middleware.Request_id())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	users := r.Group("/users")
	{
		users.POST("", userHandler.Create)
		users.GET("/:id", userHandler.Get)
		users.PUT("/:id", userHandler.Update)
		users.DELETE("/:id", userHandler.Delete)
	}
	logger.Info("服务启动成功", zap.String("port", cfg.Server.Port))
	r.Run(":" + cfg.Server.Port)
}
