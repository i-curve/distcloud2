package e

// 返回码定义
const (
	// 基本返回码
	SUCCESS        = 200
	INVALID_PARAMS = 400
	ERROR          = 500
	// 特殊返回码
	CODE_ERROR = 600
	// token相关返回码
	ERROR_AUTH                     = 700
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 701
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 702
	ERROR_AUTH_GENERNATE_TOKEN     = 703
	ERROR_PRIVILEGES               = 704
	// user相关返回码
	USER_ALREADY_EXIST  = 800
	USER_NOT_EXIST      = 801
	USER_PASSWORD_ERROR = 802
)
