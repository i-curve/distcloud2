package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

// App 全局设置定义
type App struct {
	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	DataPath string
}

// AppSetting 全局设置变量
var AppSetting = &App{}

// Server 本程序设置结构体
type Server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// ServerSetting 本程序变量
var ServerSetting = &Server{}

// Database 数据库设置结构体
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

// DatabaseSetting 数据库设置变量
var DatabaseSetting = &Database{}

var cfg *ini.File

// Setup 初始化函数
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("load the app.ini file error: \"%v\"", err)
	}
	MapTo("app", AppSetting)
	MapTo("server", ServerSetting)
	MapTo("database", DatabaseSetting)
	if ServerSetting.RunMode != "release" {
		log.Println("\nRun mode:   ", ServerSetting.RunMode)
	}
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

// MapTo 读取配置文件并映射到相应结构体中
func MapTo(name string, v interface{}) {
	err := cfg.Section(name).MapTo(v)
	if err != nil {
		log.Fatalf("load config file error: '%v'", err)
	}
}
