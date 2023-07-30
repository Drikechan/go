package dao

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"test-todolist/config"
	"time"
)

var (
	DB *gorm.DB
)

func InitMysql() {
	conn := strings.Join([]string{
		config.DbConfig.UserName, ":",
		config.DbConfig.Password,
		"@tcp(", config.DbConfig.Addr, ":",
		config.DbConfig.Port, ")/",
		config.DbConfig.DbName,
		"?charset=utf8&parseTime=true&loc=Local&timeout=1000ms"}, "")

	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		Colorful:      true,
	})

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("%+v", err)
		return
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) // 打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
}

func NewClient(ctx context.Context) *gorm.DB {
	db := DB
	return db.WithContext(ctx)
}
