package service

import (
	"lygf/backend/dao/mysql"
	"lygf/backend/model/entity"
	"lygf/backend/model/param"
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