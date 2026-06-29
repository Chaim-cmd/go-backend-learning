package service

import (
	"errors"
	"strings"

	"github.com/Chaim-cmd/go-backend-learning.git/Day5/internal/model"
	"github.com/Chaim-cmd/go-backend-learning.git/Day5/internal/repository"
	errorsa "github.com/Chaim-cmd/go-backend-learning.git/Day5/pkg/errorsa"
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
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, errorsa.Conflict("邮箱已经被注册", err)
		}
		return nil, errorsa.Internal("创建用户失败", err)
	}
	return user, nil
}
func (s *userService) GetUser(id uint) (*model.User, error) {
	user, err := s.repo.FindByID(id)

	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, errorsa.NotFound("用户不存在", err)
		}
		return nil, errorsa.Internal("查询失败", err)
	}
	return user, nil

}
func (s *userService) UpdateUser(id uint, name, email string, age int) (*model.User, error) {
	user := &model.User{Name: name, Email: email, Age: age}
	user.ID = id

	if err := s.repo.Update(user); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, errorsa.NotFound("用户不存在", err)
		}
		return nil, errorsa.Internal("更新失败", err)

	}
	return s.repo.FindByID(id)
}
func (s *userService) DeleteUser(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return errorsa.NotFound("用户不存在", err)
		}
		return errorsa.Internal("删除失败", err)
	}
	return nil

}
