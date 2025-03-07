package entity

// 商品信息

import "gorm.io/gorm"

type Product struct {
	ID           int    `json:"id" gorm:"primary_key"`	// 产品ID
	ShopID       int    `gorm:"column:shop_id" json:"shop_id"`	// 商户
	Name         string `json:"product_name"`			// 产品名称
	Price        int    `json:"product_price"`			// 价格
	Unit         string `json:"product_unit"`			// 价格单位
	Capacity     float64    `json:"product_capacity"`	// 容量
	CapacityUnit string `json:"product_capacity_unit"`	// 容量单位
	Picture      string `json:"product_picture"`		// 产品图片
	Introduction string `json:"product_introduction"`	// 产品介绍
	Score        float64`json:"product_score"`		// 产品评分
	gorm.Model
}