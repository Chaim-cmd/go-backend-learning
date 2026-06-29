package repository

import (
	"errors"

	"github.com/Chaim-cmd/go-backend-learning.git/Day5/internal/model"
	"gorm.io/gorm"
)

// ErrUserNotFound 哨兵错误，service 层可以用error.Is 判断
var ErrUserNotFound = errors.New("用户不存在")

// UserRepository 接口 ：定义数据访问能力
type UserRepository interface {
	Create(user *model.User) error
	FindByID(id uint) (*model.User, error)
	Update(user *model.User) error
	Delete(id uint) error
}

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewMysqlUserRepository(db *gorm.DB) UserRepository {
	return &mysqlUserRepository{db: db}
}

func (r *mysqlUserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error

}
func (r *mysqlUserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *mysqlUserRepository) Update(user *model.User) error {
	result := r.db.Model(&model.User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
		"age":   user.Age,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrUserNotFound
	}
	return nil
}
func (r *mysqlUserRepository) Delete(id uint) error {
	result := r.db.Delete(&model.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	//查看操作的影响行数
	if result.RowsAffected == 0 {
		return ErrUserNotFound
	}
	return nil
}
