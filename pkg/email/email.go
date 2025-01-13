package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

// 邮件发送器
type EmailSender struct {
    SMTPHost  string
    SMTPPort  int
    Username  string
    Password  string
    FromEmail string
}

// 初始化邮件发送器：邮件的 发送协议、发送端口、发送邮箱、发送邮箱的密码（一般服务商根据密码提供了授权码）、发送人
func NewEmailSender(smtpHost string, smtpPort int, username, password, fromEmail string) *EmailSender {
    return &EmailSender{
        SMTPHost:  smtpHost,
        SMTPPort:  smtpPort,
        Username:  username,
        Password:  password,
        FromEmail: fromEmail,
    }
}

// 发送邮件
func (es *EmailSender) SendEmail(toEmail, subject, body string) error {
    // 创建邮件消息
    m := gomail.NewMessage()
    // 设置邮件的基本信息
    m.SetHeader("From", es.FromEmail)
    m.SetHeader("To", toEmail)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body)

    d := gomail.NewDialer(es.SMTPHost, es.SMTPPort, es.Username, es.Password)
    err := d.DialAndSend(m)
    if err != nil {
        fmt.Println("Error sending email:", err) // 添加错误日志
    }
    return err
}

// 验证码 html 格式
const VerificationCodeFormat = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>验证码邮件</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f9;
            margin: 0;
            padding: 0;
        }
        .email-container {
            max-width: 600px;
            margin: 20px auto;
            background-color: #ffffff;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
            border-radius: 8px;
            overflow: hidden;
        }
        .email-header {
            background-color: #4CAF50;
            color: #ffffff;
            text-align: center;
            padding: 20px;
        }
        .email-header h1 {
            margin: 0;
            font-size: 24px;
        }
        .email-body {
            padding: 20px;
            color: #333333;
            line-height: 1.6;
        }
        .email-body p {
            margin: 10px 0;
        }
        .email-code {
            font-size: 24px;
            color: #4CAF50;
            font-weight: bold;
            text-align: center;
            margin: 20px 0;
            letter-spacing: 3px;
        }
        .email-footer {
            background-color: #f4f4f9;
            color: #777777;
            text-align: center;
            padding: 10px 20px;
            font-size: 12px;
        }
        .email-footer a {
            color: #4CAF50;
            text-decoration: none;
        }
        .email-footer a:hover {
            text-decoration: underline;
        }
    </style>
</head>
<body>
    <div class="email-container">
        <!-- 头部 -->
        <div class="email-header">
            <h1>验证码通知</h1>
        </div>
        
        <!-- 主体 -->
        <div class="email-body">
            <p>您好，</p>
            <p>感谢您使用我们的服务！您本次的验证码如下：</p>
            <div class="email-code">%s</div>
            <p>请注意：</p>
            <ul>
                <li>该验证码有效期为<strong>5分钟</strong>。</li>
                <li>如果您未尝试登录或注册，请忽略此邮件。</li>
            </ul>
            <p>感谢您的支持！如有任何疑问，请随时联系我们。</p>
        </div>
        
        <!-- 底部 -->
        <div class="email-footer">
            <p>此邮件由系统自动发送，请勿回复。</p>
            <p>如果您需要帮助，请访问我们的<a href="https://example.com/help">帮助中心</a>。</p>
        </div>
    </div>
</body>
</html>`