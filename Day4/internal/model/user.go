package model

import "gorm.io/gorm"

//用户模型
type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"type:varchar(50);not null"`
	Email string `json:"email" gorm:"type:varchar(100); uniqueIndex; not null"`
	Age   int    `json:"age" gorm:"type:int"`
}
