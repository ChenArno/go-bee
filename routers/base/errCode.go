package base

// 定义错误
type Err map[string]Errno

type Errno struct {
	Code         int
	Msg          string
	resultStatus bool
}

/**
五位错误码设计
第一位表示错误级别, 1 为系统错误, 2 为普通错误
第二三位表示服务模块代码
第四五位表示具体错误代码
*/
var ALLCODE_MESSAGE = Err{
	"OK":        Errno{Code: 0, Msg: "OK", resultStatus: true},
	"UnknowErr": Errno{Code: 10000, Msg: "未知错误", resultStatus: false},

	// 系统错误, 前缀为 100
	"InternalServerError": Errno{Code: 10001, Msg: "内部服务器错误", resultStatus: false},
	"ErrBind":             Errno{Code: 10002, Msg: "请求参数错误", resultStatus: false},
	"ErrTokenSign":        Errno{Code: 10003, Msg: "签名 jwt 时发生错误", resultStatus: false},
	"ErrEncrypt":          Errno{Code: 10004, Msg: "加密用户密码时发生错误", resultStatus: false},

	// 数据库错误, 前缀为 201
	"ErrDatabase": Errno{Code: 20100, Msg: "数据库错误", resultStatus: false},
	"ErrFill":     Errno{Code: 20101, Msg: "从数据库填充 struct 时发生错误", resultStatus: false},
	"HandleDbErr": Errno{Code: 20102, Msg: "从数据库执行错误", resultStatus: false},

	// 认证错误, 前缀是 202
	"ErrValidation":   Errno{Code: 20201, Msg: "验证失败", resultStatus: false},
	"ErrTokenInvalid": Errno{Code: 20202, Msg: "jwt 是无效的", resultStatus: false},
	"WITHOUTLOGIN":    Errno{Code: 20203, Msg: "账号未登录", resultStatus: false},

	// 用户错误, 前缀为 203
	"ErrUserNotFound":      Errno{Code: 20301, Msg: "用户没找到", resultStatus: false},
	"ErrPasswordIncorrect": Errno{Code: 20302, Msg: "密码错误", resultStatus: false},
}

// 调用错误返回函数
func CodeMessage(errName string) Errno {
	if errName == "" {
		return ALLCODE_MESSAGE["OK"]
	}
	if v, ok := ALLCODE_MESSAGE[errName]; ok {
		return v
	} else {
		return ALLCODE_MESSAGE["UnknowErr"]
	}
}
