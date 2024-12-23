package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shangcheng/Project/handlers"
	"github.com/shangcheng/Project/internal/dao"
	"github.com/shangcheng/Project/internal/services"
	"gorm.io/gorm"
)

// SetupPayRouter 设置支付相关路由
func SetupPayRouter(r *gin.Engine, db *gorm.DB) {
	// 创建 DAO 实例
	orderDao := &dao.OrderDao{DB: db}
	userDao := &dao.UserDao{DB: db}
	productDao := &dao.ProductDao{DB: db}

	// 创建 PayService 实例，并传入所有 DAO 依赖
	payService := services.NewPayService(orderDao, userDao, productDao)

	// 创建 PayHandler 实例，并传入 PayService 依赖
	payHandler := handlers.NewPayHandler(payService)

	// 定义支付路由组
	payRoutes := r.Group("/api/v1/payments")

	{
		// 处理支付请求
		payRoutes.POST("/pay", payHandler.Payment)
	}
}
