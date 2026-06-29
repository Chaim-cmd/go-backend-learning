# Day4 测试命令

## 启动Mysql
docker run -d --name mysql-learning \
  -e MYSQL_ROOT_PASSWORD=123456 \
  -e MYSQL_DATABASE=go_learning \
  -p 3306:3306 \
  mysql:8.0

# 等待约 10-20 秒让 MySQL 完全启动
docker logs -f mysql-learning
# 看到 "ready for connections" 字样后 Ctrl+C 退出日志查看

# 用 Docker 直接看数据库里的数据（可选，验证用）
docker exec -it mysql-learning mysql -uroot -p123456 go_learning \
  -e "SELECT * FROM users;"


# 5：完整 CRUD 测试
 ## 创建用户
curl -s -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"张三","email":"zhangsan@example.com","age":25}' | jq .
 ## 查询用户
 curl -s http://localhost:8080/users/1 | jq .
 ## 更新用户
 curl -s -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"张三丰","email":"zhangsanfeng@example.com","age":100}' | jq .

  ## 删除用户
  curl -s -X DELETE http://localhost:8080/users/1 | jq .