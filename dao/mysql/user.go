package mysql

import (
	"lygf/backend/model/entity"
	// "go.uber.org/zap"
	// "gorm.io/gorm"
)

// 关于用户的数据库操作

// handleDBError 处理数据库用户操作后的错误
// func handleDBError(result *gorm.DB) error {
//     if result.Error != nil {
//         zap.L().Error("数据库用户模块操作失败", zap.Error(result.Error))
//         return result.Error
//     }
//     return nil
// }

// func GetUserByUsername(username string) (user *entity.User) {
// 	//handleDBError(db.Where("username = ?", username).Take(user))
// 	user = new(entity.User)		// 函数上面是返回值的声明，并没有初始化
// 	db.Where("username = ?", username).First(user)
// 	return user
// }

// 删除用户信息
func DeleteUser(userID int) (err error) {
    if err = db.Delete(&entity.User{}, userID).Error; err != nil {
        return err
    }
    return nil
}

// 根据邮箱查询用户
func GetUserByEmail(email string) (user *entity.User) {
	user = new(entity.User)
	result := db.Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil
	}
	return user
}

// 插入用户
func InsertUser(user *entity.User) error{
	result := db.Create(user)
	//return handleDBError(result)
	return result.Error
}

// 根据ID获取用户
func GetUserByID(id int) (user *entity.User,err error){
	user = new(entity.User)
	result := db.Where("id = ?", id).First(user)
	if result.Error != nil {
		err = result.Error
		return nil,err
	}
	return user,err
}

// 更新用户编辑信息
func UpdateUserInfo(user *entity.User) error {
	result := db.Model(&entity.User{}).Where("id = ?",user.ID).Updates(user)
	// 如果你不指定 WHERE 条件，gorm 会尝试更新所有记录
	//result := db.Model(&entity.User{}).Updates(user)   // 错误
	return result.Error
}