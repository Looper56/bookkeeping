package plugin

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var uidRegexp = regexp.MustCompile("^[a-zA-Z0-9]{9,12}$")

var uidValidator validator.Func = func(fl validator.FieldLevel) bool {
	uid, ok := fl.Field().Interface().(string)
	if ok {
		return uidRegexp.MatchString(uid)
	}
	return false
}
