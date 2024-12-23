package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shangcheng/Project/internal/config"
	"github.com/shangcheng/Project/routers"
)

func main() {
	// 初始化配置和数据库连接
	if err := config.InitDB(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	// 设置 Gin 模式）
	gin.SetMode(gin.ReleaseMode)

	// 创建 Gin 路由实例
	r := routers.SetupRouter(config.DB)

	// 启动 HTTP 服务器
	port := ":8080"
	log.Printf("Starting server on %s", port)
	if err := r.Run(port); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to start server: %v", err)
	}
}
