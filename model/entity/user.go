package entity

import (
    "gorm.io/gorm"
    _ "lygf/backend/pkg" // 导入自定义验证器包
)

// User 用户表
type User struct {
    gorm.Model
    ID       int    `json:"id" gorm:"primary_key"` // 用户ID
    Username string `json:"username" binding:"required,phone"` // 用户名，必须是手机号
    Password string `json:"password" binding:"required,password"` // 密码，必须符合密码规则
}