package controller

import (
	"lygf/backend/model/common"
	"lygf/backend/model/response"
	"lygf/backend/service"
	"github.com/gin-gonic/gin"
)

// 首页处理方法

// 获得商家分类列表
func GetCategories(c *gin.Context)(){
	silce := []common.Category{common.FRESH,common.CATERING,common.SPOT}
	// 将所有枚举类型转换为字符串
	data := make([]string,len(silce))
	for i,v := range silce{
		data[i] = v.String()
	}
	response.ResponseSuccess(c,data)
}

// 根据分类获得所有商家列表
func GetShops(c *gin.Context)(){
	// 一级栏目类型id
	cateID1 := c.Params.ByName("cateID1")
	// 二级栏目类型id：商铺类型id
	cateID2 := c.Params.ByName("cateId2")
	data,err := service.GetShops(cateID1,cateID2) 
	if err != nil {
		response.ResponseErrorWithMsg(c,response.CodeError,err.Error())
	}
	response.ResponseSuccess(c,data)
}