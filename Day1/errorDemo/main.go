package main

import (
	"errors"
	"fmt"
)

// 自定义错误类型
type NotFoundError struct {
	Resource string
	ID       int
}

func (n *NotFoundError) Error() string {
	return fmt.Sprintf("Resource %s with ID %d not found", n.Resource, n.ID)
}

// 模拟数据库查询
func findUser(id int) (string, error) {
	db := map[int]string{2: "李四"}
	user, ok := db[id]
	if !ok {
		return "", fmt.Errorf("查询失败：%w", &NotFoundError{
			Resource: "User",
			ID:       id,
		})
	}

	return user, nil
}

func main() {
	//正常查询
	name, err := findUser(1)
	if err != nil {
		fmt.Println("错误", err)
	} else {
		fmt.Println("找到用户：", name)
	}

	//查找不到的情况
	_, err1 := findUser(99)
	if err != nil {
		fmt.Println("错误", err1)
	}

	//errors.As 提取具体的错误类型，获取结构体字段
	var notFound *NotFoundError
	if errors.As(err, &notFound) {
		fmt.Printf("具体原因 资源：%s,ID:%d\n", notFound.Resource, notFound.ID)
	}

	//errors.Is 判断是否是某个特定的错误
	var ErrPermission = errors.New("无权限")
	wrapped := fmt.Errorf("操作失败：%w", ErrPermission)
	fmt.Println("是权限错误吗：", errors.Is(wrapped, ErrPermission))

}
