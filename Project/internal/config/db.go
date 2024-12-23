package config

import (
	"github.com/shangcheng/Project/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := "root:123456@tcp(192.168.0.44:3306)/shangcheng?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// 自动迁移
	DB.AutoMigrate(&models.User{})
	return nil
}
