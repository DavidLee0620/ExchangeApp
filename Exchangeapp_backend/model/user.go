package model

import "gorm.io/gorm"

type User struct {
	//预定义
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}
