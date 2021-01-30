package e

// MsgFlags 返回码-信息对照表
var MsgFlags = map[int]string{
	// 基本返回码
	SUCCESS:        "ok",
	INVALID_PARAMS: "请求参数错误",
	ERROR:          "服务器内部错误",
	// 特殊返回码
	CODE_ERROR: "该返回码错误, 请检查e/code",
	// token返回码
	ERROR_AUTH:                     "认证失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_GENERNATE_TOKEN:     "Token生成失败",
	ERROR_PRIVILEGES:               "权限不足",
	// user相关返回码
	USER_ALREADY_EXIST:  "用户已经存在",
	USER_NOT_EXIST:      "用户不存在",
	USER_PASSWORD_ERROR: "密码错误",
}

// GetMsg 获取返回码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[CODE_ERROR]
}
