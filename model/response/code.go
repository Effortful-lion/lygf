package response

// 响应码 + 响应提示信息

type ResCode int64

// 定义状态码
const(
	CodeSuccess ResCode = 1000 + iota	// 响应成功码
	CodeInvalidParam					// 参数校验失败码
	CodeError 							// 请求错误（服务器错误）
	CodeRegisterFailed					// 注册失败
	CodeLoginFailed						// 用户名或密码错误
	CodeLoginFailedByCode				// 验证码错误
	CodeNeedLogin						// 需要登录
	CodeSendEmailFailed					// 发送邮件失败
)

// 定义状态码和响应提示信息映射
var codeMsgMap = map[ResCode]string{
	CodeSuccess: "响应成功",
	CodeInvalidParam: "请求参数无效",
	CodeError:"服务错误",
	CodeRegisterFailed:"注册失败",
	CodeLoginFailed:"用户名或密码错误",
	CodeLoginFailedByCode:"验证码错误",
	CodeNeedLogin:"需要登录",
	CodeSendEmailFailed: "发送邮件失败",
}

// 根据状态码获得相应提示信息
func (code ResCode) Msg() string{
	msg, ok := codeMsgMap[code]
	if !ok {
		// 无法获取，服务器服务繁忙（内部有问题，查看日志问题）
		msg = codeMsgMap[CodeError]
	}
	return msg
}

