package controller

import (
	"fmt"
	"lygf/backend/dao/mysql"
	"lygf/backend/model/param"
	response "lygf/backend/model/response"
	"lygf/backend/pkg"
	"lygf/backend/service"
	"time"

	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 用户退出
func UserLogout(ctx *gin.Context) {
    user_id, err := getCurrentUserId(ctx)
    if err != nil {
        response.ResponseError(ctx, response.CodeNeedLogin)
        return
    }
    if err := service.UserLogout(user_id); err != nil {
        response.ResponseError(ctx, response.CodeError)
        return
    }
    response.ResponseSuccess(ctx, nil)
}

// 注销用户
func UserDelete(ctx *gin.Context) {
    user_id,err := getCurrentUserId(ctx)
    if err != nil {
        response.ResponseError(ctx, response.CodeNeedLogin)
    }
    if err := service.DeleteUser(user_id); err != nil {
        response.ResponseError(ctx, response.CodeError)
    }
    response.ResponseSuccess(ctx, nil)
}

// 为编辑用户信息回显，查询用户信息
func GetUserInfoForEdit(ctx *gin.Context){
    user_id,err := getCurrentUserId(ctx)
    if err != nil {
        response.ResponseError(ctx, response.CodeNeedLogin)
        return
    }
    user,err := service.GetUserInfoForEdit(user_id)
    if err != nil {
        response.ResponseError(ctx, response.CodeError)
        return
    }
    response.ResponseSuccess(ctx, user)
}

// 更新用户信息
func UpdateUserInfo(ctx *gin.Context) {
    var requestdata map[string]interface{}
    if err := ctx.ShouldBindJSON(&requestdata); err != nil {
        // 将参数绑定到map变量中
        response.ResponseError(ctx, response.CodeInvalidParam)
        return
    }

    // 将map变量中的data数据提取出来
    data, ok := requestdata["data"].(map[string]interface{})
    if !ok {
        response.ResponseError(ctx, response.CodeInvalidParam)
        return
    }

    user := &param.ParamUpdateUserInfo{}
    // 进行安全的类型断言和转换
    if username, ok := data["username"].(string); ok {
        user.Username = username
    } else {
        response.ResponseError(ctx, response.CodeInvalidParam)
        return
    }

    if introduction, ok := data["user_introduction"].(string); ok {
        user.Introduction = introduction
    } else {
        response.ResponseError(ctx, response.CodeInvalidParam)
        return
    }

    if picture, ok := data["user_picture"].(string); ok {
        user.Picture = picture
    } else {
        response.ResponseError(ctx, response.CodeInvalidParam)
        return
    }

    if background, ok := data["background"].(string); ok {
        user.Background = background
    } else {
        response.ResponseError(ctx, response.CodeInvalidParam)
        return
    }

    userID,err := getCurrentUserId(ctx)
    if err != nil {
        response.ResponseError(ctx, response.CodeNeedLogin)
        return
    }

    // 调用服务层更新用户信息
    if err := service.UpdateUserInfo(userID,user); err != nil {
        response.ResponseError(ctx, response.CodeError)
        return
    }

    response.ResponseSuccess(ctx, nil)
}
// 获取用户信息
func GetUserInfo(ctx *gin.Context){
    //idStr := ctx.Param("id")
    //userID,err := strconv.Atoi(idStr)
    userID,err := getCurrentUserId(ctx)
    fmt.Println(userID)
    if err != nil {
        fmt.Println(userID)
        response.ResponseError(ctx, response.CodeError)
    }
    user,err := service.GetUserInfo(userID)
    if user == nil {
        response.ResponseErrorWithMsg(ctx, response.CodeError,"用户不存在")
        return
    }
    // //TODO 获取用户信息（还需要一些条件，比如：社区信息？地理信息？从而获取附近商品的信息）
    products := service.GetProductList()
    var userinfo param.ParamUserInfo
    userinfo.Username = user.Username
    userinfo.Picture = user.Picture
    userinfo.Introduction = user.Introduction
    userinfo.Background = user.Background
    userinfo.Products = products
    if err != nil {
        fmt.Println(2)
        response.ResponseError(ctx, response.CodeError)
        return
    }
    response.ResponseSuccess(ctx, userinfo)
}

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

    if err := service.UserRegister(user); err != nil {
        // 保存数据库失败
        response.ResponseErrorWithMsg(ctx, response.CodeRegisterFailed,err.Error())
        return
    }
    // 注册成功
    response.ResponseSuccess(ctx, nil)
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
    if token,err := service.UserLogin(p); err != nil{
        // 用户名或者密码有误
        response.ResponseError(ctx,response.CodeLoginFailed)
        return
    }else{
        // 4. 返回响应 : 登录成功后可以返回：用户信息、过期时间、refresh_token、权限信息（这样使得前端的请求减少、有利于后面的各种操作）
        user := mysql.GetUserByEmail(p.Email)     // 这里获得 user 是为了获取userID
        response.ResponseSuccess(ctx,gin.H{
          "user_id" : user.ID,          // TODO存在 安全风险 ；如果前端可以实现token解析，我就撤掉
          "token": token,
        })
    }
}

// 获得后端时间
func GetTime(c *gin.Context) {
    time_stamp := time.Now().Unix()
    response.ResponseSuccess(c,gin.H{
        "time_stamp": time_stamp,
    })
}