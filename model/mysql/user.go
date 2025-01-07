package model

import "gorm.io/gorm"

// User 用户表
type User struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primary_key"` // 用户ID
	Username string `json:"username"`              // 用户名
	Password string `json:"password"`			   // 密码
}