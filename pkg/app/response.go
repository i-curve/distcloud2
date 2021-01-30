package app

import (
	"clouddist/pkg/e"

	"github.com/gin-gonic/gin"
)

// Gin 获取前端对象
type Gin struct {
	C *gin.Context
}

// Response 规范化返回操作
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})
}
