package redis

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// 删除redis中token
func DelUserToken(ID int) error {
	return rdb.Del("token_user" + fmt.Sprint(ID)).Err()
}

// 将生成的token存储到redis中
func SetUserToken(token string, ID int) error {
	var TokenExpireDuration = viper.GetInt("auth.jwt_expire")
	return rdb.Set("token_user" + fmt.Sprint(ID),token,time.Hour * time.Duration(TokenExpireDuration)).Err()
}

// 通过用户ID得到用户token
func GetUserToken(ID int) (token string,err error){
	return rdb.Get("token_user" + fmt.Sprint(ID)).Result()
}