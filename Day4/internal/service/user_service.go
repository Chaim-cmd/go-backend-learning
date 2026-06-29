package service

import (
	"fmt"

	"github.com/Chaim-cmd/go-backend-learning.git/Day4/internal/model"
	"github.com/Chaim-cmd/go-backend-learning.git/Day4/internal/repository"
)

// UserService 接口：业务逻辑层
type UserService interface {
	CreateUser(name, email string, age int) (*model.User, error)
	GetUser(id uint) (*model.User, error)
	UpdateUser(id uint, name, email string, age int) (*model.User, error)
	DeleteUser(id uint) error
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
func (s *userService) GetUser(id uint) (*model.User, error) {
	return s.repo.FindByID(id)

}
func (s *userService) UpdateUser(id uint, name, email string, age int) (*model.User, error) {
	user := &model.User{Name: name, Email: email, Age: age}
	user.ID = id

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}
	return s.repo.FindByID(id)
}
func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)

}
