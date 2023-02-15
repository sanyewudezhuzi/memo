package model

import (
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

// MySQL 初始化
func ConnectMySQL(dsn string) {
	// 连接数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 171, // 默认字符串长度
	}), &gorm.Config{
		// 日志格式
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,   // 慢 sql 阈值
				LogLevel:      logger.Silent, // log level
				Colorful:      false,         // 禁用色彩打印
			},
		),
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "memo_", // 表名增加前缀
			SingularTable: true,    // 使用表名单数
		},
		// 禁用外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("MySQL connection failed.")
	}

	// 设置连接池
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)           // 空闲连接池最大数量
	sqlDB.SetMaxOpenConns(100)          // 打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接可复用的最大时间

	// 设置全局 DB
	DB = db
}

// 数据库迁移
func SyncSchema() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Task{})
}
