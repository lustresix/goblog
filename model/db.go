package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

// DB 数据库配置
var DB *gorm.DB
var err error
var ormLogger logger.Interface

func Database(connString string) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		// DSN data source name
		DSN: connString,
		// string 类型字段的默认长度
		DefaultStringSize: 256,
		// 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DisableDatetimePrecision: true,
		// 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameIndex: true,
		// 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		DontSupportRenameColumn: true,
		// 根据版本自动配置
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			// 表不加s，user---users
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println(err)
		panic("数据库连接错误...")
	}
	fmt.Println("数据库连接成功...", err)

	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	// 自动迁移模式
	db.AutoMigrate(&User{}, &Article{}, &Category{})

	sqlDB, _ := db.DB()
	// 设置连接池
	sqlDB.SetMaxIdleConns(20)
	// 打开
	sqlDB.SetMaxOpenConns(100)
	// 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	DB = db
}
