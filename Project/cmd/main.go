package main

import (
	"log"
	"net/http"

	"github.com/shangcheng/Project/Project/internal/config"
	"github.com/shangcheng/Project/Project/routers"
)

func main() {
	// 初始化配置和数据库连接
	if err := config.InitDB(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	// 创建 Gin 路由实例
	r := routers.SetupRouter(config.DB)

	// 获取端口号
	port := config.GetPort()
	log.Printf("Starting server on %s", port)

	// 启动 HTTP 服务器
	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	// 监听系统信号
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
   }()
}