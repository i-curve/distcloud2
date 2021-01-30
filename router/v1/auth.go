package v1

import (
	"clouddist/pkg/app"
	"clouddist/pkg/e"
	"clouddist/pkg/util"
	"clouddist/service/data_service"
	"clouddist/service/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// RegisterUser 注册用户
func RegisterUser(c *gin.Context) {
	appG := app.Gin{C: c}

	var a auth
	a.Username = c.PostForm("username")
	a.Password = c.PostForm("password")

	userService := user_service.User{Username: a.Username}
	status := userService.GetMsg()
	if status {
		appG.Response(http.StatusOK, e.USER_ALREADY_EXIST, nil)
		return
	}
	userService.Password = a.Password
	status, err := userService.Register()
	if status != true || err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	status = data_service.CreatePath(userService.Username)
	if !status {
		appG.Response(http.StatusOK, e.ERROR, nil)
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// GetAuth 获取token
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}

	var a auth
	a.Username = c.PostForm("username")
	a.Password = c.PostForm("password")

	userService := user_service.User{Username: a.Username}

	status := userService.GetMsg()
	if !status {
		appG.Response(http.StatusOK, e.USER_NOT_EXIST, nil)
		return
	}

	status = userService.Check(a.Password)
	if !status {
		appG.Response(http.StatusOK, e.USER_PASSWORD_ERROR, nil)
		return
	}

	token, err := util.GenerateToken(userService.ID, userService.Privilege, a.Username, a.Password)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_AUTH_GENERNATE_TOKEN, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"username": userService.Username,
		"token":    token,
	})
}
