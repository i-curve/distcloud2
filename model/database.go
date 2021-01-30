package model

import (
	"clouddist/pkg/setting"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

// Setup 初始化数据库函数
func Setup() {
	var err error
	logMode := logger.Silent
	if setting.ServerSetting.RunMode == "debug" {
		logMode = logger.Info
	}
	db, err = gorm.Open(
		mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				setting.DatabaseSetting.User,
				setting.DatabaseSetting.Password,
				setting.DatabaseSetting.Host,
				setting.DatabaseSetting.Name,
			),
		),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				// TablePrefix:   "t_", // 表名前缀，`User` 的表名应该是 `t_users`
				TablePrefix:   setting.DatabaseSetting.TablePrefix,
				SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
			},
			Logger: logger.Default.LogMode(logMode),
		},
	)
	if err != nil {
		log.Fatalf("connect the mysql error: '%v'", err)
	}
	DB, _ := db.DB()
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
}

// CloseDB 关闭数据库连接
func CloseDB() {
	log.Println("关闭数据库连接")
	defer db.Statement.ReflectValue.Close()
}
