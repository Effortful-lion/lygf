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

func GetUserByUsername(username string) (user *entity.User) {
	//handleDBError(db.Where("username = ?", username).Take(user))
	user = new(entity.User)		// 函数上面是返回值的声明，并没有初始化
	db.Where("username = ?", username).First(user)
	return user
}

func InsertUser(user *entity.User) error{
	result := db.Create(user)
	//return handleDBError(result)
	return result.Error
}