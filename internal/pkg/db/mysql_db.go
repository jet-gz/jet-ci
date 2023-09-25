package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var ORM *gorm.DB

func mysqlOpen() error {

	dsn := options.ConnectStr
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 这里可以设置 别的配置
		//DryRun: false,//生成 SQL 但不执行，可以用于准备或测试生成的 SQL
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err == nil {
		ORM = db
		return err
	}
	return err
}
