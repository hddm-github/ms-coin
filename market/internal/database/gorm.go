package database

import (
	"log"
	"mscoin-common/msdb"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlConfig struct {
	DataSource string
}

func ConnMysql(c MysqlConfig) *msdb.MsDB {
	var err error

	// 自定义 GORM 日志器，将慢 SQL 阈值延长至 2 秒，避免频繁记录因远程公网数据库带来的慢日志警告
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             2 * time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	_db, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	db, _ := _db.DB()
	//连接池配置
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return &msdb.MsDB{
		_db,
	}
}
