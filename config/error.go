package config

import (
	"github.com/Looper56/plugin/web"
)

var (
	UnauthorizedError web.ErrKey = "unauthorized_error" //无权限
)

// ErrorConfig 异常状态码
var ErrorConfig = map[web.ErrKey]web.ErrorConfig{
	UnauthorizedError: {40001, 401, "无访问权限"},
}
