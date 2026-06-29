package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Chaim-cmd/go-backend-learning.git/Day5/internal/model"
	"github.com/Chaim-cmd/go-backend-learning.git/Day5/internal/repository"
	apperrors "github.com/Chaim-cmd/go-backend-learning.git/Day5/pkg/errorsa"
)

// MockUserRepository 手动实现 mock，不依赖 mockgen 工具
// 实际项目可以用 go install go.uber.org/mock/mockgen 自动生成
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByID(id uint) (*model.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) Update(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// ---- 测试用例 ----

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo)

	// 设定：调用 Create 时返回 nil（成功）
	mockRepo.On("Create", mock.AnythingOfType("*model.User")).Return(nil)

	user, err := svc.CreateUser("张三", "zhangsan@example.com", 25)

	assert.NoError(t, err)
	assert.Equal(t, "张三", user.Name)
	assert.Equal(t, "zhangsan@example.com", user.Email)
	mockRepo.AssertExpectations(t) // 验证 mock 被正确调用
}

func TestCreateUser_DuplicateEmail(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo)

	// 模拟 MySQL 唯一索引冲突
	mockRepo.On("Create", mock.AnythingOfType("*model.User")).
		Return(errors.New("Duplicate entry 'zhangsan@example.com' for key 'email'"))

	_, err := svc.CreateUser("张三", "zhangsan@example.com", 25)

	assert.Error(t, err)
	var appErr *apperrors.AppError
	assert.True(t, errors.As(err, &appErr))
	assert.Equal(t, apperrors.CodeConflict, appErr.Code)
	assert.Equal(t, "邮箱已被注册", appErr.Message)
}

func TestGetUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo)

	expected := &model.User{Name: "张三", Email: "zhangsan@example.com", Age: 25}
	expected.ID = 1
	mockRepo.On("FindByID", uint(1)).Return(expected, nil)

	user, err := svc.GetUser(1)

	assert.NoError(t, err)
	assert.Equal(t, "张三", user.Name)
}

func TestGetUser_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo)

	mockRepo.On("FindByID", uint(999)).Return(nil, repository.ErrUserNotFound)

	_, err := svc.GetUser(999)

	assert.Error(t, err)
	var appErr *apperrors.AppError
	assert.True(t, errors.As(err, &appErr))
	assert.Equal(t, apperrors.CodeNotFound, appErr.Code)
}

func TestDeleteUser_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo)

	mockRepo.On("Delete", uint(999)).Return(repository.ErrUserNotFound)

	err := svc.DeleteUser(999)

	assert.Error(t, err)
	var appErr *apperrors.AppError
	assert.True(t, errors.As(err, &appErr))
	assert.Equal(t, apperrors.CodeNotFound, appErr.Code)
}
