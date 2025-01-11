package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// 请求的预处理，取出context的东西，获取重要信息，为后面controller的处理请求准备调用


var ErrorUserNotLogin = errors.New("用户未登录")
var ContextUserIDKey = "userID"

// userid是通过解析token后，将token负载中的userid解析并设置到context中
// 获取当前登录的用户的ID
func getCurrentUserId(c *gin.Context) (userId int, err error) {
	uid, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return 
	}
	userId,ok = uid.(int)
	if !ok {
		err = ErrorUserNotLogin
		return 
	}
	return userId,nil
}