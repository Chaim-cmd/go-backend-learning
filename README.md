# Go Backend Learning

Go 高并发后端学习项目。

Show Image

技术栈


框架：Gin
数据库：MySQL + GORM
日志：Zap（结构化日志）
配置：Viper
测试：testify + mock


本地启动

bash# 启动 MySQL
docker run -d --name mysql-learning \
  -e MYSQL_ROOT_PASSWORD=123456 \
  -e MYSQL_DATABASE=go_learning \
  -p 3306:3306 mysql:8.0

# 启动服务
go mod tidy
go run cmd/server/main.go

运行测试

bash# 单元测试（不需要数据库）
go test -v -race ./internal/service/...

# 全部测试 + 覆盖率
go test -v -race -coverprofile=coverage.out ./...
go tool cover -func=coverage.out

API 文档

方法路径描述
POST/users创建用户
GET/users/:id查询用户
PUT/users/:id更新用户
DELETE/users/:id删除用户