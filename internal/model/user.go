package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size: 50"`
	Sex      int    `gorm:"default: 1"`
}
