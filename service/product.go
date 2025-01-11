package service

import (
	"lygf/backend/dao/mysql"
	"lygf/backend/model/param"
)

// GetProductList 获取商品列表
func GetProductList() (products []param.Product) {
	return mysql.GetProductList()
}