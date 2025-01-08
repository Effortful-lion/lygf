package param



// 注册参数
type ParamRegister struct {
    Username string `json:"username" binding:"required,phone"` // 用户名，必须是手机号
    Password string `json:"password" binding:"required,password"` // 密码，必须符合密码规则
}