package entity

import (
	"lygf/backend/model/common"
	"gorm.io/gorm"
)

// User 用户表
type User struct {
    ID       int    `json:"id" gorm:"primary_key"` // 用户ID（唯一）
    Email    string `json:"email" binding:"required,email" gorm:"unique"` // 邮箱（唯一）
    Username string `json:"username"`  // 用户名(不唯一)
    Password string `json:"password" binding:"required,password"` // 密码（不唯一） 
    Type     common.UserType `json:"type" binding:"required"`
    Introduction string `json:"user_introduction"`
    Background string `json:"background"`
    Picture string `json:"user_picture"`
    gorm.Model
}