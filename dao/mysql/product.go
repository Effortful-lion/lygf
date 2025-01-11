package mysql

import (
	"lygf/backend/model/param"
)

// 获取所有商品列表
func GetProductList() (products []param.Product) {
	db.Find(&products)
	return products
}