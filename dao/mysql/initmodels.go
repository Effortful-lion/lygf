package mysql

import "lygf/backend/model/entity"

// 这里是关于 mysql 中的各种记录模型的定义：比如用户表、文章表、评论表等
// 一般来说，每个模型对应数据库中的一张表，每个字段对应表中的一个字段
// 初始化一次就可以了，不需要每次都初始化

// 初始化各种模型（）
func InitModels(){
	// 迁移数据库
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Product{})
	db.AutoMigrate(&entity.Shop{})
}