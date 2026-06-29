package repository

import (
	"fmt"

	"github.com/Chaim-cmd/go-backend-learning.git/Day5/internal/config"
	"github.com/Chaim-cmd/go-backend-learning.git/Day5/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewDB 建立数据库连接 + 自动建表
func NewDB(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //打印sql语句
	})
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败:%w", err)
	}

	//AutoMigrate:根据struct 自动建表/加字段 （生成环境通常用专门的迁移工具）
	if err := db.AutoMigrate(&model.User{}); err != nil {
		return nil, fmt.Errorf("自动迁移失败:%w", err)
	}
	return db, nil

}
