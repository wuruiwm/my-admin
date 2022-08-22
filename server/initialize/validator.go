package initialize

import (
	"app/util"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translation "github.com/go-playground/validator/v10/translations/zh"
)

func Validator() {
	util.Trans, _ = ut.New(zh.New()).GetTranslator("zh")
	err := translation.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), util.Trans)
	if err != nil {
		panic("gin validator translator register error: " + err.Error())
	}
}
