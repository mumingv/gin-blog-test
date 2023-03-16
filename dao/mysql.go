package dao

import (
	"fmt"
	"github.com/mumingv/gin-blog/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

func InitMySQL(cfg *settings.MySQLConfig) (err error) {
	// 连接 mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "tb_",
		},
	})
	if err != nil {
		return
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Ping(); err != nil {
		fmt.Println("连接MySQL失败")
	}

	return nil
}

func Close() {
	sqlDB, _ := DB.DB()
	sqlDB.Close()
}
