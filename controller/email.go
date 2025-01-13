package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"lygf/backend/model/response"
	"lygf/backend/pkg/email"
	"lygf/backend/service"
	"lygf/backend/setting"
	"github.com/gin-gonic/gin"
)

// 生成随机验证码
func generateVerificationCode() string {
    rand.Seed(time.Now().UnixNano())
    code := fmt.Sprintf("%06d", rand.Intn(1000000))
    fmt.Println("验证码生成成功：", code)
    return code // 6位数字验证码
}

// 发送验证码
func SendEmailCode(c *gin.Context) {
    type RequestBody struct {
        Email string `json:"email" binding:"required,email"`
    }
    var req RequestBody
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 生成验证码
    code := generateVerificationCode()

    // 保存验证码到 Redis（或者数据库）
    if err := service.SaveVerificationCode(req.Email, code, 5*time.Minute); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "验证码保存失败"})
        return
    }

    cfg := setting.Conf.EmailConfig

    fmt.Println(cfg.Host,cfg.Port,cfg.Username,cfg.Password,cfg.FromEmail)

    // 导入邮件发送模块发送邮件
    emailSender := email.NewEmailSender(
        cfg.Host,               // SMTP 服务器
        cfg.Port,                         // 端口号（QQ邮箱通常使用 SSL 加密）
        cfg.Username,         // 用户名(一般默认是邮箱)
        cfg.Password,          // 授权码（非邮箱密码）
        cfg.FromEmail,         // 发件邮箱
    )
    subject := "邻优果坊验证码"
    body := fmt.Sprintf(email.VerificationCodeFormat,code)
    err := emailSender.SendEmail(req.Email, subject,body)
    if err != nil {
        fmt.Println("验证码发送失败", err) // 添加错误日志
        c.JSON(http.StatusInternalServerError, gin.H{"error": "验证码发送失败"})
        return
    }

    response.ResponseSuccess(c, gin.H{"code": code})
}


