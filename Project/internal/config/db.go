package config

import (
	"fmt"

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
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Product{},
		&models.Order{},
		&models.Pay{},
		&models.MultiPay{},
		&models.PayItem{},
	}

	for _, model := range modelsToMigrate {
		if err := DB.AutoMigrate(model); err != nil {
			return fmt.Errorf("failed to migrate table for model %T: %w", model, err)
		}
	}

	return nil
}
