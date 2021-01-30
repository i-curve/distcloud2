package v1

import (
	"clouddist/pkg/app"
	"clouddist/pkg/e"
	"clouddist/pkg/file"
	"clouddist/service/data_service"
	"clouddist/service/user_service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetFiles 获取目录下所有内容
func GetFiles(c *gin.Context) {
	appG := app.Gin{C: c}

	id, _ := c.Get("id")
	userService := user_service.User{ID: id.(int)}

	status := userService.GetMsg()
	if !status {
		appG.Response(http.StatusOK, e.ERROR, nil)
	}

	loc := c.DefaultQuery("path", "/")
	data := data_service.GetPath(userService.Username, loc)
	log.Println(data)
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// DownloadFile 文件下载
func DownloadFile(c *gin.Context) {
	// appG := app.Gin{C: c}

	id, _ := c.Get("id")
	userService := user_service.User{ID: id.(int)}

	userService.GetMsg()
	// if !status {
	// appG.Response(http.StatusOK, e.ERROR, nil)
	// }

	FilePath := c.DefaultQuery("path", "/")
	FileName := c.DefaultQuery("filename", "")
	PATH := "./data/" + userService.Username + FilePath

	log.Println(PATH, PATH+FileName)

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", FileName))
	c.Writer.Header().Set("Content-Type", "text/plain")
	// appG := app.Gin{C: c}
	c.File(PATH + FileName)
}

// UploadFile 上传文件
func UploadFile(c *gin.Context) {
	appG := app.Gin{C: c}

	FilePath := c.DefaultQuery("path", "/")
	id, _ := c.Get("id")
	userService := user_service.User{ID: id.(int)}
	userService.GetMsg()

	file, err := c.FormFile("upload")
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	filename := file.Filename
	PATH := "./data/" + userService.Username + FilePath
	c.SaveUploadedFile(file, PATH+filename)
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// DeleteFile 删除文件
func DeleteFile(c *gin.Context) {
	appG := app.Gin{C: c}

	id, _ := c.Get("id")
	userService := user_service.User{ID: id.(int)}
	userService.GetMsg()

	filepath := c.PostForm("path")
	filename := c.PostForm("file")
	PATH := "./data/" + userService.Username + filepath
	file.MkDel(PATH + filename)
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// NewFolder 创建文件夹
func NewFolder(c *gin.Context) {
	appG := app.Gin{C: c}

	id, _ := c.Get("id")
	userService := user_service.User{ID: id.(int)}
	userService.GetMsg()
	log.Println(id)

	filepath := c.PostForm("path")
	filename := c.PostForm("name")
	PATH := "./data/" + userService.Username + filepath
	log.Println(PATH, filename)
	file.MkDir(PATH + filename)
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
