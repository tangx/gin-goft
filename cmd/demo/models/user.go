package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId   int    `gorm:"index"`
	UserName string `gorm:"index"`
}
