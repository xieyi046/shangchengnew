package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config 应用的全局配置结构体
type Config struct {
	Port       string
	DBUsername string
	DBPassword string
	DBHost     string // 数据库主机地址
	DBPort     string // 数据库端口
	DBName     string
}

var Cfg *Config
// LoadConfig 加载配置文件并初始化配置结构体
func LoadConfig() error {
	// 加载 .env 文件中的环境变量（可选）
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	Cfg = &Config{
		Port:       os.Getenv("PORT"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"), // 新增环境变量
		DBPort:     os.Getenv("DB_PORT"), // 新增环境变量
		DBName:     os.Getenv("DB_NAME"),
	}

	// 如果没有设置 PORT 环境变量，则使用默认值
	if Cfg.Port == "" {
		Cfg.Port = "8080" // 默认端口
	}

	return nil
}

// GetPort 获取服务器监听的端口号
func GetPort() string {
	if Cfg == nil {
		log.Fatal("Configuration not loaded")
	}
	return Cfg.Port
}

