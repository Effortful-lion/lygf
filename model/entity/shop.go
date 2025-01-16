package entity

import (
	"lygf/backend/model/common"

	"gorm.io/gorm"
)


// 商铺结构体
type Shop struct {
	ID          int    `gorm:"primary_key;auto_increment" json:"id"`
	UserID      int    `gorm:"column:user_id" json:"user_id"`	// 商家的用户ID
	Name        string `gorm:"column:name" json:"name"`		// 商家名称
	Category    common.Category `gorm:"column:category" json:"category"`	// 商家分类
	HOT			bool `gorm:"column:hot" json:"hot"`			// 是否为热门(0 否 1 是)
	Location    string `gorm:"column:location" json:"location"`		// 商家位置
	Introduction string `gorm:"column:introduction" json:"introduction"`	// 商家介绍
	Score       float64 `gorm:"column:score" json:"score"`					// 商家评分
	Phone       string `gorm:"column:phone" json:"phone"`
	*gorm.Model
}
