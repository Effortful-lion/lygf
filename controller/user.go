package controller

import (
	"fmt"
	"lygf/backend/dao/mysql"
	"lygf/backend/model/param"
	response "lygf/backend/model/response"
	"lygf/backend/pkg"
	"lygf/backend/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 用户注册
func UserRegister(ctx *gin.Context) {
    // 1/2. 获取参数和参数校验
    user := &param.ParamRegister{}
    if err := ctx.ShouldBindJSON(user); err != nil {
        errs, ok := err.(validator.ValidationErrors)
        if !ok {
            // 非validator规则
            response.ResponseError(ctx, response.CodeInvalidParam)
            return
        }
        // validator 规则
        errsStr := pkg.RemoveStructName(errs.Translate(pkg.Trans))
        response.ResponseErrorWithMsg(ctx, response.CodeInvalidParam, errsStr)
        return
    }

    if ok := service.UserRegister(user); !ok{
		// 保存数据库失败
		response.ResponseError(ctx,response.CodeRegisterFailed)
		return
	}
	// 注册成功
	response.ResponseSuccess(ctx,nil)
}

// 用户登录
func UserLogin(ctx *gin.Context) {
    // 1/2. 获取参数/参数校验
    p := &param.ParamLogin{}
    if err := ctx.ShouldBindJSON(p); err != nil{
        errs, ok := err.(validator.ValidationErrors)
        if !ok{
            response.ResponseError(ctx,response.CodeInvalidParam)
            return
        } 
        errStr := pkg.RemoveStructName(errs.Translate(pkg.Trans))
        response.ResponseErrorWithMsg(ctx, response.CodeInvalidParam, errStr)
    } 
    // 3. 业务处理
    if token,ok := service.UserLogin(p); !ok{
        // 用户名或者密码有误
        fmt.Println(ok)
        response.ResponseError(ctx,response.CodeLoginFailed)
        return
    }else{
        // 4. 返回响应 : 登录成功后可以返回：用户信息、过期时间、refresh_token、权限信息（这样使得前端的请求减少、有利于后面的各种操作）
        user := mysql.GetUserByUsername(p.Username)     // 这里获得 user 是为了获取userID
        response.ResponseSuccess(ctx,gin.H{
          "user_id" : user.ID,          // TODO存在 安全风险 ；如果前端可以实现token解析，我就撤掉
          "username": user.Username,
          "token": token,
        })
    }
    
}