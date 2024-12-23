package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shangcheng/Project/Project/handlers"
	"github.com/shangcheng/Project/Project/internal/dao"
	"github.com/shangcheng/Project/Project/internal/services"
	"gorm.io/gorm"
)

// SetupPaysRouter 设置支付相关路由
func SetupPaysRouter(r *gin.Engine, db *gorm.DB) {
	// 创建 DAO 实例
	orderDao := &dao.OrderDao{DB: db}
	userDao := &dao.UserDao{DB: db}
	productDao := &dao.ProductDao{DB: db}

	// 创建 PaysService 实例，并传入所有 DAO 依赖
	paysService := services.NewPaysService(orderDao, userDao, productDao)

	// 创建 PaysHandler 实例，并传入 PaysService 依赖
	paysHandler := handlers.NewPaysHandler(paysService)

	// 定义支付路由组
	paysRoutes := r.Group("/api/v1/paysments")

	{
		// 处理支付请求
		paysRoutes.POST("/pay", paysHandler.MultiPayment)
	}
}
