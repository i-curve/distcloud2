package main

import (
	"clouddist/model"
	"clouddist/pkg/cloud"
	"clouddist/pkg/logging"
	"clouddist/pkg/setting"
	"clouddist/router"
	"fmt"
	"log"
	"syscall"

	"github.com/fvbock/endless"
)

func init() {
	setting.Setup()
	logging.Setup()
	model.Setup()
	cloud.Setup()
	// cloud.CreateUSER("admin")
	// cloud.ShowDir("admin", "/Documents")
}

// main 程序入口
func main() {
	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort)

	server := endless.NewServer(endPoint, router.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err:%v", err)
	}
	if setting.ServerSetting.RunMode == "release" {
		logging.Info("运行成功...")
	}
	model.CloseDB()
}
