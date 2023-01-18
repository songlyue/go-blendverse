package initialize

import (
	"blendverse/global"
	"blendverse/utils"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

type Func func(fl validator.FieldLevel) bool

func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New()
		enT := en.New()
		uni := ut.New(enT, zhT, enT)
		var ok bool
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("")
		}
		switch locale {
		case "en":
			_ = en_translations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			_ = zh_translations.RegisterDefaultTranslations(v, global.Trans)
		default:
			_ = en_translations.RegisterDefaultTranslations(v, global.Trans)
		}
		// 注册自定义校验
		RegisterValidatorFunc(v, "mobile", "手机号违法", utils.ValidateMobile)

		return
	}
	return
}

func RegisterValidatorFunc(v *validator.Validate, tag string, msgStr string, fn Func) {
	// 注册tag校验
	_ = v.RegisterValidation(tag, validator.Func(fn))
	// 自定义错误内容
	_ = v.RegisterTranslation(tag, global.Trans, func(ut ut.Translator) error {
		return ut.Add(tag, "{0}"+msgStr, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())
		return t
	})
	return
}
