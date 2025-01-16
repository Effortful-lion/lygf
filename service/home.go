package service

import (
	"errors"
	"lygf/backend/dao/mysql"
	"lygf/backend/model/entity"
)

func GetShops(cateID1, cateID2 string) (*[]entity.Shop,error){
	data := &[]entity.Shop{}
	data = mysql.GetShops(cateID1, cateID2)
	if data == nil { 
		return nil,errors.New("没有找到相关数据")
	}
	return data,nil
}