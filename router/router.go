package router

import (
	"clouddist/middle"
	"clouddist/pkg/setting"
	v1 "clouddist/router/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由表
func InitRouter() *gin.Engine {
	gin.SetMode(setting.ServerSetting.RunMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.Default())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "ok",
		})
	})
	router.POST("/login", v1.GetAuth)
	router.POST("/register", v1.RegisterUser)
	apiv1 := router.Group("cloud")
	apiv1.Use(middle.JWT())
	{
		apiv1.GET("/show", v1.GetFiles)
		apiv1.GET("/download", v1.DownloadFile)
		apiv1.POST("/upload", v1.UploadFile)
		apiv1.POST("/new", v1.NewFolder)
		apiv1.POST("/delete", v1.DeleteFile)
	}
	return router
}
