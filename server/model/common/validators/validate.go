package validators

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ja"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	et "github.com/go-playground/validator/v10/translations/en"
	zt "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"regexp"
	"strings"
)

var Trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(locale string) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := "-"
			if fld.Tag.Get("json") != "" {
				name = strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			} else if fld.Tag.Get("form") != "" {
				name = strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
			} else if fld.Tag.Get("query") != "" {
				name = strings.SplitN(fld.Tag.Get("query"), ",", 2)[0]
			}
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器
		jaT := ja.New() // 日语翻译器
		uni := ut.New(enT, enT, zhT, jaT)
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			fmt.Printf("uni.GetTranslator(%s) failed", locale)
			return
		}
		// 注册新的自定义验证规则
		regCustom(v, Trans)
		// 注册翻译器
		switch locale {
		case "en":
			_ = et.RegisterDefaultTranslations(v, Trans)
		case "ja":
			_ = et.RegisterDefaultTranslations(v, Trans)
		case "zh":
			_ = zt.RegisterDefaultTranslations(v, Trans)
		default:
			_ = et.RegisterDefaultTranslations(v, Trans)
		}
	}
}

func regCustom(v *validator.Validate, Trans ut.Translator) {
	for _, eoo := range []struct {
		tag string
		fn  func(fl validator.FieldLevel) bool
		msg string
	}{
		{"phone", phone, "Mobile number format is incorrect"},
		//{"password", password, "The password format is entered incorrectly"},
		//{"excludeChinese", excludeChinese, "cannot contain Chinese"},
	} {
		_ = v.RegisterValidation(eoo.tag, eoo.fn)
		// 添加额外翻译
		_ = v.RegisterTranslation(eoo.tag, Trans, func(ut ut.Translator) error {
			return ut.Add(eoo.tag, "{0} "+eoo.msg, false)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(fe.Tag(), fe.Field())
			return t
		})
	}
}

//手机号码验证
func phone(fl validator.FieldLevel) bool {
	regular := "^((13[0-9])|(14[0-9])|(15[0-9])|(16[0-9])|(17[0-9])|(18[0-9])|(19[0-9]))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(fl.Field().String())
}

//
////密码校验
//func password(fl validator.FieldLevel) bool {
//	ps := fl.Field().String()
//	// 默认密码必须到达100分也就是 大小写字母加数字加特殊字符
//	v := 100
//	if param, err := strconv.Atoi(fl.Param()); err == nil {
//		v = param
//	}
//	return utils.PassWordCheck(ps) >= v
//}
//
//// 禁止出现中文
//func excludeChinese(fl validator.FieldLevel) bool {
//	if utils.IsChineseChar(fl.Field().String()) {
//		return false
//	}
//	return true
//}
