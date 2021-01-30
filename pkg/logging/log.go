package logging

import (
	"clouddist/pkg/file"
	"clouddist/pkg/setting"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Level 日志等级
type Level int

// 相关文件定义
var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	logger             *log.Logger
	logPrefix          = ""
	levelFlags         = []string{"INFO", "DEBUG", "WARN", "ERROR", "FATAL"}
)

// 等级
const (
	INFO Level = iota
	DEBUG
	WARNING
	ERROR
	FATAL
)

// Setup 初始化
func Setup() {
	if setting.AppSetting.LogSavePath != "" {
		LogSavePath = setting.AppSetting.RuntimeRootPath + setting.AppSetting.LogSavePath
	}
	if setting.AppSetting.LogSaveName != "" {
		LogSaveName = setting.AppSetting.LogSaveName
	}
	if setting.AppSetting.LogFileExt != "" {
		LogFileExt = setting.AppSetting.LogFileExt
	}
	if setting.ServerSetting.RunMode == "release" {
		TimeFormat = setting.AppSetting.TimeFormat
	}

	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup: %v", err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// Info 正常显示
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

// Debug 调试
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

// Warn 警告
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

// Error 错误
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

// Fatal 致命错误
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Println(v)
}

// setPrefix 内部方法,文件命令
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
