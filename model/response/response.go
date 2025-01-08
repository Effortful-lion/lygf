package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 响应结构体
type ResponseData struct {
	Code ResCode    `json:"code"` // 自定义的code
	Msg  any   `json:"msg"`  // 自定义的msg
	Data any 	`json:"data"`// 自定义的数据  ,omitempty可忽略空值不展示  Data any 	`json:"data,omitempty"`
}

// 响应错误信息：code+错误信息
func ResponseError(c *gin.Context, code ResCode){
	c.JSON(http.StatusOK,&ResponseData{
		Code: code,
		Msg:code.Msg(),
		Data:nil,
	})
}

// 响应成功信息：
func ResponseSuccess(c *gin.Context, data any){
	c.JSON(http.StatusOK,&ResponseData{
		Code: CodeSuccess,
		Msg: CodeSuccess.Msg(),
		Data: data,
	})
}

// 响应具体错误信息
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg string){
	c.JSON(http.StatusOK,&ResponseData{
		Code: code,
		Msg: msg,
		Data: nil,
	})
}