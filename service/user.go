package service

import (
	"fmt"
	"lygf/backend/dao/mysql"
	"lygf/backend/dao/redis"
	"lygf/backend/model/entity"
	"lygf/backend/model/param"
	"lygf/backend/pkg"

	"go.uber.org/zap"
)

// 用户注册
func UserRegister(p *param.ParamRegister) bool{
    // 先查mysql数据库是否存在用户名
    //user,err := mysql.GetUserByUsername(p.Username)
	user := mysql.GetUserByUsername(p.Username)
    // 从数据库查到的都不是参数，是与数据库对应的结构体
    //if err != nil && user == nil{
	if user == nil{
        // 创建新用户实例
        user = &entity.User{
            Username: p.Username,
            Password: p.Password,
        }
        // 插入新记录
        mysql.InsertUser(user)
        return true
    }
    // 用户名重复
    return false
}

// 用户登录
func UserLogin(p *param.ParamLogin) (token string,ok bool) {
    // 从数据库中查user比对 用户名和密码
    // 比对验证码（代码变量中/redis中）
    fmt.Println(p.Username)
    user := mysql.GetUserByUsername(p.Username)
    fmt.Println(user)
    if user == nil{
        // 用户不存在
        fmt.Println(1)
        return "",false
    }
    if user.Password != p.Password{
        // 密码错误
        fmt.Println(2)
        return "",false
    }
    // 验证码功能

    // 全部正确,生成token返回
    token,err := pkg.GenToken(user.ID,user.Username)
    if err != nil{
        // 系统繁忙，token生成错误。但是不中断登录过程，记录token生成错误日志
        token = ""
        zap.L().Error("token生成失败，请尽快修复！")
    }
    // 生成成功,存储到redis中
    redis.SetUserToken(token,user.ID)
    
    return token,true
}