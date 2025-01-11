package pkg

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/go-playground/locales/en"                            // 英语语言包的资源
	"github.com/go-playground/locales/zh"                            // 中文语言包的资源
	ut "github.com/go-playground/universal-translator"               // 翻译器
	zhtrans "github.com/go-playground/validator/v10/translations/zh" // 汉语翻译包
)

// 全局验证器
var Validate *validator.Validate
// 全局验证器的翻译器
var Trans ut.Translator

func Init() error {
    // 获得gin框架的validator引擎（后面还可以定制语言包）
    Validate, _= binding.Validator.Engine().(*validator.Validate) 

	// 注册GetTag，获取参数结构体中的tag
    Validate.RegisterTagNameFunc(GetTag)
	// 初始化翻译器（固定为汉语翻译包）
    InitTrans("zh")

    // 注册验证规则（名字 -- 方法 映射）

	// 用户名验证
    // if err := Validate.RegisterValidation("phone", validatePhone); err != nil {
    //     zap.L().Error("用户名验证规则注册失败", zap.Error(err))
    //     return err
    // }

	// 密码验证
    if err := Validate.RegisterValidation("password", validatePassword); err != nil {
        zap.L().Error("密码验证规则注册失败", zap.Error(err))
        return err
    }

    // 邮箱验证
	if err := Validate.RegisterValidation("email", validateEmail); err != nil {
		zap.L().Error("邮箱验证规则注册失败", zap.Error(err))
		return err
	}

    // 验证码验证
	if err := Validate.RegisterValidation("code", validateCode); err != nil {
		zap.L().Error("验证码验证规则注册失败", zap.Error(err))
		return err
	}

    // 注册验证规则对应的翻译包
    //Validate.RegisterTranslation("phone", Trans, tranPhoneRegis, tranPhone)
    Validate.RegisterTranslation("password", Trans, tranPasswordRegis, tranPassword)
    Validate.RegisterTranslation("email", Trans, tranEmailRegis, tranEmail)
    Validate.RegisterTranslation("code", Trans, tranCodeRegis, tranCode)

    return nil
}

// 另外的作用（可选）：用于删除错误信息前面的结构体信息
func RemoveStructName(fields map[string]string) string {
    result := map[string]string{}
    for field, err := range fields {
        result[field[strings.Index(field, ".")+1:]] = err
    }
	errStr := ""
	for _, v := range result {
		errStr += v + "; "
	}
    return errStr
}

// 初始化翻译器：根据传入的方言选择翻译方向
func InitTrans(locale string) (err error) {
    en := en.New()
    zh := zh.New()
    uni := ut.New(en, zh)
	var ok bool
    // 注意：这里一定要给全局的trans，否则空指针报错
    Trans, ok = uni.GetTranslator(locale)
    if!ok {
        return fmt.Errorf("无法获取 %s 语言包对应的翻译器",locale)
    }
    switch locale {
    case "zh":
        err = zhtrans.RegisterDefaultTranslations(Validate, Trans)
    default:
        return fmt.Errorf("不支持的语言环境: %s",locale)
    }
    return
}

//获取参数结构体中的tag
func GetTag(fld reflect.StructField) string {
    name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
    if name == "-" {
        return ""
    }
    return name
}

// 定义验证规则

//------------------------------------------------------------------------------------------
// 对用户的用户名字段验证
// func validatePhone(fl validator.FieldLevel) bool{
//     // 利用反射拿到字段
//     phone := fl.Field().String()
//     regex := `^1[3-9]\d{9}$`
//     // 使用regexp包进行正则表达式匹配
//     matched, _ := regexp.MatchString(regex, phone)
//     return matched
// }

// // 自定义校验函数对应的注册翻译函数
// func tranPhoneRegis(ut ut.Translator) error {
// 	// 添加特定标签和翻译的错误信息
//     return ut.Add("phone", "非法的用户名!", true)
// }

// // 自定义校验函数对应的翻译函数
// func tranPhone(ut ut.Translator, fe validator.FieldError) string {
// 	// 拼接错误信息后返回
//     t, _ := ut.T("phone")
//     return t
// }
//------------------------------------------------------------------------------------------

// go语言的regexp包不支持预先断言：regex := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{6,}$` 中 (?=
func validatePassword(fl validator.FieldLevel) bool{
    password := fl.Field().String()
    // 至少一个大写字母，一个小写字母，一个数字，长度>=6个字符
    if len(password) < 6 {
        return false
    }
    hasLower := false
    hasUpper := false
    hasDigit := false
    for _, char := range password {
        switch {
        case 'a' <= char && char <= 'z':
            hasLower = true
        case 'A' <= char && char <= 'Z':
            hasUpper = true
        case '0' <= char && char <= '9':
            hasDigit = true
        }
    }
    return hasLower && hasUpper && hasDigit
}

// 自定义校验函数对应的注册翻译函数
func tranPasswordRegis(ut ut.Translator) error {
	// 添加特定标签和翻译的错误信息
    return ut.Add("password", "非法的密码!", true)
}

// 自定义校验函数对应的翻译函数
func tranPassword(ut ut.Translator, fe validator.FieldError) string {
	// 拼接错误信息后返回
    t, _ := ut.T("password")
    return t
}
//-----------------------------------------------------------------------

// 自定义邮箱验证规则
func validateEmail(fl validator.FieldLevel) bool {
    email := fl.Field().String()
    // qq 邮箱验证规则
    regex := `^[a-zA-Z0-9_-]+@qq\.com$`
    matched, _ := regexp.MatchString(regex, email)
    return matched
}

// 自定义校验函数对应的注册翻译函数
func tranEmailRegis(ut ut.Translator) error {
    return ut.Add("email", "必须是qq邮箱", true)
}

// 自定义校验函数对应的翻译函数
func tranEmail(ut ut.Translator, fe validator.FieldError) string {
    // 拼接错误信息后返回
    t, _ := ut.T("email")
    return t
}

//-----------------------------------------------------------------------

// 自定义验证码校验函数
func validateCode(fl validator.FieldLevel) bool {
    code := fl.Field().String()
    // 验证码就是6位数字的字符串
    return regexp.MustCompile(`^[0-9]{6}$`).MatchString(code)
}

// 自定义验证码校验函数对应的注册翻译函数
func tranCodeRegis(ut ut.Translator) error {
    return ut.Add("code", "必须是6位数字的字符串", true)
}

// 自定义验证码校验函数对应的翻译函数
func tranCode(ut ut.Translator, fe validator.FieldError) string {
    t, _ := ut.T("code")
    return t
}

//---------------------------------------------