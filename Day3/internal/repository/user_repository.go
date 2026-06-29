package repository

import (
	"errors"
	"sync"

	"github.com/Chaim-cmd/go-backend-learning.git/Day3/internal/model"
)

// ErrUserNotFound 哨兵错误，service 层可以用error.Is 判断
var ErrUserNotFound = errors.New("用户不存在")

// UserRepository 接口 ：定义数据访问能力
type UserRepository interface {
	Create(user *model.User) error
	FindByID(id int) (*model.User, error)
	Update(user *model.User) error
	Delete(id int) error
}

// memoryUserRepository 是UserRepository的内存实现
type memoryUserRepository struct {
	mu     sync.RWMutex //保护并发读写
	data   map[int]*model.User
	NextID int
}

// NewMemoryUserRepository 构造函数，返回接口类型
func NewMemoryUserRepository() UserRepository {
	return &memoryUserRepository{
		data:   make(map[int]*model.User),
		NextID: 1,
	}

}

func (r *memoryUserRepository) Create(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.NextID
	r.NextID++
	r.data[user.ID] = user
	return nil
}
func (r *memoryUserRepository) FindByID(id int) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	user, ok := r.data[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}
func (r *memoryUserRepository) Update(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.data[user.ID]; !ok {
		return ErrUserNotFound
	}
	r.data[user.ID] = user
	return nil
}
func (r *memoryUserRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; !ok {
		return ErrUserNotFound
	}
	delete(r.data, id)
	return nil

}
