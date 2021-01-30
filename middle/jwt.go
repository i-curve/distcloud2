package middle

import (
	"clouddist/pkg/e"
	"clouddist/pkg/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT 中间件验证
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var err error
		var data interface{}
		var claims *util.Claims
		code = e.SUCCESS

		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}

		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err = util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Set("id", claims.ID)
		c.Set("privilege", claims.Privilege)
		c.Next()
	}
}
