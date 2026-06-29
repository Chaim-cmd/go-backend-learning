# Day3 接口命令测试

## 1.测试创建用户
curl.exe -s -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"张三","email":"zhangsan@example.com","age":25}'
## 2.查询用户
curl.exe -s http://localhost:8080/users/1 
## 3.更新用户
curl.exe -s -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"张三丰","email":"zhangsanfeng@example.com","age":100}'
## 4.删除用户
curl.exe -s -X DELETE http://localhost:8080/users/1
## 5.查询已删除的用户（应返回 404）
curl -s http://localhost:8080/users/1 
#期望: {"code":404,"msg":"用户不存在","data":null}
## 6.查询不存在的 ID
curl -s http://localhost:8080/users/999

