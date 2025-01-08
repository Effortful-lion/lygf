package controller

import (
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
    // 1. 获取参数

    // 2. 参数校验

    // 3. 业务处理

    // 4. 返回响应

}