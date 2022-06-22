package validator

// ref : https://qiita.com/syoimin/items/b3923fea6070b0a3df8f
// ref : https://github.com/go-playground/validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func Run(m interface{}) error {
	validate.RegisterValidation("notblank", notBlank)
	validate.RegisterValidation("include_low", includeLowercase)
	validate.RegisterValidation("include_up", includeUppercase)
	validate.RegisterValidation("include_num", includeNumeric)
	return validate.Struct(m)
}

// 空白で無いかどうか
func notBlank(fl validator.FieldLevel) bool {
	return len(fl.Field().String()) > 0
}

// 小文字が含まれるかどうか
func includeLowercase(fl validator.FieldLevel) bool {
	return checkRegexp("[a-z]", fl.Field().String())
}

// 大文字が含まれるかどうか
func includeUppercase(fl validator.FieldLevel) bool {
	return checkRegexp("[A-Z]", fl.Field().String())
}

// 数値が含まれるかどうか
func includeNumeric(fl validator.FieldLevel) bool {
	return checkRegexp("[0-9]", fl.Field().String())
}

// 特殊記号が含まれるかどうか
// func includeSymbol(fl validator.FieldLevel) bool {
//     availableChar := checkRegexp(`^[0-9a-zA-Z\-^$*.@]+$`, fl.Field().String())
//     checkIsSymbol := checkRegexp(`[\-^$*.@]`, fl.Field().String())
//
//     return availableChar && checkIsSymbol
// }

// 正規表現共通関数
func checkRegexp(reg, str string) bool {
	r := regexp.MustCompile(reg).Match([]byte(str))
	return r
}
