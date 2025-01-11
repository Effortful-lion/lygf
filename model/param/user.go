package param

import (
	"lygf/backend/model/common"
)

// 注册参数
type ParamRegister struct {
	Email    string   `json:"email" binding:"required,email"`
	Password string   `json:"password" binding:"required,password"`
	Type     *common.UserType `json:"type" binding:"required"`
	Code     string   `json:"code" binding:"required,code"`
}

// 登录参数
type ParamLogin struct {
	Email string `json:"email" binding:"required,email"`    // 用户名，必须是手机号
	Password string `json:"password" binding:"required,password"` // 密码，必须符合密码规则
}

// 展示的用户信息参数
type ParamUserInfo struct {
    Picture string `json:"user_picture"`
    Username string `json:"username"`
    Introduction string `json:"user_introduction"`
    Background string `json:"background"`
	Products []Product `json:"products"`
}


// 个人信息编辑参数
type ParamUpdateUserInfo struct {
    Username string `json:"username"`
    Introduction string `json:"user_introduction"`
    Picture string `json:"user_picture"`
    Background string `json:"background"`
}

// // 用户主页参数
// type ParamUserHome struct {
//     UserID int `json:"user_id"`
//     Introduction string `json:"user_introduction"`
//     Background string `json:"background"`
//     Picture string `json:"user_picture"`
//     // TODO 
// }
