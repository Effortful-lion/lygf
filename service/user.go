package service

import (
	"errors"
	"fmt"
	"lygf/backend/dao/mysql"
	"lygf/backend/dao/redis"
	"lygf/backend/model/common"
	"lygf/backend/model/entity"
	"lygf/backend/model/param"
	"lygf/backend/pkg"
	"time"

	"go.uber.org/zap"
)

// 更新用户信息
func UpdateUserInfo(userID int,userInfo *param.ParamUpdateUserInfo) (err error) {
    user := &entity.User{}
    user.ID = userID
    user.Username = userInfo.Username
    user.Picture = userInfo.Picture
    user.Background = userInfo.Background
    user.Introduction = userInfo.Introduction
    if err = mysql.UpdateUserInfo(user); err != nil {
        zap.L().Error("更新用户信息失败", zap.Error(err))
        return err
    }
    return nil
}

// 编辑用户信息回显
func GetUserInfoForEdit(userID int) (result *param.ParamUpdateUserInfo,err error) {
    user,err := GetUserInfo(userID)
    if err != nil {
        return nil,err
    }
    result = &param.ParamUpdateUserInfo{}
    result.Username = user.Username
    result.Picture = user.Picture
    result.Background = user.Background
    result.Introduction = user.Introduction
    return result,err
}

// 获取用户信息
func GetUserInfo(userID int) (user *entity.User,err error){
    user,err = mysql.GetUserByID(userID)
    return user,err
}

// 用户注册
func UserRegister(p *param.ParamRegister) error{
	user := mysql.GetUserByEmail(p.Email)
    currentTimestamp := time.Now().Unix()
    var err error
    // 判断验证码是否正确
    code := p.Code
    key := "email_code:" + p.Email
    result := redis.GetRedis().Get(key).Val()
    if code != result {
        err = errors.New("验证码错误")
        return err
    }

	if user == nil{
        // 创建新用户实例
        user = &entity.User{
            Email: p.Email,
            Username: fmt.Sprint("果友",currentTimestamp), // 默认用户名
            Password: p.Password,
            Type: common.UserType(*p.Type),
        }
        // 插入新记录
        mysql.InsertUser(user)
        return nil
    }
    // 邮箱存在
    err = errors.New("邮箱已存在")
    return err
}

// 用户登录
func UserLogin(p *param.ParamLogin) (token string,err error) {
    // 从数据库中查user比对 用户名和密码
    // 比对验证码（代码变量中/redis中）
    user := mysql.GetUserByEmail(p.Email)
    if user == nil{
        // 用户不存在
        err = errors.New("用户不存在")
        return "",err
    }
    if user.Password != p.Password{
        // 密码错误
        err = errors.New("密码错误")
        return "",err
    }
    // 全部正确,生成token返回
    token,err = pkg.GenToken(user.ID,user.Username)
    if err != nil{
        // 系统繁忙，token生成错误。但是不中断登录过程，记录token生成错误日志
        token = ""
        zap.L().Error("token生成失败，请尽快修复！")
    }
    // 生成成功,存储到redis中
    redis.SetUserToken(token,user.ID)
    
    return token,nil
}