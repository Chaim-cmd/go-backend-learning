package service

import (
	"errors"
	"fmt"

	"github.com/Chaim-cmd/go-backend-learning.git/Day3/internal/model"
	"github.com/Chaim-cmd/go-backend-learning.git/Day3/internal/repository"
)

// ErrEmailExists 业务错误：邮箱重复
var ErrEmailExists = errors.New("邮箱已被注册")

// UserService 接口：业务逻辑层
type UserService interface {
	CreateUser(name, email string, age int) (*model.User, error)
	GetUser(id int) (*model.User, error)
	UpdateUser(id int, name, email string, age int) (*model.User, error)
	DeleteUser(id int) error
}

type userService struct {
	repo repository.UserRepository //依赖倒置
}

// NewUserService 通过构造函数注入 repository(依赖注入)
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(name, email string, age int) (*model.User, error) {
	user := &model.User{
		Name:  name,
		Email: email,
		Age:   age,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, fmt.Errorf("创建用户失败：%w", err)
	}
	return user, nil
}
func (s *userService) GetUser(id int) (*model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, err
}
func (s *userService) UpdateUser(id int, name, email string, age int) (*model.User, error) {
	//先查询，确认存在
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	user.Name = name
	user.Email = email
	user.Age = age

	if err := s.repo.Update(user); err != nil {
		return nil, fmt.Errorf("更新用户失败：%w", err)
	}
	return user, nil
}
func (s *userService) DeleteUser(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
